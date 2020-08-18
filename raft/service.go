package raft

type raftService interface {
	vote(req vote, resp voteResp)
	heartbeat()
	append()
	apply()
}

type FollowerService struct{
}
