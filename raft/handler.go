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
	onVoteReq(req msg.VoteReq)
	onVoteResp(req msg.VoteResp)
	onHeartbeatReq(req msg.VoteReq)
	onHeartbeatResp(req msg.VoteReq)
	onAppendReq(req msg.VoteReq)
	onAppendResp(req msg.VoteReq)
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
	ruleDict := [3]rpcHandler{}
	ruleDict[follower] = newFollowerHandler()
	ruleDict[candidate] = newCandidateHandler()
	ruleDict[leader] = newLeaderHandler()
}
func handle() {
	server := stupid.DefaultServerInstance()
	register := stupid.DefaultRegisterInstance()
	regProtoMsg(register)
}

