package msg

import (
	"encoding/json"
	"testing"
)

func BenchmarkMarshalProto(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = (&AppendReq{
			Id:           uint32(i),
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
	}
}

func BenchmarkMarshalJson(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(&AppendReq{
			Id:           uint32(i),
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
	}
}
