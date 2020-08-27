package raft

import "github.com/Mstch/naruto/raft/msg"

type candidateHandler struct{}

func newCandidateHandler() *candidateHandler {
	return &candidateHandler{}
}

func (f *candidateHandler) onVoteReq(req msg.VoteReq) {
	panic("implement me")
}

func (f *candidateHandler) onVoteResp(req msg.VoteResp) {
	panic("implement me")
}

func (f *candidateHandler) onHeartbeatReq(req msg.VoteReq) {
	panic("implement me")
}

func (f *candidateHandler) onHeartbeatResp(req msg.VoteReq) {
	panic("implement me")
}

func (f *candidateHandler) onAppendReq(req msg.VoteReq) {
	panic("implement me")
}

func (f *candidateHandler) onAppendResp(req msg.VoteReq) {
	panic("implement me")
}
