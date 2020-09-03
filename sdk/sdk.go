package sdk

import (
	"github.com/Mstch/naruto/helper/rpc"
	"github.com/Mstch/naruto/sdk/msg"
	"github.com/gogo/protobuf/proto"
	"log"
	"net"
	"sync"
)

type NarutoKV struct {
	ServersAddress []string
	Leader         string
	leaderClient   rpc.Client
	connectOnce    sync.Once
	clients        map[string]rpc.Client
}

func NewNarutoKV(servers []string) *NarutoKV {
	return &NarutoKV{
		ServersAddress: servers,
		clients:        make(map[string]rpc.Client, len(servers)),
	}
}

func (n *NarutoKV) Connect() {
	n.connectOnce.Do(func() {
		rpc.DefaultRegister().RegMessageFactory(1, true, func() proto.Message {
			return &msg.CmdResp{}
		})
		for _, addr := range n.ServersAddress {
			client := rpc.NewDefaultClient()
			conn, err := net.Dial("tcp", addr)
			if err != nil {
				continue
			}
			err = client.Conn(conn)
			if err != nil {
				continue
			}
			err = client.RegHandler("Cmd", n.handleCmdResp, 1)
			if err != nil {
				continue
			}
			n.clients[addr] = client
		}
	})
}

func (n *NarutoKV) SendCmd(cmd *msg.Cmd) {
}

func (n *NarutoKV) handleCmdResp(resp proto.Message) {
	cmdResp := resp.(*msg.CmdResp)
	if cmdResp.IsLeader == false {
		n.leaderClient = n.clients[cmdResp.LeaderAddr]
	}
	log.Printf(cmdResp.Res)
}

func (n *NarutoKV) probe() {
	for _, client := range n.clients {
		client.Notify("Cmd",&msg.Cmd{
		})
	}
}
