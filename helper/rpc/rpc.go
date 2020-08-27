package rpc

import (
	"github.com/gogo/protobuf/proto"
)

/**
序列化逻辑:
1. 把入参的message序列化为[]byte
2. 与name,seq封装为stupidmsg,再序列化
*/

type Server interface {
	Listen(address string) error
	RegHandler(name string, handler func(arg *proto.Message) (res proto.Message), argId, resId uint8) error
}

type Client interface {
	Dial(address string) error
	Call(name string, arg proto.Message) (res proto.Message, err error)
	AsyncCall(name string, arg proto.Message) (resC chan proto.Message, err error)
	RegHandler(name string, handler func(arg *proto.Message) (res proto.Message), argId, resId uint8) error
	Notify(name string, arg proto.Message) error
}

type MessageFactoryRegister interface {
	RegMessageFactory(id uint8, usePool bool, factory func() proto.Message)
}

type Rpc interface {
	Server
	Client
	MessageFactoryRegister
}
