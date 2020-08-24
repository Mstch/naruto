package msg

import (
	"net"
	"net/rpc"
	"testing"
)

func Test(t *testing.T) {
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
