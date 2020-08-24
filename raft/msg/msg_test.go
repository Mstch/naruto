package msg

import (
	"log"
	"net"
	"net/rpc"
	"testing"
)

func Test(t *testing.T) {
	s := rpc.NewServer()
	lis, _ := net.Listen("tcp", ":1234")
	for {
		conn, err := lis.Accept()
		codec := &TestCodec{Conn: conn}
		if err != nil {
			log.Print("rpc.Serve: accept:", err.Error())
			return
		}
		go rpc.ServeCodec()
	}
}

type TestCodec struct {
	Conn net.Conn
}

func (t *TestCodec) ReadRequestHeader(request *rpc.Request) error {

}

func (t *TestCodec) ReadRequestBody(i interface{}) error {

}

func (t *TestCodec) WriteResponse(response *rpc.Response, i interface{}) error { panic("implement me") }

func (t *TestCodec) Close() error { panic("implement me") }
