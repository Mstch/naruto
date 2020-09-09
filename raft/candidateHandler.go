package raft

import (
	"github.com/Mstch/naruto/helper/logger"
	"github.com/Mstch/naruto/helper/quorum"
	"github.com/Mstch/naruto/helper/timer"
	"github.com/Mstch/naruto/raft/msg"
	"sync/atomic"
)

type candidateHandler struct{}

func (c *candidateHandler) onElection() {
	timer.Loop(electionTimerOption)
}

func (c *candidateHandler) onVoteResp(arg *msg.VoteResp) {
	if arg.Grant {
		quorum.Approve(arg.Id)
	}
}

func (c *candidateHandler) onHeartbeatReq(arg *msg.HeartbeatReq) *msg.HeartbeatResp {
	atomic.StoreUint32(&nodeRule, follower)
	return fh.onHeartbeatReq(arg)
}

func (c *candidateHandler) onAppendReq(arg *msg.AppendReq) *msg.AppendResp {
	atomic.StoreUint32(&nodeRule, follower)
	return fh.onAppendReq(arg)
}

func (c *candidateHandler) OnVoteMajority() {
	logger.Info("become leader")
	atomic.StoreUint32(&nodeRule, leader)
	timer.Loop(heartbeatTimerOption)
	broadcast("Heartbeat", &msg.HeartbeatReq{
		Term:         atomic.LoadUint32(&nodeTerm),
		LeaderCommit: atomic.LoadUint64(&lastCommitIndex),
	})
}
