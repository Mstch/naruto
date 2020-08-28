package raft

import "github.com/Mstch/naruto/raft/msg"

type candidateHandler struct{}


func (f *candidateHandler) onElection() {
	panic("implement me")
}

func (f *candidateHandler) onHeartbeat() {
	panic("implement me")
}

func (f *candidateHandler) onVoteReq(arg *msg.VoteReq) *msg.VoteResp {
	panic("implement me")
}

func (f *candidateHandler) onVoteResp(arg *msg.VoteResp) {
	panic("implement me")
}

func (f *candidateHandler) onHeartbeatReq(arg *msg.HeartbeatReq) *msg.HeartbeatResp {
	panic("implement me")
}

func (f *candidateHandler) onHeartbeatResp(arg *msg.HeartbeatResp) {
	panic("implement me")
}

func (f *candidateHandler) onAppendReq(arg *msg.AppendReq) *msg.AppendResp {
	panic("implement me")
}

func (f *candidateHandler) onAppendResp(arg *msg.AppendResp) {
	panic("implement me")
}