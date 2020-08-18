package raft

import "github.com/Mstch/naruto/helper/timer"

var (
	electionTimerOption  = timer.NewRTimerOption("election-timer", becomeCandidate, 4000, 5000)
	heartbeatTimerOption = timer.NewRTimerOption("heartbeat-timer", broadcastHeartbeat, 400, 400)
)

func Startup() {

}
