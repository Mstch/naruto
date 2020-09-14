package sbuf

import (
	"golang.org/x/sys/cpu"
	"sync"
	"sync/atomic"
)

type RingBuffer struct {
	qbufChangeLock sync.RWMutex
	qBuf           []*quoterBuf
	workbuf        *quoterBuf
	restbuf        *quoterBuf
}

type quoterBuf struct {
	_      cpu.CacheLinePad
	index  int64
	_      cpu.CacheLinePad
	waiter *sync.WaitGroup
	buf    []byte
}

func NewRingBuffer() *RingBuffer {
	return &RingBuffer{
		qbufChangeLock: sync.RWMutex{},
		workbuf: &quoterBuf{
			index:  0,
			waiter: &sync.WaitGroup{},
			buf:    make([]byte, 1024),
		},
		restbuf: &quoterBuf{
			index:  0,
			waiter: &sync.WaitGroup{},
			buf:    make([]byte, 1024),
		},
	}
}

func (r *RingBuffer) Take(takeSize int) (buf []byte, releaser func()) {
	r.qbufChangeLock.RLock()
	size := ceilingSize(takeSize)
	workbuf := r.workbuf
	c := cap(workbuf.buf)
	if i := int(atomic.AddInt64(&workbuf.index, int64(size))); i < c {
		workbuf.waiter.Add(1)
		r.qbufChangeLock.RUnlock()
		return workbuf.buf[i-size : i+1], workbuf.release
	} else if size <= c {
		r.restbuf.clean()
		r.restbuf, r.workbuf = r.workbuf, r.restbuf
		r.qbufChangeLock.RUnlock()
		return r.Take(size)
	} else {
		r.qbufChangeLock.RUnlock()
		r.qbufChangeLock.Lock()
		var newSize int
		if size < 1024 {
			newSize = size * 2
		} else {
			newSize = size + size/4
		}
		r.workbuf = &quoterBuf{
			index:  0,
			waiter: &sync.WaitGroup{},
			buf:    make([]byte, newSize),
		}
		r.restbuf = &quoterBuf{
			index:  0,
			waiter: &sync.WaitGroup{},
			buf:    make([]byte, newSize),
		}
		r.qbufChangeLock.Unlock()
		return r.Take(size)
	}
}

func (q *quoterBuf) release() {
	q.waiter.Done()
}

func (q *quoterBuf) clean() {
	q.waiter.Wait()
	atomic.StoreInt64(&q.index, 0)
}
