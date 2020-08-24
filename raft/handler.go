package raft

type raftService interface {
	vote()
	heartbeat()
	append()
	apply()
}

type FollowerService struct{
}
