package raft

import (
	"github.com/Mstch/naruto/helper/db"
	"github.com/Mstch/naruto/helper/util"
	"sync"
)

type rule = uint32
type ruleSet = map[rule]struct{}

const (
	follower  rule = 0x001
	candidate      = 0x010
	leader         = 0x100
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

func changeRule(should ruleSet, new rule) error {
	ruleLock.Lock()
	if _, ok := should[nodeRule]; ok {
		nodeRule = new
	}
	ruleLock.Unlock()
	return nil
}

func becomeCandidate() {
}
