package raft

import (
	"github.com/Mstch/naruto/helper/logger"
	"github.com/Mstch/naruto/helper/quorum"
	"github.com/Mstch/naruto/helper/timer"
	"github.com/Mstch/naruto/raft/msg"
	"sync/atomic"
)

type followerHandler struct {
}

func (f *followerHandler) onElection() {
	logger.Info("election timeout as follower")
	atomic.StoreUint32(&nodeRule, candidate)
	broadcast("Vote", &msg.VoteReq{
		Id: quorum.RegQuorum(majority, func() {
			if atomic.LoadUint32(&nodeRule) == candidate {
				ch.OnVoteMajority()
			}
		}),
		Term:         atomic.LoadUint32(&nodeTerm),
		LastLogIndex: atomic.LoadUint64(&lastLogIndex),
		LastLogTerm:  atomic.LoadUint32(&lastLogTerm),
	})
}

func (f *followerHandler) onVoteReq(arg *msg.VoteReq) *msg.VoteResp {
	logger.Info("receive vote req as follower")
	if (atomic.LoadUint32(&voteFor) == 0 || atomic.LoadUint32(&voteFor) == arg.From) &&
		arg.GetTerm() >= atomic.LoadUint32(&lastLogTerm) &&
		arg.LastLogIndex >= atomic.LoadUint64(&lastLogIndex) {
		return &msg.VoteResp{
			Id:    arg.Id,
			Term:  atomic.LoadUint32(&nodeTerm),
			Grant: true,
		}
	}
	return &msg.VoteResp{
		Id:    arg.Id,
		Term:  atomic.LoadUint32(&nodeTerm),
		Grant: false,
	}
}

func (f *followerHandler) onHeartbeatReq(_ *msg.HeartbeatReq) *msg.HeartbeatResp {
	timer.Loop(electionTimerOption)
	return &msg.HeartbeatResp{
		Term:         atomic.LoadUint32(&nodeTerm),
		Success:      true,
		LastLogIndex: atomic.LoadUint64(&lastLogIndex),
	}
}

func (f *followerHandler) onAppendReq(arg *msg.AppendReq) *msg.AppendResp {
	panic("implement me")
}
