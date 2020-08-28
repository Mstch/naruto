package raft

import (
	"github.com/Mstch/naruto/raft/msg"
)

type followerHandler struct{

}

func (f *followerHandler) onElection() {
	panic("implement me")
}

func (f *followerHandler) onHeartbeat() {}

func (f *followerHandler) onVoteReq(arg *msg.VoteReq) *msg.VoteResp {
	panic("implement me")
}

func (f *followerHandler) onVoteResp(arg *msg.VoteResp) {
	panic("implement me")
}

func (f *followerHandler) onHeartbeatReq(arg *msg.HeartbeatReq) *msg.HeartbeatResp {
	panic("implement me")
}

func (f *followerHandler) onHeartbeatResp(arg *msg.HeartbeatResp) {
	panic("implement me")
}

func (f *followerHandler) onAppendReq(arg *msg.AppendReq) *msg.AppendResp {
	panic("implement me")
}

func (f *followerHandler) onAppendResp(arg *msg.AppendResp) {
	panic("implement me")
}

