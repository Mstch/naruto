package raft

import (
	"github.com/Mstch/naruto/helper/logger"
	"github.com/Mstch/naruto/helper/quorum"
	"github.com/Mstch/naruto/helper/timer"
	"github.com/Mstch/naruto/raft/msg"
	"sync/atomic"
)

type candidateHandler struct{}

func (f *candidateHandler) onElection() {
	timer.Loop(electionTimerOption)
}

func (f *candidateHandler) onVoteResp(arg *msg.VoteResp) {
	logger.Info("receive vote resp as candidate")
	if arg.Grant {
		quorum.Approve(arg.Id)
	}
}

func (f *candidateHandler) onHeartbeatReq(arg *msg.HeartbeatReq) *msg.HeartbeatResp {
	atomic.StoreUint32(&nodeRule, follower)
	return fh.onHeartbeatReq(arg)
}

func (f *candidateHandler) onAppendReq(arg *msg.AppendReq) *msg.AppendResp {
	atomic.StoreUint32(&nodeRule, follower)
	return fh.onAppendReq(arg)
}

func (f *candidateHandler) OnVoteMajority() {
	atomic.StoreUint32(&nodeRule, leader)
	broadcast("Heartbeat", &msg.HeartbeatReq{
		Term:         atomic.LoadUint32(&nodeTerm),
		LeaderCommit: atomic.LoadUint64(&lastCommitIndex),
	})
}
