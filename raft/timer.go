package raft

import (
	"github.com/Mstch/naruto/helper/timer"
	"sync/atomic"
)

type timeoutType uint8

const (
	election timeoutType = iota
	heartbeat
)

var (
	electionTimerOption  *timer.RTimerOption
	heartbeatTimerOption *timer.RTimerOption
)

func StartupTimer() {
	electionTimerOption = &timer.RTimerOption{Name: "election-timer", Task: func() {
		dispatchTimeout(election)
	}, Min: 4000, Max: 5000}

	heartbeatTimerOption = &timer.RTimerOption{Name: "heartbeat-timer", Task: func() {
		dispatchTimeout(heartbeat)
	}, Min: 400, Max: 400}
	if nodeRule == follower || nodeRule == candidate {
		timer.Loop(electionTimerOption)
	} else if nodeRule == leader {
		timer.Loop(heartbeatTimerOption)
	}
}

func dispatchTimeout(tt timeoutType) {
	switch tt {
	case election:
		{
			r := atomic.LoadUint32(&nodeRule)
			if r == follower {
				fh.onElection()
			} else if r == candidate {
				ch.onElection()
			}
		}
	case heartbeat:
		{
			if atomic.LoadUint32(&nodeRule) == leader {
				lh.onHeartbeat()
			}
		}
	}
}
