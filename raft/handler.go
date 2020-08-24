package raft

import "github.com/Mstch/naruto/helper/event"

type handler interface {
	vote()
	heartbeat()
	append()
}

type followerHandler struct {
}

func registerHandlers() {
	event.RegHandler(&event.Handler{
		Handle: followerHandler{}.vote,
		Next:   nil,
	}, "follower-vote")
}

func (f *followerHandler) vote(req interface{}) error {
	return nil
}

func (f *followerHandler) heartbeat() {
}

func (f *followerHandler) append() {
	panic("implement me")
}

func (f *followerHandler) apply() {
	panic("implement me")
}
