package logger

import (
	"github.com/cockroachdb/pebble"
	"testing"
)

func BenchmarkSelfLogger(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Info("hello %d", i)
	}
}

func BenchmarkPebble(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pebble.DefaultLogger.Infof("hello %d", i)
	}
}
