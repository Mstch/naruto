package sbuf

import "sync"

type CompactBuffer [64][64]*sync.Pool

func NewCompactBuffer() *CompactBuffer {
	c := &CompactBuffer{}
	for i := 0; i < 64; i++ {
		for j := 0; j < 64; j++ {
			c[i][j] = &sync.Pool{}
		}
	}
	return c
}

func (c *CompactBuffer) Take(takeSize int) []byte {
	size := ceilingSize(takeSize)
	segmentIndex := (size/64 - 1) / 64
	bufferIndex := (size/64 - 1) % 64
	pool := c[segmentIndex][bufferIndex]
	buf := pool.Get()
	if buf == nil {
		stealBufIndex := bufferIndex
		stealSegmentIndex := segmentIndex
		for i := 0; i < 16; i++ {
			if stealBufIndex+i >= 64 {
				stealBufIndex = 0
				stealSegmentIndex++
				if stealSegmentIndex > 64 {
					return make([]byte, size)[:takeSize]
				}
			}
			if stealBuf := c[stealSegmentIndex][stealBufIndex].Get(); stealBuf != nil {
				return stealBuf.([]byte)[:takeSize]
			}
		}
		return make([]byte, size)[:takeSize]
	} else {
		return buf.([]byte)[:takeSize]
	}
}

func (c *CompactBuffer) Release(buf []byte) {
	size := cap(buf)
	segmentIndex := (size/64 - 1) / 64
	bufferIndex := (size/64 - 1) % 64
	c[segmentIndex][bufferIndex].Put(buf)
}
