package stupid

import (
	"github.com/Mstch/naruto/helper/logger"
	"github.com/Mstch/naruto/raft/msg"
	"github.com/gogo/protobuf/proto"
	"log"
	"sync"
	"testing"
)

var serverInit = &sync.Once{}

func TestClientImpl_Notify(t *testing.T) {
	c := &ClientImpl{}
	err := c.Dial("localhost:1234")
	if err != nil {
		panic(err)
	}
	for i := 0; i < 10; i++ {
		err = c.Notify("fuck", &msg.HeartbeatReq{
			Id:           uint32(i),
			Term:         uint32(i),
			LeaderCommit: uint64(i),
		})
		if err != nil {
			panic(err)
		}
	}
	log.Println("client test done")
	select {}
}
func TestServerImpl_Listen(t *testing.T) {
	s := DefaultServerInstance()
	f := DefaultRegisterInstance()
	f.RegMessageFactory(1, true, func() proto.Message {
		return &msg.HeartbeatReq{}
	})
	s.RegHandler("fuck", func(arg proto.Message) (res proto.Message) {
		req := arg.(*msg.HeartbeatReq)
		logger.Info("req:%s", req.String())
		return nil
	}, uint8(1), uint8(2))
	s.Listen(":1234")
}

func BenchmarkStupid(b *testing.B) {
	serverInit.Do(func() {
		s := DefaultServerInstance()
		f := DefaultRegisterInstance()
		f.RegMessageFactory(1, true, func() proto.Message {
			return &msg.HeartbeatReq{}
		})
		s.RegHandler("fuck", func(arg proto.Message) (res proto.Message) {
			req := arg.(*msg.HeartbeatReq)
			logger.Info("req:%s", req.String())
			return nil
		}, uint8(1), uint8(2))
		go s.Listen(":1234")
	})
	for i := 0; i < 100; i++ {
		go func() {
			c := &ClientImpl{}
			err := c.Dial("localhost:1234")
			if err != nil {
				panic(err)
			}
			for i := 0; i < b.N; i++ {
				err = c.Notify("fuck", &msg.HeartbeatReq{
					Id:           uint32(i),
					Term:         uint32(i),
					LeaderCommit: uint64(i),
				})
				if err != nil {
					panic(err)
				}
			}
		}()
	}
	log.Println("client test done")
}
