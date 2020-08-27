package stupid

import (
	"github.com/gogo/protobuf/proto"
	"sync"
)

var (
	registerInstance     *MsgFactoryRegister
	registerInstanceOnce = &sync.Once{}
)

type factory struct {
	produce func() proto.Message
	pool    *sync.Pool
}

type MsgFactoryRegister struct {
	factoryMap map[uint8]*factory
}

func DefaultRegisterInstance() *MsgFactoryRegister {
	registerInstanceOnce.Do(func() {
		registerInstance = &MsgFactoryRegister{
			factoryMap: make(map[uint8]*factory, 8),
		}
	})
	return registerInstance
}

func (m *MsgFactoryRegister) RegMessageFactory(id uint8, usePool bool, f func() proto.Message) {
	m.factoryMap[id] = &factory{
		produce: f,
	}
	if usePool {
		m.factoryMap[id].pool = &sync.Pool{New: func() interface{} { return f() }}
	}
}
