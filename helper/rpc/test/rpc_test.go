package test

import (
	"github.com/Mstch/naruto/helper/rpc"
	"github.com/Mstch/naruto/helper/rpc/stupid"
	"github.com/gogo/protobuf/proto"
	"net"
	stdrpc "net/rpc"
	"sync"
	"testing"
)

var (
	stupidServerInit = &sync.Once{}
	stdServerInit    = &sync.Once{}
	testMsg          = &Msg{
		Content:   GenString(64),
		Content2:  GenString(128),
		Content3:  GenString(31),
		Content4:  GenString(17),
		Content5:  GenString(1024),
		Content6:  GenString(88),
		Content7:  GenString(62),
		Content8:  GenString(6),
		Content9:  GenString(5),
		Cont0Nt19: genStrings(18, 10),
	}
)

func genStrings(size int, ssize int) []string {
	ss := make([]string, ssize)
	for i := 0; i < ssize; i++ {
		ss[i] = GenString(size)
	}
	return ss
}

func GenString(size int) string {
	b := make([]byte, size)
	for i := 0; i < size; i++ {
		b[i] = byte(i)
	}
	return string(b)
}

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
	b.N = 100 * b.N
	stupid.UseBufPool = true
	stupidServerInit.Do(func() {
		register := rpc.DefaultRegister()
		register.RegMessageFactory(1, false, func() proto.Message {
			return &Msg{}
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
	for i := 0; i < 100; i++ {
		go notify(waiter, b.N/100)
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
	b.N = 100 * b.N
	stupid.UseBufPool = false
	stupid.UseRWPool = true
	stupidServerInit.Do(func() {
		register := rpc.DefaultRegister()
		register.RegMessageFactory(1, false, func() proto.Message {
			return &Msg{}
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
	for i := 0; i < 100; i++ {
		go notify(waiter, b.N/100)
	}
	waiter.Wait()
}
func BenchmarkStupidUnPooledRpc(b *testing.B) {
	b.N = 100 * b.N
	stupid.UseBufPool = false
	stupid.UseRWPool = false
	buf := make([]byte, 1024)
	for i, _ := range buf {
		buf[i] = 1
	}
	content := string(buf)
	stupidServerInit.Do(func() {
		register := rpc.DefaultRegister()
		register.RegMessageFactory(1, false, func() proto.Message {
			return &Msg{}
		})
		server := rpc.DefaultServer()
		err := server.Serve(":8739")
		if err != nil {
			panic(err)
		}
		err = server.RegHandler("Test", func(arg proto.Message) (res proto.Message) {
			return &Msg{Content: content}
		}, 1)
	})
	b.ResetTimer()
	waiter := &sync.WaitGroup{}
	waiter.Add(b.N)
	for i := 0; i < 100; i++ {
		go notify(waiter, b.N/100)
	}
	waiter.Wait()
}
func BenchmarkGoStdRpc(b *testing.B) {

	b.N = 100 * b.N
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
	for i := 0; i < 100; i++ {
		go func(n int) {
			c, err := stdrpc.Dial("tcp", "localhost:8888")
			if err != nil {
				panic(err)
			}
			for j := 0; j < n; j++ {
				respmsg := &Msg{}
				c.Go("TestStdRpc.Test", testMsg, respmsg, done)
			}
		}(b.N / 100)
	}
	for i := 0; i < b.N; i++ {
		<-done
	}
}

type TestStdRpc struct {
}

func (t *TestStdRpc) Test(req *Msg, resp *Msg) error {
	resp = testMsg
	return nil
}
