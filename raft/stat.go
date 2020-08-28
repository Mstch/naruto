package raft

import (
	"github.com/Mstch/naruto/helper/db"
	"github.com/Mstch/naruto/helper/util"
	"github.com/cockroachdb/pebble"
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
	lastCommitIndex uint64
	lastLogIndex    uint64
	lastApplyIndex  uint64
	nodeTerm        uint32
)
var nodeDB *db.DB

func init() {
	var err error
	nodeDB, err = db.NewDB("node-data", &pebble.Options{})
	if err != nil {
		panic(err)
	}
}
func loadFromDB() {
	_, err := nodeDB.Get(key, func(k, v []byte) (interface{}, error) {
		nodeRule = util.BytesToUInt32(v[:4])
		lastCommitIndex = util.BytesToUInt64(v[4:12])
		lastApplyIndex = util.BytesToUInt64(v[12:20])
		nodeTerm = util.BytesToUInt32(v[20:24])
		id = util.BytesToUInt32(v[24:28])
		voteFor = util.BytesToUInt32(v[28:32])
		return nil, nil
	})
	if err != nil {
		panic(err)
	}
}

func becomeCandidate() {
}

func becomeFollower() {
}
