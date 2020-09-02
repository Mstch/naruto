package stupid

import (
	"errors"
	"github.com/Mstch/naruto/helper/logger"
	"github.com/gogo/protobuf/proto"
	"net"
	"sync"
)

var (
	serverInstance     *serverImpl
	serverInstanceOnce = &sync.Once{}
)

type serverImpl struct {
	regLock  sync.Mutex
	handlers map[string]*handler
}

func ServerInstance() *serverImpl {
	serverInstanceOnce.Do(func() {
		serverInstance = &serverImpl{
			regLock:  sync.Mutex{},
			handlers: make(map[string]*handler, 8),
		}
	})
	return serverInstance
}

func (s *serverImpl) Serve(address string) error {
	l, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	go func() {
		for {
			c, err := l.Accept()
			if err != nil {
				logger.Error("accept error caused by %s", err)
			}
			go serveConn(s.handlers, c)
		}
	}()
	return nil
}

func (s *serverImpl) RegHandler(name string, h func(arg proto.Message) (res proto.Message), argId uint8) error {
	if len(name) > 0xff {
		return errors.New("name too lang, maxsize is 255")
	}
	if _, ok := MsgFactoryRegisterInstance().factoryMap[argId]; !ok {
		return errors.New("argId not registered")
	}
	s.regLock.Lock()
	s.handlers[name] = &handler{
		handleFunc: h,
		argId:      argId,
	}
	s.regLock.Unlock()
	return nil
}
