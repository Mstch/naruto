package raft

import (
	"errors"
	"github.com/Mstch/naruto/helper/db"
	"github.com/Mstch/naruto/helper/logger"
	"github.com/Mstch/naruto/helper/util"
	"github.com/Mstch/naruto/raft/msg"
	"github.com/cockroachdb/pebble"
	"github.com/gogo/protobuf/proto"
	"sync/atomic"
)

var (
	logDB          *db.DB
	generatedIndex uint64
	lastLogIndex   uint64
	lastLogTerm    uint32
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

func appendOne(log *msg.Log, fromReplicate bool) (uint64, uint32, error) {
	buf := make([]byte, log.Size())
	err := log.Unmarshal(buf)
	if err != nil {
		return 0, 0, err
	}
	if fromReplicate {
		if log.Index == 0 || log.Term == 0 {
			return 0, 0, errors.New("log marked as replicate,but not set index")
		}
		err = logDB.Set(util.UInt64ToBytes(log.Index), buf)
		if err != nil {
			return 0, 0, err
		}
	} else {
		//先占地,然后再写
		genIndex := atomic.AddUint64(&generatedIndex, 1)
		log.Index = genIndex
		log.Term = atomic.LoadUint32(&nodeTerm)
	}
	err = logDB.Set(util.UInt64ToBytes(log.Index), buf)
	if err != nil {
		return 0, 0, err
	}
	prevLogIndex := util.SwapToMaxUint64(&lastLogIndex, log.Index)
	prevLogTerm := util.SwapToMaxUint32(&lastLogTerm, log.Term)
	return prevLogIndex, prevLogTerm, nil
}
func batchAppend(logs []*msg.Log, fromReplicate bool) error {
	l := len(logs)
	if l == 0 {
		return nil
	}
	ks := make([][]byte, l)
	vs := make([][]byte, l)
	if fromReplicate {
		lastIndex := logs[l-1].Index
		lastTerm := logs[l-1].Term
		for i, log := range logs {
			ks[i] = util.UInt64ToBytes(log.Index)
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
		util.SwapToMaxUint64(&lastLogIndex, lastIndex)
		util.SwapToMaxUint32(&lastLogTerm, lastTerm)
		return nil
	}
	//先占地,然后再写
	lastIndex := atomic.AddUint64(&generatedIndex, uint64(l))
	curIndex := lastIndex - uint64(l)
	term := atomic.LoadUint32(&nodeTerm)
	for i, log := range logs {
		curIndex++
		log.Index = curIndex
		log.Term = term
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
	util.SwapToMaxUint64(&lastLogIndex, lastIndex)
	util.SwapToMaxUint32(&lastLogTerm, term)
	return nil
}

func applyTo(index uint64) error {
	atomic.SwapUint64(&lastCommitIndex, util.MaxUint64(atomic.LoadUint64(&lastCommitIndex), index))
	start := util.UInt64ToBytes(atomic.LoadUint64(&lastApplyIndex))
	end := make([]byte, index)
	_, err := logDB.Iter(start, end, false, func(k, v []byte) (interface{}, error) {
		log := &msg.Log{}
		err := proto.Unmarshal(v, log)
		if err != nil {
			return nil, err
		}
		_, err = apply(log.Cmd)
		if err == nil {
			util.SwapToMaxUint64(&lastApplyIndex, log.Index)
		}
		return nil, err
	})
	if err != nil {
		return err
	}
	return nil
}
func getLog(index uint64) (*msg.Log, error) {
	log, err := logDB.Get(util.UInt64ToBytes(index), func(_, v []byte) (interface{}, error) {
		l := &msg.Log{}
		err := proto.Unmarshal(v, l)
		if err != nil {
			logger.Error("unmarshal log %d failed when 	get log:%s", index, err)
			return nil, err
		}
		return l, nil
	})
	if err != nil {
		logger.Error("get log %d failed :%s", index, err)
		return nil, err
	}
	return log.(*msg.Log), nil
}
