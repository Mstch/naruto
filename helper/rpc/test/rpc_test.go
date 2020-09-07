package test

import (
	"github.com/Mstch/naruto/helper/rpc"
	"github.com/Mstch/naruto/helper/rpc/stupid"
	"github.com/Mstch/naruto/raft/msg"
	"github.com/gogo/protobuf/proto"
	"net"
	stdrpc "net/rpc"
	"sync"
	"testing"
)

var (
	stupidServerInit = &sync.Once{}
	stdServerInit    = &sync.Once{}
	testMsg          = &msg.AppendReq{
		Id:           1,
		Term:         2,
		PrevLogIndex: 3,
		PrevLogTerm:  4,
		LeaderCommit: 5,
		Logs: []*msg.Log{{
			Term:  6,
			Index: 7,
			Cmd: &msg.Cmd{
				Opt:   msg.Get,
				Key:   "get",
				Value: "0",
			},
		}},
	}
)

func notify(w *sync.WaitGroup, n int) {
	client := rpc.NewDefaultClient()
	conn, err := net.Dial("tcp", "localhost:8739")
	if err != nil {
		panic(err)
	}
	err = client.Conn(conn)
	if err != nil {
		panic(err)
	}
	client.RegHandler("Test", func(arg proto.Message) {
		w.Done()
	}, 1)
	for i := 0; i < n; i++ {
		err := client.Notify("Test", testMsg)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkBufPoolStupidRpc(b *testing.B) {
	stupid.UseBufPool = true
	stupidServerInit.Do(func() {
		register := rpc.DefaultRegister()
		register.RegMessageFactory(1, true, func() proto.Message {
			return &msg.AppendReq{}
		})
		server := rpc.DefaultServer()
		err := server.Serve(":8739")
		if err != nil {
			panic(err)
		}
		err = server.RegHandler("Test", func(arg proto.Message) (res proto.Message) {
			return testMsg
		}, 1)
	})
	b.ResetTimer()
	waiter := &sync.WaitGroup{}
	waiter.Add(b.N)
	div := b.N / 100
	mod := b.N - div*100
	if div < 100 {
		for i := 0; i < div; i++ {
			go notify(waiter, 100)
		}
	} else {
		for i := 0; i < 100; i++ {
			go notify(waiter, div)
		}
	}
	if mod > 0 {
		go notify(waiter, mod)
	}
	waiter.Wait()
	//b.StopTimer()
	//min := math.MaxInt32
	//count := 0
	//for {
	//	sb := stupid.BufPool.Get().(*sbuf.Buffer)
	//	if sb.Size() == 0 {
	//		break
	//	}
	//	if sb.Size() < min {
	//		min = sb.Size()
	//	}
	//	count++
	//}
	//log.Println("count:", count, "min", min)
}
func BenchmarkStupidRWPooledRpc(b *testing.B) {

	stupid.UseBufPool = false
	stupid.UseRWPool = true
	stupidServerInit.Do(func() {
		register := rpc.DefaultRegister()
		register.RegMessageFactory(1, false, func() proto.Message {
			return &msg.AppendReq{}
		})
		server := rpc.DefaultServer()
		err := server.Serve(":8739")
		if err != nil {
			panic(err)
		}
		err = server.RegHandler("Test", func(arg proto.Message) (res proto.Message) {
			return testMsg
		}, 1)
	})
	b.ResetTimer()
	waiter := &sync.WaitGroup{}
	waiter.Add(b.N)
	div := b.N / 100
	mod := b.N - div*100
	if div < 100 {
		for i := 0; i < div; i++ {
			go notify(waiter, 100)
		}
	} else {
		for i := 0; i < 100; i++ {
			go notify(waiter, div)
		}
	}
	if mod > 0 {
		go notify(waiter, mod)
	}
	waiter.Wait()
}
func BenchmarkStupidUnPooledRpc(b *testing.B) {

	stupid.UseBufPool = false
	stupid.UseRWPool = false
	stupidServerInit.Do(func() {
		register := rpc.DefaultRegister()
		register.RegMessageFactory(1, false, func() proto.Message {
			return &msg.AppendReq{}
		})
		server := rpc.DefaultServer()
		err := server.Serve(":8739")
		if err != nil {
			panic(err)
		}
		err = server.RegHandler("Test", func(arg proto.Message) (res proto.Message) {
			return testMsg
		}, 1)
	})
	b.ResetTimer()
	waiter := &sync.WaitGroup{}
	waiter.Add(b.N)
	div := b.N / 100
	mod := b.N - div*100
	if div < 100 {
		for i := 0; i < div; i++ {
			go notify(waiter, 100)
		}
	} else {
		for i := 0; i < 100; i++ {
			go notify(waiter, div)
		}
	}
	if mod > 0 {
		go notify(waiter, mod)
	}
	waiter.Wait()
}
func BenchmarkGoStdRpc(b *testing.B) {

	stdServerInit.Do(func() {
		s := stdrpc.NewServer()
		err := s.Register(&TestStdRpc{})
		if err != nil {
			panic(err)
		}
		l, err := net.Listen("tcp", ":8888")
		if err != nil {
			panic(err)
		}
		go s.Accept(l)
	})
	b.ResetTimer()
	done := make(chan *stdrpc.Call, b.N)

	div := b.N / 100
	mod := b.N - div*100
	if div < 100 {
		for i := 0; i < div; i++ {
			go stdNotify(done, 100)
		}
	} else {
		for i := 0; i < 100; i++ {
			go stdNotify(done, div)
		}
	}
	if mod > 0 {
		go stdNotify(done, mod)
	}

	for i := 0; i < b.N; i++ {
		<-done
	}
}
func stdNotify(done chan *stdrpc.Call, n int) {
	c, err := stdrpc.Dial("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	for j := 0; j < n; j++ {
		respmsg := &msg.AppendReq{}
		c.Go("TestStdRpc.Test", testMsg, respmsg, done)
	}
}

type TestStdRpc struct {
}

func (t *TestStdRpc) Test(req *msg.AppendReq, resp *msg.AppendReq) error {
	resp = testMsg
	return nil
}
