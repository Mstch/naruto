package stupid

import (
	"bufio"
	"github.com/Mstch/naruto/helper/logger"
	"github.com/Mstch/naruto/raft/msg"
	"github.com/gogo/protobuf/proto"
	"log"
	"net"
	"sync"
	"sync/atomic"
	"testing"
)

var (
	serverInit = &sync.Once{}
	inited     int32
	gl         net.Listener
)

func TestClientImpl_Notify(t *testing.T) {
	c := &clientImpl{}
	err := c.Dial("localhost:1234")
	if err != nil {
		panic(err)
	}
	for i := 0; i < 10; i++ {
		err = c.Notify("fuck", &msg.HeartbeatReq{
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
	}, uint8(1))
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
			_ = arg.(*msg.HeartbeatReq)
			//logger.Info("req:%s", req.String())
			return nil
		}, uint8(1))
		go s.Listen(":1234")
	})
	c := &clientImpl{}
	err := c.Dial("localhost:1234")
	if err != nil {
		panic(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err = c.Notify("fuck", &msg.HeartbeatReq{
			Term:         uint32(i),
			LeaderCommit: uint64(i),
		})
		if err != nil {
			panic(err)
		}
	}
	log.Println("client test done")
}

func BenchmarkByte(t *testing.B) {
	serverInit.Do(func() {
		l, err := net.Listen("tcp", ":1234")
		if err != nil {
			panic(err)
		}
		gl = l
		atomic.StoreInt32(&inited, 1)
	})
	done := make(chan bool)
	for ; atomic.LoadInt32(&inited) != 1; {
	}
	t.ResetTimer()
	go func() {
		c, err := gl.Accept()
		if err != nil {
			panic(err)
		}
		r := bufio.NewReader(c)
		for i := 0; i < t.N; i++ {
			_, err := r.ReadByte()
			if err != nil {
				panic(err)
			}
		}
		done <- false
	}()
	cc, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		panic(err)
	}
	t.ResetTimer()
	w := bufio.NewWriter(cc)
	for i := 0; i < t.N; i++ {
		err = w.WriteByte(byte(i))
		if err != nil {
			panic(err)
		}
		err = w.Flush()
		if err != nil {
			panic(err)
		}
	}
	<-done
}
