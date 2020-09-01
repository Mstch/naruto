package db

import (
	"encoding/binary"
	"github.com/cockroachdb/pebble"
	"testing"
)

func TestIter(t *testing.T) {
	db, err := NewDB("test-data", &pebble.Options{})
	if err != nil {
		panic(err)
	}
	for i := 0; i < 100; i++ {
		kbuf := make([]byte, 8)
		binary.BigEndian.PutUint64(kbuf, uint64(i))
		err = db.Set(kbuf, kbuf)
		if err != nil {
			panic(err)
		}
	}
	start := make([]byte, 8)
	end := make([]byte, 8)
	binary.BigEndian.PutUint64(start, 10)
	binary.BigEndian.PutUint64(end, 20)
}
func TestDB_Get(t *testing.T) {
	db, err := NewDB("test-data", &pebble.Options{})
	if err != nil {
		panic(err)
	}
	_, err = db.Get([]byte("fuck"), func(k, v []byte) (interface{}, error) {
		return nil, nil
	})
	if err != nil {
		panic(err)
	}
}
