package rpc

type Manager interface {
	RegHandler(handler func(req interface{}) (resp interface{}))
	RegAsyncHandler(handler func(req interface{}, respC chan interface{}))
}


