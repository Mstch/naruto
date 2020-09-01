package raft

import (
	"github.com/Mstch/naruto/raft/msg"
	"sync/atomic"
)

type leaderHandler struct{}

func (f *leaderHandler) onHeartbeat() {
	broadcast("Heartbeat", &msg.HeartbeatReq{
		Term:         atomic.LoadUint32(&nodeTerm),
		LeaderCommit: atomic.LoadUint64(&lastCommitIndex),
	})
}

func (f *leaderHandler) onHeartbeatResp(arg *msg.HeartbeatResp) {
	if arg.LastLogIndex < atomic.LoadUint64(&lastLogIndex) {
		go func(resp *msg.HeartbeatResp) {
			//TODO send append
		}(arg)
	}
}

func (f *leaderHandler) onAppendResp(arg *msg.AppendResp) {
	panic("implement me")
}

func (f *leaderHandler) OnAppendMajority() {

}
