package raft

import (
	"errors"
	"github.com/Mstch/naruto/helper/db"
	"github.com/Mstch/naruto/raft/msg"
	"github.com/cockroachdb/pebble"
)

var kvDB *db.DB

func init() {
	var err error
	kvDB, err = db.NewDB("raft-kv", &pebble.Options{})
	if err != nil {
		panic(err)
	}
}
func apply(cmd *msg.Cmd) (string, error) {
	switch cmd.Opt {
	case msg.Get:
		v, err := kvDB.Get([]byte(cmd.Key), func(k, v []byte) (interface{}, error) {
			return string(v), nil
		})
		if err != nil {
			return "", err
		}
		return v.(string), err
	case msg.Set:
		err := kvDB.Set([]byte(cmd.Key), []byte(cmd.Value))
		if err != nil {
			return "", err
		}
		return "", nil
	default:
		return "", errors.New("unknown cmd opt")
	}
}
