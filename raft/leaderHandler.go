package raft

import (
	"github.com/Mstch/naruto/helper/quorum"
	"github.com/Mstch/naruto/raft/msg"
	"sync/atomic"
	"time"
)

type leaderHandler struct{}

func (l *leaderHandler) onHeartbeat() {
	broadcast("Heartbeat", &msg.HeartbeatReq{
		Term:         atomic.LoadUint32(&nodeTerm),
		LeaderCommit: atomic.LoadUint64(&lastCommitIndex),
	})
}

func (l *leaderHandler) onHeartbeatResp(arg *msg.HeartbeatResp) {
	if arg.Id != 0 {
		quorum.Approve(arg.Id)
	}
	leaseTimeout = time.Now().Add(4 * time.Second).UnixNano()
	if arg.LastLogIndex < atomic.LoadUint64(&lastLogIndex) {
		go func(resp *msg.HeartbeatResp) {
			//TODO send append
		}(arg)
	}
}

func (l *leaderHandler) onAppendResp(arg *msg.AppendResp) {
	leaseTimeout = time.Now().Add(4 * time.Second).UnixNano()
}

func (l *leaderHandler) onLeaseRead(cmd *msg.Cmd) *msg.CmdResp {
	if time.Now().UnixNano() > leaseTimeout {
		return l.onReadIndex(cmd)
	} else {
		if res, err := apply(cmd); err != nil {
			return &msg.CmdResp{
				Res:      err.Error(),
				Success:  false,
				IsLeader: true,
			}
		} else {
			return &msg.CmdResp{
				Res:      res,
				Success:  true,
				IsLeader: true,
			}
		}
	}
}

func (l *leaderHandler) onReadIndex(cmd *msg.Cmd) *msg.CmdResp {
	quorumId := quorum.RegQuorum(majority, func() {})
	broadcast("Heartbeat", &msg.HeartbeatReq{
		Id:           quorumId,
		From:         self.Id,
		Term:         atomic.LoadUint32(&nodeTerm),
		LeaderCommit: atomic.LoadUint64(&lastCommitIndex),
	})
	quorum.Wait(quorumId)
	if res, err := apply(cmd); err != nil {
		return &msg.CmdResp{
			Res:      err.Error(),
			Success:  false,
			IsLeader: true,
		}
	} else {
		return &msg.CmdResp{
			Res:      res,
			Success:  true,
			IsLeader: true,
		}
	}
}
