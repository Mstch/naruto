package raft

import (
	"github.com/Mstch/naruto/raft/msg"
	"sync/atomic"
)

type leaderHandler struct{}

func (l *leaderHandler) onHeartbeat() {
	broadcast("Heartbeat", &msg.HeartbeatReq{
		Term:         atomic.LoadUint32(&nodeTerm),
		LeaderCommit: atomic.LoadUint64(&lastCommitIndex),
	})
}

func (l *leaderHandler) onHeartbeatResp(arg *msg.HeartbeatResp) {
	if arg.LastLogIndex < atomic.LoadUint64(&lastLogIndex) {
		go func(resp *msg.HeartbeatResp) {
			//TODO send append
		}(arg)
	}
}



func (l *leaderHandler) onAppendResp(arg *msg.AppendResp) {
}

func (l *leaderHandler) OnAppendMajority(index uint64) {
	if index > atomic.LoadUint64(&lastCommitIndex) {
		applyTo(index)
	}
}
