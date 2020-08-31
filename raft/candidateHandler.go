package raft

import (
	"github.com/Mstch/naruto/helper/logger"
	"github.com/Mstch/naruto/helper/quorum"
	"github.com/Mstch/naruto/raft/msg"
)

type candidateHandler struct{}

func (f *candidateHandler) onElection() {
	panic("implement me")
}

func (f *candidateHandler) onVoteResp(arg *msg.VoteResp) {
	logger.Info("receive vote resp as candidate")
	if arg.Grant {
		quorum.Approve(arg.Id)
	}
}

func (f *candidateHandler) onHeartbeatReq(arg *msg.HeartbeatReq) *msg.HeartbeatResp {
	panic("implement me")
}

func (f *candidateHandler) onAppendReq(arg *msg.AppendReq) *msg.AppendResp {
	panic("implement me")
}

func (f *candidateHandler) OnVoteMajority() {

}
