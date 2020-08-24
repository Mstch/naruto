package helper

import "sync"

type Helper interface {
	Meta() sync.Map
	Startup() chan bool
}
