package raft

import (
	"github.com/Mstch/naruto/helper/db"
	"github.com/Mstch/naruto/helper/util"
	"sync"
)

type rule = uint32

const (
	follower rule = iota
	candidate
	leader
)

var key = []byte("node-state")

var (
	ruleLock        sync.Mutex
	nodeRule        rule
	voteFor         uint32
	id              uint32
	lastCommitIndex uint32
	lastApplyIndex  uint32
	term            uint32
)

func loadFromDB() {
	_, err := db.Get(key, func(src []byte) (interface{}, error) {
		nodeRule = util.BytesToUInt32(src[:4])
		lastCommitIndex = util.BytesToUInt32(src[4:8])
		lastApplyIndex = util.BytesToUInt32(src[8:12])
		term = util.BytesToUInt32(src[12:16])
		id = util.BytesToUInt32(src[16:20])
		return nil, nil
	})
	if err != nil {
		panic(err)
	}
}

func becomeCandidate() {
}
