package raft

import (
	"errors"
	"github.com/Mstch/naruto/helper/logger"
	"github.com/Mstch/naruto/helper/quorum"
	"github.com/Mstch/naruto/helper/timer"
	"github.com/Mstch/naruto/raft/msg"
	"sync/atomic"
)

var (
	ErrorNotLeader = errors.New("current node is not leader")
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

func (f *followerHandler) onHeartbeatReq(req *msg.HeartbeatReq) *msg.HeartbeatResp {
	if oldLeader := atomic.SwapUint32(&leaderId, req.From); oldLeader != req.From {
		logger.Info("new leader is %d", req.From)
	}
	timer.Loop(electionTimerOption)
	return &msg.HeartbeatResp{
		Id:           req.Id,
		From:         self.Id,
		Term:         atomic.LoadUint32(&nodeTerm),
		Success:      true,
		LastLogIndex: atomic.LoadUint64(&lastLogIndex),
	}
}

func (f *followerHandler) onAppendReq(req *msg.AppendReq) *msg.AppendResp {
	if oldLeader := atomic.SwapUint32(&nodeRule, req.From); oldLeader != req.From {
		logger.Info("new leader is %d", req.From)
	}
	timer.Loop(electionTimerOption)
	logs := req.Logs
	if req.PrevLogIndex == 0 {
		err := f.doAppend(req.Logs)
		if err != nil {
			return nil
		}
	}
	checkLog, err := getLog(req.PrevLogIndex)
	if err != nil {
		return nil
	}
	success := false
	if checkLog.Term == req.PrevLogTerm {
		err := f.doAppend(logs)
		if err != nil {
			return nil
		}
		success = true
	}
	return &msg.AppendResp{
		Id:           req.Id,
		From:         id,
		Term:         nodeTerm,
		Success:      success,
		LastLogIndex: atomic.LoadUint64(&lastLogIndex),
	}
}

func (f *followerHandler) doAppend(logs []*msg.Log) error {
	if len(logs) == 1 {
		_, _, err := appendOne(logs[0], true)
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

func (f *followerHandler) onCmd(cmd *msg.Cmd) *msg.CmdResp {
	resp := &msg.CmdResp{
		IsLeader:   false,
		LeaderAddr: memberManager.GetMembers()[leaderId].Address,
	}
	if isReadCmd(cmd) && cmd.ReadMode == msg.FollowerRead {
		if res, err := apply(cmd); err != nil {
			resp.Res = err.Error()
			resp.Success = false
			return resp
		} else {
			resp.Success = true
			resp.Res = res
			return resp
		}
	}
	return resp
}
