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

var (
	ruleLock        sync.Mutex
	nodeRule        rule
	voteFor         uint32
	id              uint32
	nodeTerm        uint32
	lastCommitIndex uint64
	lastApplyIndex  uint64
)
var nodeDB *db.DB

func StartStatDB() {
	var err error
	nodeDB, err = db.NewDB("node-data", &pebble.Options{})
	if err != nil {
		panic(err)
	}
	_, err = nodeDB.Get([]byte("nodeRule"), func(k, v []byte) (interface{}, error) {
		nodeRule = util.BytesToUInt32(v)
		return nil, nil
	})
	if err != nil {
		panic(err)
	}
	_, err = nodeDB.Get([]byte("voteFor"), func(k, v []byte) (interface{}, error) {
		voteFor = util.BytesToUInt32(v)
		return nil, nil
	})
	if err != nil {
		panic(err)
	}
	_, err = nodeDB.Get([]byte("id"), func(k, v []byte) (interface{}, error) {
		id = util.BytesToUInt32(v)
		return nil, nil
	})
	if err != nil {
		panic(err)
	}
	_, err = nodeDB.Get([]byte("nodeTerm"), func(k, v []byte) (interface{}, error) {
		return util.BytesToUInt32(v), nil
	})
	if err != nil {
		panic(err)
	}
	_, err = nodeDB.Get([]byte("lastCommitIndex"), func(k, v []byte) (interface{}, error) {
		lastCommitIndex = util.BytesToUInt64(v)
		return nil, nil
	})
	if err != nil {
		panic(err)
	}
	_, err = nodeDB.Get([]byte("lastApplyIndex"), func(k, v []byte) (interface{}, error) {
		lastApplyIndex = util.BytesToUInt64(v)
		return nil, nil
	})
	if err != nil {
		panic(err)
	}
}
