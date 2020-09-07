package stupid

import (
	"github.com/gogo/protobuf/proto"
	"sync"
)

var (
	registerInstance     *msgFactoryRegister
	registerInstanceOnce = &sync.Once{}
)

type factory struct {
	produce func() proto.Message
	pool    *sync.Pool
	usePool bool
}

func (f *factory) take() proto.Message {
	if f.usePool {
		return f.pool.Get().(proto.Message)
	} else {
		return f.produce()
	}
}

func (f *factory) release(pb proto.Message) {
	if f.usePool {
		f.pool.Put(pb)
	}
}

type msgFactoryRegister struct {
	factoryMap map[uint8]*factory
}

func MsgFactoryRegisterInstance() *msgFactoryRegister {
	registerInstanceOnce.Do(func() {
		registerInstance = &msgFactoryRegister{
			factoryMap: make(map[uint8]*factory, 8),
		}
	})
	return registerInstance
}

func (m *msgFactoryRegister) RegMessageFactory(id uint8, usePool bool, f func() proto.Message) {
	m.factoryMap[id] = &factory{
		produce: f,
	}
	if usePool {
		m.factoryMap[id].pool = &sync.Pool{New: func() interface{} { return f() }}
		m.factoryMap[id].usePool = true
	}
}
