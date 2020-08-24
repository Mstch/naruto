package raft

import (
	"net/rpc"
	"testing"
)

type testCodec struct {
}

func (t testCodec) ReadRequestHeader(request *rpc.Request) error {

	return nil
}

func (t testCodec) ReadRequestBody(i interface{}) error {
	panic("implement me")
}

func (t testCodec) WriteResponse(response *rpc.Response, i interface{}) error {
	panic("implement me")
}

func (t testCodec) Close() error {
	panic("implement me")
}

func TestReg(t *testing.T) {
}
