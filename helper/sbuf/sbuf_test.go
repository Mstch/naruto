package sbuf

import (
	"math/rand"
	"sync"
	"testing"
)

var c = NewCompactBuffer()

func testCompactBuffer(c *CompactBuffer, n int, waiter *sync.WaitGroup, b *testing.B) {
	for i := 0; i < n; i++ {
		size := int(rand.NormFloat64() + 1024)
		if size < 64 {
			size = 64
		} else if size > 4096 {
			size = 4096
		}
		buf := c.Take(size)
		b.SetBytes(int64(cap(buf)))
		c.Release(buf)
	}
	waiter.Done()
}
func BenchmarkCompactBuffer_Take(b *testing.B) {
	waiter := &sync.WaitGroup{}
	waiter.Add(100)
	for i := 0; i < 100; i++ {
		go testCompactBuffer(c, b.N, waiter, b)
	}
	waiter.Wait()
}

func BenchmarkTreeBuffer_Take(b *testing.B) {
	waiter := &sync.WaitGroup{}
	waiter.Add(100)
	for i := 0; i < 100; i++ {
		go testTreeBuffer(b.N, waiter, b)
	}
	waiter.Wait()
}

func BenchmarkRingBuffer_Take(b *testing.B) {
	rb := NewRingBuffer()
	waiter := &sync.WaitGroup{}
	waiter.Add(100)
	for i := 0; i < 100; i++ {
		go func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				size := int(rand.NormFloat64() + 1024)
				if size < 64 {
					size = 64
				}
				_, releaser := rb.Take(size)
				releaser()
			}
			waiter.Done()
		}(b)
	}
	waiter.Wait()
}

func testTreeBuffer(n int, waiter *sync.WaitGroup, b *testing.B) {
	tree := NewTree(false)
	for i := 0; i < n; i++ {
		size := int(rand.NormFloat64() + 1024)
		if size < 64 {
			size = 64
		}
		tree.Take(size)
		b.SetBytes(int64(size))
	}
	waiter.Done()
}
