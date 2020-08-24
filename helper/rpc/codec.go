package rpc

type Encoder interface {
	Encode(src interface{}) []byte
}

type Decoder interface {
	Decode([]byte) interface{}
}

type Msg interface {
	Encoder
	Decoder
}
