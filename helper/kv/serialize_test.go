package kv

import (
	"bytes"
	"encoding/gob"
	"testing"
)

func BenchmarkGob(b *testing.B) {
	c := &Cmd{
		Opt:   "fuck",
		Key:   "fuck",
		Value: "fuck",
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.ToGobBytes()
	}
}
func BenchmarkJson(b *testing.B) {
	c := &Cmd{
		Opt:   "fuck",
		Key:   "fuck",
		Value: "fuck",
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.ToJsonBytes()
	}
}

func TestCmd_ToJsonBytes(t *testing.T) {
	c := &Cmd{
		Opt:   "fuck",
		Key:   "fuck",
		Value: "fuck",
	}
	println(string(c.ToJsonBytes()))
}
func TestCmd_ToGobBytes(t *testing.T) {
	c := &Cmd{
		Opt:   "fuck",
		Key:   "fuck",
		Value: "fuck",
	}
	nc := &Cmd{}
	buf := bytes.NewBuffer(c.ToGobBytes())
	_ = gob.NewDecoder(buf).Decode(nc)
	println(nc)
}
