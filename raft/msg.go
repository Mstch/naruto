package raft

import "github.com/Mstch/naruto/helper/log"

type vote struct {
	index int
	term  int
}

type heartbeat struct {
	index int
	term  int
}

type appendEntry struct {
	term      int
	logs      []log.Log
	prevIndex int
	prevTerm  int
}

type apply struct {
	index int
	term  int
}

type voteResp struct {
	term  int
	grant bool
}
