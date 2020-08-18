package raft

import (
	"github.com/Mstch/naruto/helper/db"
	"github.com/Mstch/naruto/helper/util"
)

type rule = int32

const (
	follower rule = iota
	candidate
	leader
)

var key = []byte("node-state")

var (
	nodeRule        rule
	lastCommitIndex int32
	lastApplyIndex  int32
	term            int32
)

func loadFromDB() {
	_, err := db.Get(key, func(src []byte) (interface{}, error) {
		r, err := util.BytesToInt32(src[:3])
		if err != nil {
			return nil, err
		}
		lastCommit, err := util.BytesToInt32(src[4:7])
		if err != nil {
			return nil, err
		}
		lastApply, err := util.BytesToInt32(src[8:11])
		if err != nil {
			return nil, err
		}
		t, err := util.BytesToInt32(src[12:15])
		if err != nil {
			return nil, err
		}

		nodeRule = r
		lastCommitIndex = lastCommit
		lastApplyIndex = lastApply
		term = t
		return nil, nil
	})
	if err != nil {
		panic(err)
	}
}

func becomeCandidate() {
	
}
