package stupid

import (
	"errors"
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

func DefaultServerInstance() *serverImpl {
	serverInstanceOnce.Do(func() {
		serverInstance = &serverImpl{
			regLock:  sync.Mutex{},
			handlers: make(map[string]*handler, 8),
		}
	})
	return serverInstance
}

func (s *serverImpl) Listen(address string) {
	l, err := net.Listen("tcp", address)
	defer l.Close()
	if err != nil {
		panic(err)
	}
	for {
		c, _ := l.Accept()
		go serveConn(s.handlers, c)
	}
}

func (s *serverImpl) RegHandler(name string, h func(arg proto.Message) (res proto.Message), argId, resId uint8) error {
	if len(name) > 0xff {
		return errors.New("name too lang, maxsize is 255")
	}
	s.regLock.Lock()
	s.handlers[name] = &handler{
		handleFunc: h,
		argId:      argId,
		resId:      resId,
	}
	s.regLock.Unlock()
	return nil
}
