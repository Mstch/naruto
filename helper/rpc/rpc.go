package rpc

import (
	"github.com/Mstch/naruto/helper/rpc/stupid"
	"github.com/gogo/protobuf/proto"
	"net"
)

/**
序列化逻辑:
1. 把入参的message序列化为[]byte
2. 与name,seq封装为stupidmsg,再序列化
*/

type Server interface {
	//not block
	Serve(address string) error
	RegHandler(name string, handler func(arg proto.Message) (res proto.Message), argId uint8) error
}

type Client interface {
	//not block
	Conn(conn net.Conn) error
	Call(name string, arg proto.Message) (res proto.Message, err error)
	AsyncCall(name string, arg proto.Message) (resC chan proto.Message, err error)
	RegHandler(name string, handler func(arg proto.Message), argId uint8) error
	Notify(name string, arg proto.Message) error
	RemoteAddr() net.Addr
}

type MessageFactoryRegister interface {
	RegMessageFactory(id uint8, usePool bool, factory func() proto.Message)
}

func DefaultServer() Server {
	return stupid.ServerInstance()
}
func NewDefaultClient() Client {
	return stupid.NewClientImpl()
}
func DefaultRegister() MessageFactoryRegister {
	return stupid.MsgFactoryRegisterInstance()
}
