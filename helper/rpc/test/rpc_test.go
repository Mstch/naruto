package test

import (
	"github.com/Mstch/naruto/helper/rpc"
	"github.com/Mstch/naruto/helper/rpc/stupid"
	"github.com/Mstch/naruto/helper/sbuf"
	"github.com/gogo/protobuf/proto"
	"log"
	"math"
	"net"
	stdrpc "net/rpc"
	"sync"
	"testing"
)

var (
	stupidServerInit = &sync.Once{}
	stdServerInit    = &sync.Once{}
)

func notify(w *sync.WaitGroup, content string, n int) {
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
		err := client.Notify("Test", &Msg{
			Content: content,
		})
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkBufPoolStupidRpc(b *testing.B) {
	b.N = 100 * b.N
	stupid.UseBufPool = true
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
		go notify(waiter, content, b.N/100)
	}
	waiter.Wait()
	b.StopTimer()
	min := math.MaxInt32
	count := 0
	for {
		sb := stupid.BufPool.Get().(*sbuf.Buffer)
		if sb.Cap() == 0 {
			break
		}
		if sb.Cap() < min {
			min = sb.Cap()
		}
		count++
	}
	log.Println("count:", count, "min", min)
}
func BenchmarkStupidRWPooledRpc(b *testing.B) {
	b.N = 100 * b.N
	stupid.UseBufPool = false
	stupid.UseRWPool = true
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
		go notify(waiter, content, b.N/100)
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
		go notify(waiter, content, b.N/100)
	}
	waiter.Wait()
}
func BenchmarkGoStdRpc(b *testing.B) {
	b.N = 100 * b.N
	buf := make([]byte, 1024)
	for i, _ := range buf {
		buf[i] = 1
	}
	content := string(buf)
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
				c.Go("TestStdRpc.Test", &Msg{Content: content}, &Msg{}, done)
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
	resp.Content = req.Content
	return nil
}
