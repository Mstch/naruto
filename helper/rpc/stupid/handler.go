package stupid

import "github.com/gogo/protobuf/proto"

type handler struct {
	handleFunc func(arg proto.Message) (res proto.Message)
	argId      uint8
}
