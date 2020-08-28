package raft

import "github.com/Mstch/naruto/raft/msg"

type leaderHandler struct{}

func (f *leaderHandler) onElection() {
	panic("implement me")
}

func (f *leaderHandler) onHeartbeat() {
	panic("implement me")
}

func (f *leaderHandler) onVoteReq(arg *msg.VoteReq) *msg.VoteResp {
	panic("implement me")
}

func (f *leaderHandler) onVoteResp(arg *msg.VoteResp) {
	panic("implement me")
}

func (f *leaderHandler) onHeartbeatReq(arg *msg.HeartbeatReq) *msg.HeartbeatResp {
	panic("implement me")
}

func (f *leaderHandler) onHeartbeatResp(arg *msg.HeartbeatResp) {
	panic("implement me")
}

func (f *leaderHandler) onAppendReq(arg *msg.AppendReq) *msg.AppendResp {
	panic("implement me")
}

func (f *leaderHandler) onAppendResp(arg *msg.AppendResp) {
	panic("implement me")
}