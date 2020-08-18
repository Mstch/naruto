package log

import (
	"github.com/Mstch/naruto/helper/kv"
)

type Log struct {
	Index int
	Term  int
	Cmd   kv.Cmd
}

