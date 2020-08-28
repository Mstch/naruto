package raft

import (
	"encoding/binary"
	"github.com/Mstch/naruto/helper/db"
	"github.com/Mstch/naruto/helper/logger"
	"github.com/Mstch/naruto/raft/msg"
	"github.com/cockroachdb/pebble"
	"github.com/gogo/protobuf/proto"
	"sync/atomic"
)

var logDB *db.DB

func init() {
	var err error
	logDB, err = db.NewDB("raft-log", &pebble.Options{})
	if err != nil {
		panic(err)
	}
}
func applyTo(index uint64) {
	start := make([]byte, 8)
	end := make([]byte, 8)
	binary.BigEndian.PutUint64(start, atomic.LoadUint64(&lastApplyIndex))
	binary.BigEndian.PutUint64(start, index)
	_, err := logDB.Iter(start, end, false, func(k, v []byte) (interface{}, error) {
		log := &msg.Log{}
		err := proto.Unmarshal(v, log)
		if err != nil {
			return nil, err
		}
		_, err = apply(log.Cmd)
		return nil, err
	})
	logger.Error("apply to %d failed:%s", index, err)
}
func append(log msg.Log) {

}
func batchAppend(logs []msg.Log) {

}
func getLastlogIndex() uint64 {

}
