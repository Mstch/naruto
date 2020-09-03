package domain

import "net"

type Member struct {
	Id    uint32
	Host  string
	Port  uint32
	Conn  net.Conn
	UPort uint32
	//ip:port to bind
	Address string
}