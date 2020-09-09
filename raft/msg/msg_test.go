package msg

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"testing"
)

func BenchmarkMarshalProto(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf, _ := (&AppendReq{
			Id:           uint64(i),
			Term:         uint32(i),
			PrevLogIndex: uint64(i),
			PrevLogTerm:  uint32(i),
			LeaderCommit: uint64(i),
			Logs: []*Log{{
				Term:  uint32(i),
				Index: uint64(i),
				Cmd: &Cmd{
					Opt:   Get,
					Key:   "get",
					Value: "0",
				},
			}},
		}).Marshal()
		_ = (&AppendReq{}).Unmarshal(buf)
	}
}

func BenchmarkMarshalJson(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf, _ := json.Marshal(&AppendReq{
			Id:           uint64(i),
			Term:         uint32(i),
			PrevLogIndex: uint64(i),
			PrevLogTerm:  uint32(i),
			LeaderCommit: uint64(i),
			Logs: []*Log{{
				Term:  uint32(i),
				Index: uint64(i),
				Cmd: &Cmd{
					Opt:   Set,
					Key:   "get",
					Value: "0",
				},
			}},
		})
		_ = (&AppendReq{}).Unmarshal(buf)

	}
}

func BenchmarkMarshalGob(b *testing.B) {
	buf := &bytes.Buffer{}
	e := gob.NewEncoder(buf)
	d := gob.NewDecoder(buf)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := e.Encode(&AppendReq{
			Id:           uint64(i),
			Term:         uint32(i),
			PrevLogIndex: uint64(i),
			PrevLogTerm:  uint32(i),
			LeaderCommit: uint64(i),
			Logs: []*Log{{
				Term:  uint32(i),
				Index: uint64(i),
				Cmd: &Cmd{
					Opt:   Set,
					Key:   "get",
					Value: "0",
				},
			}},
		})
		d.Decode(&AppendReq{})
		if err != nil {
			panic(err)
		}
	}
}
