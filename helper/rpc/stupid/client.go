package stupid

import (
	"errors"
	"github.com/gogo/protobuf/proto"
	"net"
	"sync"
)

type clientImpl struct {
	connected bool
	conn      net.Conn
	regLock   *sync.Mutex
	handlers  map[string]*handler
}

func NewClientImpl() *clientImpl {
	return &clientImpl{
		regLock:  &sync.Mutex{},
		handlers: make(map[string]*handler),
	}
}

//func (c *clientImpl) Dial(address string) error {
//	if c.connected {
//		return errors.New("there is already a connection successfully")
//	}
//	conn, err := net.Dial("tcp", address)
//	if err != nil {
//		return err
//	}
//	c.connected = true
//	c.conn = conn
//	go serveConn(c.handlers, conn)
//	return nil
//}

func (c *clientImpl) Conn(conn net.Conn) error {
	if c.connected {
		return errors.New("there is already a connection successfully")
	}
	c.connected = true
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

func (c *clientImpl) RegHandler(name string, h func(arg proto.Message), argId uint8) error {
	if len(name) > 255 {
		return errors.New("name too lang, maxsize is 255")
	}
	if _, ok := MsgFactoryRegisterInstance().factoryMap[argId]; !ok {
		return errors.New("argId not registered")
	}
	c.regLock.Lock()
	c.handlers[name] = &handler{
		handleFunc: func(arg proto.Message) (res proto.Message) {
			h(arg)
			return nil
		},
		argId: argId,
	}
	c.regLock.Unlock()
	return nil
}
func (c *clientImpl) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}
