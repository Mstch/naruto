package raft

import "github.com/Mstch/naruto/raft/msg"

type followerHandler struct{}

func newFollowerHandler() *followerHandler {
	return &followerHandler{}
}

func (f *followerHandler) onVoteReq(req msg.VoteReq) {
	panic("implement me")
}

func (f *followerHandler) onVoteResp(req msg.VoteResp) {
	panic("implement me")
}

func (f *followerHandler) onHeartbeatReq(req msg.VoteReq) {
	panic("implement me")
}

func (f *followerHandler) onHeartbeatResp(req msg.VoteReq) {
	panic("implement me")
}

func (f *followerHandler) onAppendReq(req msg.VoteReq) {
	panic("implement me")
}

func (f *followerHandler) onAppendResp(req msg.VoteReq) {
	panic("implement me")
}
