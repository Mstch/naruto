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
	generatedKey uint64
	lastLogIndex uint64
	lastLogTerm  uint32
)

func StartLogDB() {
	var err error
	logDB, err = db.NewDB("raft-log", &pebble.Options{})
	if err != nil {
		panic(err)
	}
	_, _, err = logDB.Last(func(k, v []byte) (interface{}, error) {
		lastLogIndex = util.BytesToUInt64(k)
		lastLog := &msg.Log{}
		err := proto.Unmarshal(v, lastLog)
		if err != nil {
			return nil, err
		}
		lastLogTerm = lastLog.GetTerm()
		return nil, nil
	})
	if err != nil {
		panic(err)
	}
}

func appendOne(log *msg.Log) error {
	buf := make([]byte, log.Size())
	err := log.Unmarshal(buf)
	if err != nil {
		return err
	}
	//先占地,然后再写
	genKey := atomic.AddUint64(&generatedKey, 1)
	log.Index = genKey
	err = logDB.Set(util.UInt64ToBytes(genKey), buf)
	if err != nil {
		return err
	}
	atomic.StoreUint64(&lastLogIndex, util.Max(atomic.LoadUint64(&lastLogIndex), genKey))
	atomic.StoreUint32(&lastLogTerm, atomic.LoadUint32(&nodeTerm))
	return nil
}
func batchAppend(logs []*msg.Log) error {
	l := len(logs)
	ks := make([][]byte, l)
	vs := make([][]byte, l)
	//先占地,然后再写
	lastIndex := atomic.AddUint64(&generatedKey, uint64(l))
	curIndex := lastIndex - uint64(l)
	for i, log := range logs {
		curIndex++
		log.Index = curIndex
		ks[i] = util.UInt64ToBytes(curIndex)
		vs[i] = make([]byte, log.Size())
		err := log.Unmarshal(vs[i])
		if err != nil {
			return err
		}
	}
	err := logDB.BatchSet(ks, vs)
	if err != nil {
		return err
	}
	atomic.StoreUint64(&lastLogIndex, util.Max(atomic.LoadUint64(&lastLogIndex), lastIndex))
	atomic.StoreUint32(&lastLogTerm, atomic.LoadUint32(&nodeTerm))
	return nil
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
