package raft

import (
	"github.com/Mstch/naruto/helper/logger"
	"github.com/Mstch/naruto/helper/member"
	"github.com/Mstch/naruto/helper/rpc"
	"github.com/Mstch/naruto/helper/rpc/stupid"
	"github.com/Mstch/naruto/raft/msg"
	"github.com/gogo/protobuf/proto"
	"sync/atomic"
)

type quorumType = uint8

var (
	majority uint32
	clients  = make([]rpc.Client, 0)
)

func regClientHandlers(client rpc.Client) {
	err := client.RegHandler("Vote", func(arg proto.Message) {
		if termInterceptor(arg.(*msg.VoteReq).Term) {
			if h, ok := clientHandlerDict[atomic.LoadUint32(&nodeRule)]["Vote"]; ok {
				h(arg)
			}
		}
	}, voteResp)
	if err != nil {
		panic(err)
	}
	err = client.RegHandler("Heartbeat", func(arg proto.Message) {
		if termInterceptor(arg.(*msg.HeartbeatReq).Term) {
			if h, ok := clientHandlerDict[atomic.LoadUint32(&nodeRule)]["Heartbeat"]; ok {
				h(arg)
			}
		}
	}, heartbeatResp)
	if err != nil {
		panic(err)
	}
	err = client.RegHandler("Append", func(arg proto.Message) {
		if termInterceptor(arg.(*msg.AppendReq).Term) {
			if h, ok := clientHandlerDict[atomic.LoadUint32(&nodeRule)]["Append"]; ok {
				h(arg)
			}
		}
	}, appendResp)
	if err != nil {
		panic(err)
	}
}

func StartupClient() {
	<-member.OK
	majority = uint32((len(member.ConnectedMembers) + 1) / 2)
	for _, m := range member.ConnectedMembers {
		client := stupid.NewClientImpl()
		err := client.Dial(m.Address)
		if err != nil {
			logger.Error("connect to %s failed, caused by:%s", m.Address, err)
		}
		clients = append(clients, client)
	}
	for _, client := range clients {
		regClientHandlers(client)
	}
}

func broadcast(name string, msg proto.Message) {
	for _, client := range clients {
		go func(c rpc.Client) {
			err := c.Notify(name, msg)
			if err != nil {
				logger.Error("send %s error,caused by:%s", name, err)
			}
		}(client)
	}
}
