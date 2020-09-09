package raft

import (
	"github.com/Mstch/naruto/conf"
	"github.com/Mstch/naruto/helper/db"
	"github.com/Mstch/naruto/helper/logger"
	"github.com/Mstch/naruto/helper/util"
	"github.com/cockroachdb/pebble"
	"strings"
)

type rule = uint32

const (
	follower rule = iota
	candidate
	leader
)

var (
	nodeRule        rule
	voteFor         uint32
	id              uint32
	nodeTerm        uint32
	lastCommitIndex uint64
	lastApplyIndex  uint64
	leaderId        uint32
	matchIndex      map[uint32]uint64
	leaseTimeout    int64
)
var statDB *db.DB

func StartupStatDB() {
	var err error
	id = conf.Conf.Id
	statDB, err = db.NewDB("stat-data", &pebble.Options{ErrorIfNotExists: true})
	if err != nil {
		if strings.HasSuffix(err.Error(), "no such file or directory") {
			logger.Info("stat-data db not exist,init node stat as follower first at startup")
			initStatDB()
			return
		} else {
			panic(err)
		}
	}
	_, err = statDB.Get([]byte("nodeRule"), func(k, v []byte) (interface{}, error) {
		nodeRule = util.BytesToUInt32(v)
		return nil, nil
	})
	if err != nil {
		panic(err)
	}
	_, err = statDB.Get([]byte("voteFor"), func(k, v []byte) (interface{}, error) {
		voteFor = util.BytesToUInt32(v)
		return nil, nil
	})
	if err != nil {
		panic(err)
	}
	_, err = statDB.Get([]byte("nodeTerm"), func(k, v []byte) (interface{}, error) {
		return util.BytesToUInt32(v), nil
	})
	if err != nil {
		panic(err)
	}
	_, err = statDB.Get([]byte("lastCommitIndex"), func(k, v []byte) (interface{}, error) {
		lastCommitIndex = util.BytesToUInt64(v)
		return nil, nil
	})
	if err != nil {
		panic(err)
	}
	_, err = statDB.Get([]byte("lastApplyIndex"), func(k, v []byte) (interface{}, error) {
		lastApplyIndex = util.BytesToUInt64(v)
		return nil, nil
	})
	if err != nil {
		panic(err)
	}
	_, err = statDB.Get([]byte("leaderId"), func(k, v []byte) (interface{}, error) {
		leaderId = util.BytesToUInt32(v)
		return nil, nil
	})
	if err != nil {
		panic(err)
	}
}

func initStatDB() {
	var err error
	statDB, err = db.NewDB("stat-data", &pebble.Options{})
	if err != nil {
		panic(err)
	}
	err = statDB.Set([]byte("nodeRule"), util.UInt32ToBytes(follower))
	if err != nil {
		panic(err)
	}
	err = statDB.Set([]byte("voteFor"), util.UInt32ToBytes(0))
	if err != nil {
		panic(err)
	}
	err = statDB.Set([]byte("nodeTerm"), util.UInt32ToBytes(0))
	if err != nil {
		panic(err)
	}
	err = statDB.Set([]byte("lastCommitIndex"), util.UInt64ToBytes(0))
	if err != nil {
		panic(err)
	}
	err = statDB.Set([]byte("lastApplyIndex"), util.UInt64ToBytes(0))
	if err != nil {
		panic(err)
	}
	err = statDB.Set([]byte("leaderId"), util.UInt64ToBytes(0))
	if err != nil {
		panic(err)
	}
	nodeRule = follower
	voteFor = 0
	nodeTerm = 0
	lastCommitIndex = 0
	lastApplyIndex = 0
	leaderId = 0
}
