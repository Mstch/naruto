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
	"time"
	_ "unsafe"
)

var (
	stupidServerInit = &sync.Once{}
	stdServerInit    = &sync.Once{}
	testMsg          = make([]*msg.AppendReq, 10)
	Delay            time.Duration
)

func init() {
	for i := range testMsg {
		m := &msg.AppendReq{}
		m.LeaderCommit = uint64(i)
		m.PrevLogTerm = uint32(i)
		m.Term = uint32(i)
		for j := 0; j < i; j++ {
			m.Logs = append(m.Logs, &msg.Log{
				Cmd: &msg.Cmd{
					Opt:      msg.Get,
					ReadMode: msg.Lease,
					Key:      "Set",
					Value:    "v",
				},
				Term:  uint32(i),
				Index: uint64(i),
			})
		}
	}
}

// Uint32 returns a lock free uint32 value.
//go:linkname fastrandn runtime.fastrandn
func fastrandn(n uint32) uint32

func notify(w *sync.WaitGroup, n int, b *testing.B) {
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
		if Delay > 0 {
			time.Sleep(Delay)
		}
		err := client.Notify("Test", testMsg[int(fastrandn(10))])
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkBufPool200StupidRpc(b *testing.B) {
	Delay = 200 * time.Microsecond
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
			return testMsg[int(fastrandn(10))]
		}, 1)
	})
	b.ResetTimer()
	waiter := &sync.WaitGroup{}
	waiter.Add(b.N)
	div := b.N / 100
	mod := b.N - div*100
	if div < 100 {
		for i := 0; i < div; i++ {
			go notify(waiter, 100, b)
		}
	} else {
		for i := 0; i < 100; i++ {
			go notify(waiter, div, b)
		}
	}
	if mod > 0 {
		go notify(waiter, mod, b)
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
func BenchmarkBufPool500StupidRpc(b *testing.B) {
	Delay = 500 * time.Microsecond
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
			return testMsg[int(fastrandn(10))]
		}, 1)
	})
	b.ResetTimer()
	waiter := &sync.WaitGroup{}
	waiter.Add(b.N)
	div := b.N / 100
	mod := b.N - div*100
	if div < 100 {
		for i := 0; i < div; i++ {
			go notify(waiter, 100, b)
		}
	} else {
		for i := 0; i < 100; i++ {
			go notify(waiter, div, b)
		}
	}
	if mod > 0 {
		go notify(waiter, mod, b)
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
	done := &sync.WaitGroup{}
	done.Add(b.N)
	div := b.N / 100
	mod := b.N - div*100
	if div < 100 {
		for i := 0; i < div; i++ {
			go stdCall(done, 100)
		}
	} else {
		for i := 0; i < 100; i++ {
			go stdCall(done, div)
		}
	}
	if mod > 0 {
		go stdCall(done, mod)
	}
	done.Wait()

}
func stdNotify(done chan *stdrpc.Call, n int) {
	c, err := stdrpc.Dial("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	for j := 0; j < n; j++ {
		respmsg := &msg.AppendReq{}
		c.Go("TestStdRpc.Test", testMsg[int(fastrandn(10))], respmsg, done)
		err = c.Call("TestStdRpc.Test", testMsg[int(fastrandn(10))], respmsg)
		if err != nil {
			panic(err)
		}
	}
}
func stdCall(waiter *sync.WaitGroup, n int) {
	c, err := stdrpc.Dial("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	for j := 0; j < n; j++ {
		if Delay > 0 {
			time.Sleep(Delay)
		}
		respmsg := &msg.AppendReq{}
		err = c.Call("TestStdRpc.Test", testMsg[int(fastrandn(10))], respmsg)
		if err != nil {
			panic(err)
		}
		waiter.Done()
	}
}

type TestStdRpc struct {
}

func (t *TestStdRpc) Test(req *msg.AppendReq, resp *msg.AppendReq) error {
	resp.Id = testMsg[int(fastrandn(10))].Id
	resp.LeaderCommit = testMsg[int(fastrandn(10))].LeaderCommit
	resp.PrevLogIndex = testMsg[int(fastrandn(10))].PrevLogIndex
	copy(resp.Logs, testMsg[int(fastrandn(10))].Logs)
	resp.PrevLogTerm = testMsg[int(fastrandn(10))].PrevLogTerm
	return nil
}
