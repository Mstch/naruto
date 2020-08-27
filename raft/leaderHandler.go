package raft

import "github.com/Mstch/naruto/raft/msg"

type leaderHandler struct{}

func newLeaderHandler() *leaderHandler {
	return &leaderHandler{}
}

func (f *leaderHandler) onVoteReq(req msg.VoteReq) {
	panic("implement me")
}

func (f *leaderHandler) onVoteResp(req msg.VoteResp) {
	panic("implement me")
}

func (f *leaderHandler) onHeartbeatReq(req msg.VoteReq) {
	panic("implement me")
}

func (f *leaderHandler) onHeartbeatResp(req msg.VoteReq) {
	panic("implement me")
}

func (f *leaderHandler) onAppendReq(req msg.VoteReq) {
	panic("implement me")
}

func (f *leaderHandler) onAppendResp(req msg.VoteReq) {
	panic("implement me")
}
