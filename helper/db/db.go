package db

import (
	"github.com/Mstch/naruto/helper/logger"
	"github.com/cockroachdb/pebble"
)

var (
	db *pebble.DB
	wo = pebble.Sync
)

func init() {
	var err error
	db, err = pebble.Open("raft-data",&pebble.Options{})
	if err != nil {
		logger.Fatal("db初始化失败", err)
	}
}

func Get(k []byte, serializer func([]byte) (interface{}, error)) (interface{}, error) {
	bytes, closer, err := db.Get(k)
	if err != nil {
		return nil, err
	}
	err = closer.Close()
	if err != nil {
		return nil, err
	}
	result, err := serializer(bytes)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func Set(k, v []byte) error {
	return db.Set(k, v, wo)
}
