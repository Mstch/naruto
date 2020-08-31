package raft

import "github.com/Mstch/naruto/raft/msg"

type leaderHandler struct{}


func (f *leaderHandler) onHeartbeat() {
	panic("implement me")
}



func (f *leaderHandler) onHeartbeatResp(arg *msg.HeartbeatResp) {
	panic("implement me")
}

func (f *leaderHandler) onAppendResp(arg *msg.AppendResp) {
	panic("implement me")
}