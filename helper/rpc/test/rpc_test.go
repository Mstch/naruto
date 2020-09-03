package test

import (
	"bytes"
	"github.com/Mstch/naruto/helper/rpc"
	"github.com/gogo/protobuf/proto"
	"net"
	"net/http"
	"sync"
	"testing"
)

var (
	stupidServerInit = &sync.Once{}
	httpServerInit   = &sync.Once{}
)

func BenchmarkStupidRpc(b *testing.B) {
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
	client := rpc.NewDefaultClient()
	conn, err := net.Dial("tcp", "localhost:8739")
	if err != nil {
		panic(err)
	}
	err = client.Conn(conn)
	if err != nil {
		panic(err)
	}
	waiter := &sync.WaitGroup{}
	waiter.Add(b.N)
	client.RegHandler("Test", func(arg proto.Message) {
		waiter.Done()
	}, 1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := client.Notify("Test", &Msg{
			Content: content,
		})
		if err != nil {
			panic(err)
		}
	}
	waiter.Wait()
}
func TestStupidRpc(t *testing.T) {

	buf := make([]byte, 1024)
	for i, _ := range buf {
		buf[i] = 1
	}
	content := string(buf)
	register := rpc.DefaultRegister()
	register.RegMessageFactory(1, false, func() proto.Message {
		return &Msg{}
	})
	server := rpc.DefaultServer()
	err := server.Serve(":8888")
	if err != nil {
		panic(err)
	}
	err = server.RegHandler("Test", func(arg proto.Message) (res proto.Message) {
		println("server receive success ")
		return &Msg{Content: content}
	}, 1)
	client := rpc.NewDefaultClient()
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	err = client.Conn(conn)
	if err != nil {
		panic(err)
	}
	done := make(chan bool)
	err = client.RegHandler("Test", func(arg proto.Message) {
		println("receive success ")
		done <- false
	}, 1)
	if err != nil {
		panic(err)
	}
	err = client.Notify("Test", &Msg{
		Content: content,
	})
	println("notify success ")
	if err != nil {
		panic(err)
	}
	<-done
}
func BenchmarkHttpRpc(b *testing.B) {
	buf := make([]byte, 1024)
	for i, _ := range buf {
		buf[i] = 1
	}
	httpServerInit.Do(func() {
		http.HandleFunc("/Test", func(writer http.ResponseWriter, request *http.Request) {
			b := make([]byte, 1024)
			_, err := request.Body.Read(b)
			if err != nil {
				panic(err)
			}
			_, err = writer.Write(buf)
			if err != nil {
				panic(err)
			}
		})
		go func() {
			_ = http.ListenAndServe(":1234", nil)
		}()
	})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req := bytes.NewBuffer(buf)
		resp, err := http.Post("http://localhost:1234", "text/plain", req)
		if err != nil {
			panic(err)
		}
		b := make([]byte, 1024)
		_, _ = resp.Body.Read(b)
	}
}
