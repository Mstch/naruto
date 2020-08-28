package raft

import (
	"github.com/Mstch/naruto/helper/rpc"
	"github.com/Mstch/naruto/helper/rpc/stupid"
	"github.com/Mstch/naruto/raft/msg"
	"github.com/gogo/protobuf/proto"
	"sync/atomic"
)

const (
	voteReq uint8 = iota
	voteResp
	heartbeatReq
	heartbeatResp
	appendReq
	appendResp
)

var (
	serverHandlerDict = [3]map[string]func(req proto.Message) proto.Message{}
	clientHandlerDict = [3]map[string]func(req proto.Message){}
)

func init() {
	f := &followerHandler{}
	c := &candidateHandler{}
	l := &leaderHandler{}
	serverHandlerDict[follower] = make(map[string]func(req proto.Message) proto.Message, 3)
	serverHandlerDict[candidate] = make(map[string]func(req proto.Message) proto.Message, 2)
	serverHandlerDict[leader] = make(map[string]func(req proto.Message) proto.Message, 0)
	serverHandlerDict[follower]["Vote"] = func(req proto.Message) proto.Message {
		return f.onVoteReq(req.(*msg.VoteReq))
	}
	serverHandlerDict[follower]["Heartbeat"] = func(req proto.Message) proto.Message {
		return f.onHeartbeatReq(req.(*msg.HeartbeatReq))
	}
	serverHandlerDict[follower]["Append"] = func(req proto.Message) proto.Message {
		return f.onAppendReq(req.(*msg.AppendReq))
	}
	serverHandlerDict[candidate]["Heartbeat"] = func(req proto.Message) proto.Message {
		return c.onHeartbeatReq(req.(*msg.HeartbeatReq))
	}
	serverHandlerDict[candidate]["Append"] = func(req proto.Message) proto.Message {
		return c.onAppendReq(req.(*msg.AppendReq))
	}
	clientHandlerDict[follower] = make(map[string]func(req proto.Message), 0)
	clientHandlerDict[candidate] = make(map[string]func(req proto.Message), 1)
	clientHandlerDict[leader] = make(map[string]func(req proto.Message), 2)
	clientHandlerDict[candidate]["Vote"] = func(resp proto.Message) {
		c.onVoteResp(resp.(*msg.VoteResp))
	}
	clientHandlerDict[leader]["Heartbeat"] = func(resp proto.Message) {
		l.onHeartbeatResp(resp.(*msg.HeartbeatResp))
	}
	clientHandlerDict[leader]["Append"] = func(resp proto.Message) {
		l.onAppendResp(resp.(*msg.AppendResp))
	}
}

func regProtoMsg(register rpc.MessageFactoryRegister) {
	register.RegMessageFactory(voteReq, true, func() proto.Message {
		return &msg.VoteReq{}
	})
	register.RegMessageFactory(voteResp, true, func() proto.Message {
		return &msg.VoteReq{}
	})
	register.RegMessageFactory(heartbeatReq, true, func() proto.Message {
		return &msg.VoteReq{}
	})
	register.RegMessageFactory(heartbeatResp, true, func() proto.Message {
		return &msg.VoteReq{}
	})
	register.RegMessageFactory(appendReq, true, func() proto.Message {
		return &msg.VoteReq{}
	})
	register.RegMessageFactory(appendResp, true, func() proto.Message {
		return &msg.VoteReq{}
	})
}
func regServerHandlers(server rpc.Server) {
	err := server.RegHandler("Vote", func(arg proto.Message) (res proto.Message) {
		if termInterceptor(arg.(*msg.VoteReq).Term) {
			r := atomic.LoadUint32(&nodeRule)
			if h, ok := serverHandlerDict[r]["Vote"]; ok {
				return h(arg)
			}
			return nil
		} else {
			return &msg.VoteResp{
				Term:  atomic.LoadUint32(&nodeTerm),
				Grant: false,
			}
		}
	}, voteReq)
	if err != nil {
		panic(err)
	}
	err = server.RegHandler("Heartbeat", func(arg proto.Message) (res proto.Message) {
		if termInterceptor(arg.(*msg.HeartbeatReq).Term) {
			r := atomic.LoadUint32(&nodeRule)
			if h, ok := serverHandlerDict[r]["Heartbeat"]; ok {
				return h(arg)
			}
			return nil
		} else {
			return &msg.HeartbeatResp{
				Term:         atomic.LoadUint32(&nodeTerm),
				Success:      false,
				LastLogIndex: atomic.LoadUint64(&lastLogIndex),
			}
		}
	}, heartbeatReq)
	if err != nil {
		panic(err)
	}
	err = server.RegHandler("Append", func(arg proto.Message) (res proto.Message) {
		if termInterceptor(arg.(*msg.AppendReq).Term) {
			r := atomic.LoadUint32(&nodeRule)
			if h, ok := serverHandlerDict[r]["Append"]; ok {
				return h(arg)
			}
			return nil
		} else {
			return &msg.AppendResp{
				Term:         atomic.LoadUint32(&nodeTerm),
				Success:      false,
				LastLogIndex: atomic.LoadUint64(&lastLogIndex),
			}
		}
	}, appendReq)
	if err != nil {
		panic(err)
	}
}

func regClientHandlers(client rpc.Client) {
	err := client.RegHandler("Vote", func(arg proto.Message) {
		if termInterceptor(arg.(*msg.VoteReq).Term) {
			r := atomic.LoadUint32(&nodeRule)
			if h, ok := clientHandlerDict[r]["Vote"]; ok {
				h(arg)
			}
		}
	}, voteReq)
	if err != nil {
		panic(err)
	}
	err = client.RegHandler("Heartbeat", func(arg proto.Message) {
		if termInterceptor(arg.(*msg.HeartbeatReq).Term) {
			r := atomic.LoadUint32(&nodeRule)
			if h, ok := clientHandlerDict[r]["Heartbeat"]; ok {
				h(arg)
			}
		}
	}, heartbeatReq)
	if err != nil {
		panic(err)
	}
	err = client.RegHandler("Append", func(arg proto.Message) {
		if termInterceptor(arg.(*msg.AppendReq).Term) {
			r := atomic.LoadUint32(&nodeRule)
			if h, ok := clientHandlerDict[r]["Append"]; ok {
				h(arg)
			}
		}
	}, appendReq)
}

func startup(clients []rpc.Client) {
	server := stupid.DefaultServerInstance()
	register := stupid.DefaultRegisterInstance()
	regProtoMsg(register)
	regServerHandlers(server)
	for _, client := range clients {
		regClientHandlers(client)
	}
}

func termInterceptor(term uint32) bool {
	selfTerm := atomic.LoadUint32(&nodeTerm)
	if term > selfTerm {
		atomic.StoreUint32(&nodeTerm, term)
		becomeFollower()
		return true
	} else if term < selfTerm {
		return false
	}
	return true
}

func commitIndexInterceptor(commitIndex uint64) {
	selfAppliedIndex := atomic.LoadUint64(&lastApplyIndex)
	if commitIndex > selfAppliedIndex {

	}
}
