package member

import (
	"github.com/Mstch/naruto/helper/member/file"
	"net"
)

type Member struct {
	Id    uint32
	Host  string
	Port  uint32
	Conn  net.Conn
	UPort uint32
	//ip:port to bind
	Address string
}
type Members interface {
	//block until discover complete
	Discover()
	GetMembers() map[uint32]*Member
	Self() *Member
}

func Default() Members {
	return file.NewFileMembers()
}
