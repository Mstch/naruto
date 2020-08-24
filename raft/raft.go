package raft

import (
	"github.com/Mstch/naruto/helper/event"
	"github.com/Mstch/naruto/helper/timer"
)

var (
	electionTimerOption = &timer.RTimerOption{Name: "election-timer", Task: func() {
		event.Notify("election-timeout", nil)
	}, Min: 4000, Max: 5000}
	heartbeatTimerOption = &timer.RTimerOption{Name: "heartbeat-timer", Task: func() {
		event.Notify("heartbeat-timeout", nil)
	}, Min: 400, Max: 400}
)

func Startup() {

}
