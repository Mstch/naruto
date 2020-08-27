package stupid

import (
	"errors"
	"github.com/gogo/protobuf/proto"
	"net"
	"sync"
)

type clientImpl struct {
	conn     net.Conn
	regLock  *sync.Mutex
	handlers map[string]*handler
}

func NewClientImpl() *clientImpl {
	return &clientImpl{
		regLock: &sync.Mutex{},
	}
}
func (c *clientImpl) Dial(address string) error {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}
	c.conn = conn
	go serveConn(c.handlers, conn)
	return nil
}

func (c *clientImpl) Call(name string, arg proto.Message) (res proto.Message, err error) {
	panic("implement me")
}

func (c clientImpl) AsyncCall(name string, arg proto.Message) (resC chan proto.Message, err error) {
	panic("implement me")
}

func (c *clientImpl) Notify(name string, arg proto.Message) error {
	return write(c.conn, name, arg)
}

func (c *clientImpl) RegHandler(name string, h func(arg proto.Message) (res proto.Message), argId, resId uint8) error {
	if len(name) > 255 {
		return errors.New("name too lang, maxsize is 255")
	}
	c.regLock.Lock()
	c.handlers[name] = &handler{
		handleFunc: h,
		argId:      argId,
		resId:      resId,
	}
	c.regLock.Unlock()
	return nil
}
