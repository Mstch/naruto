package raft

import (
	"github.com/Mstch/naruto/helper/rpc"
	"github.com/Mstch/naruto/helper/rpc/stupid"
	"github.com/Mstch/naruto/raft/msg"
	"github.com/gogo/protobuf/proto"
)

const (
	voteReq uint8 = iota
	voteResp
	heartbeatReq
	heartbeatResp
	appendReq
	appendResp
)

type rpcHandler interface {
	onVoteReq(req msg.VoteReq) msg.VoteResp
	onVoteResp(req msg.VoteResp)
	onHeartbeatReq(req msg.HeartbeatReq) msg.HeartbeatResp
	onHeartbeatResp(req msg.HeartbeatResp)
	onAppendReq(req msg.AppendReq) msg.AppendResp
	onAppendResp(req msg.AppendResp)
}
type timeoutHandler interface {
	onElection()
	onHeartbeat()
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
func regHandlers(server rpc.Server) {
	f := &followerHandler{}
	c := &candidateHandler{}
	l := &leaderHandler{}
	dict := [3]map[string]func(req proto.Message) proto.Message{}
	dict[follower]["voteReq"] = f.onVoteReq
	dict[follower]["voteReq"] = f.onVoteReq
	dict[follower]["voteReq"] = f.onVoteReq
	dict[follower]["voteReq"] = f.onVoteReq
	server.RegHandler("voteReq", func(arg proto.Message) (res proto.Message) {

	}, voteReq, voteResp)
}
func handle() {
	server := stupid.DefaultServerInstance()
	register := stupid.DefaultRegisterInstance()
	regProtoMsg(register)
	regHandlers(server)
}
