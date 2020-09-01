package raft

import (
	"github.com/Mstch/naruto/helper/logger"
	"github.com/Mstch/naruto/helper/member"
	"github.com/Mstch/naruto/helper/rpc"
	"github.com/Mstch/naruto/raft/msg"
	"github.com/gogo/protobuf/proto"
	"sync/atomic"
)

var (
	memberManager = member.Default()
	majority      uint32
	clientMap     = make(map[uint32]rpc.Client)
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
	members := memberManager.GetMembers()
	majority = uint32((len(members) + 1) / 2)
	for id, m := range members {
		client := rpc.NewDefaultClient()
		err := client.Conn(m.Conn)
		if err != nil {
			logger.Error("connect to %s failed, caused by:%s", m.Address, err)
		}
		clientMap[id] = client
	}
	for _, client := range clientMap {
		regClientHandlers(client)
	}
}

func broadcast(name string, msg proto.Message) {
	for _, client := range clientMap {
		go func(c rpc.Client) {
			err := c.Notify(name, msg)
			if err != nil {
				logger.Error("send %s error,caused by:%s", name, err)
			}
		}(client)
	}
}
func sendTo(to uint32, name string, msg proto.Message) {
	client := clientMap[to]
	err := client.Notify(name, msg)
	if err != nil {
		logger.Error("send to %d %s error,caused by:%s", to, name, err)
	}
}
