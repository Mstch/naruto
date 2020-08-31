package quorum

import "sync/atomic"

type q struct {
	cur            uint32
	quorum         uint32
	quorumCallback func()
}

var (
	quorumMap        = make(map[uint64]*q)
	lastId    uint64 = 0
)

func RegQuorum(quorum uint32, quorumCallback func()) uint64 {
	id := atomic.AddUint64(&lastId, 1)
	quorumMap[id] = &q{
		cur:            0,
		quorum:         quorum,
		quorumCallback: quorumCallback,
	}
	return id
}
func Approve(id uint64) {
	q := quorumMap[id]
	if atomic.AddUint32(&q.cur, 1) > atomic.LoadUint32(&q.quorum) {
		delete(quorumMap,id)
	}
}
