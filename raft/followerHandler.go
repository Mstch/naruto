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
		Term:         atomic.AddUint32(&nodeTerm, 1),
		LastLogIndex: atomic.LoadUint64(&lastLogIndex),
		LastLogTerm:  atomic.LoadUint32(&lastLogTerm),
	})
}

func (f *followerHandler) onVoteReq(arg *msg.VoteReq) *msg.VoteResp {
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
		From:         self.Id,
		Term:         atomic.LoadUint32(&nodeTerm),
		Success:      true,
		LastLogIndex: atomic.LoadUint64(&lastLogIndex),
	}
}

func (f *followerHandler) onAppendReq(arg *msg.AppendReq) *msg.AppendResp {
	timer.Loop(electionTimerOption)
	logs := arg.Logs
	if arg.PrevLogIndex == 0 {
		err := f.doAppend(arg.Logs)
		if err != nil {
			return nil
		}
	}
	checkLog, err := getLog(arg.PrevLogIndex)
	if err != nil {
		return nil
	}
	success := false
	if checkLog.Term == arg.PrevLogTerm {
		err := f.doAppend(logs)
		if err != nil {
			return nil
		}
		success = true
	}
	return &msg.AppendResp{
		Id:           arg.Id,
		From:         id,
		Term:         nodeTerm,
		Success:      success,
		LastLogIndex: atomic.LoadUint64(&lastLogIndex),
	}
}

func (f *followerHandler) doAppend(logs []*msg.Log) error {
	if len(logs) == 1 {
		err := appendOne(logs[0], true)
		if err != nil {
			logger.Error("append one log failed,caused by %s", err)
			return err
		}
	} else {
		err := batchAppend(logs, true)
		if err != nil {
			logger.Error("batch append logs failed,caused by %s", err)
			return err
		}
	}
	return nil
}
