package raft

import (
	"github.com/Mstch/naruto/helper/db"
	"github.com/Mstch/naruto/helper/util"
	"github.com/Mstch/naruto/raft/msg"
	"github.com/cockroachdb/pebble"
	"github.com/gogo/protobuf/proto"
	"sync/atomic"
)

var (
	logDB        *db.DB
	lastLogIndex uint64
)

func init() {
	var err error
	logDB, err = db.NewDB("raft-log", &pebble.Options{})
	if err != nil {
		panic(err)
	}
}
func appendOne(log msg.Log) error {
	buf := make([]byte, log.Size())
	err := log.Unmarshal(buf)
	if err != nil {
		return err
	}
	return logDB.Set(util.UInt64ToBytes(atomic.AddUint64(&lastLogIndex, 1)), buf)
}
func batchAppend(logs []msg.Log) error {
	l := len(logs)
	ks := make([][]byte, l)
	vs := make([][]byte, l)
	curIndex := atomic.AddUint64(&lastLogIndex, uint64(l)) - uint64(l)
	for i, log := range logs {
		curIndex++
		ks[i] = util.UInt64ToBytes(curIndex)
		vs[i] = make([]byte, log.Size())
		err := log.Unmarshal(vs[i])
		if err != nil {
			return err
		}
	}
	return logDB.BatchSet(ks, vs)
}
func commitTo(index uint64) {
	atomic.SwapUint64(&lastCommitIndex, index)
}

func applyTo(index uint64) error {
	start := util.UInt64ToBytes(atomic.LoadUint64(&lastApplyIndex))
	end := make([]byte, index)
	_, err := logDB.Iter(start, end, false, func(k, v []byte) (interface{}, error) {
		log := &msg.Log{}
		err := proto.Unmarshal(v, log)
		if err != nil {
			return nil, err
		}
		_, err = apply(log.Cmd)
		return nil, err
	})
	return err
}
