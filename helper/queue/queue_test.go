package queue

import (
	"runtime"
	"sync"
	"testing"
	"time"
)

func BenchmarkSPSC(b *testing.B) {
	b.ResetTimer()
	c := make(chan int, 1024)
	go func() {
		for i := 0; i < b.N; i++ {
			c <- i
		}
	}()

	for i := 0; i < b.N; i++ {
		<-c
		time.Sleep(30 * time.Millisecond)
	}
}

func BenchmarkSPMC(b *testing.B) {
	b.ResetTimer()
	c := make(chan int, 1024)
	wait := sync.WaitGroup{}
	wait.Add(2 * runtime.NumCPU())
	for i := 0; i < 2*runtime.NumCPU(); i++ {
		go func() {
			for j := 0; j < b.N/(2*runtime.NumCPU()); j++ {
				<-c
				time.Sleep(30 * time.Millisecond)
			}
			wait.Done()
		}()
	}
	for i := 0; i < (b.N/(2*runtime.NumCPU()))*2*runtime.NumCPU(); i++ {
		c <- i
	}
	wait.Wait()
}

func BenchmarkMPSC(b *testing.B) {
	b.ResetTimer()
	c := make(chan int, 1024)
	for i := 0; i < 2*runtime.NumCPU(); i++ {
		go func() {
			for j := 0; j < b.N/(2*runtime.NumCPU()); j++ {
				c <- j
			}
		}()
	}

	for j := 0; j < (b.N/(2*runtime.NumCPU()))*2*runtime.NumCPU(); j++ {
		<-c
		time.Sleep(30 * time.Millisecond)
	}
}

func BenchmarkMPMC(b *testing.B) {
	b.ResetTimer()
	c := make(chan int, 1024)
	wait := sync.WaitGroup{}
	wait.Add(4 * runtime.NumCPU())
	for i := 0; i < 2*runtime.NumCPU(); i++ {
		go func() {
			for j := 0; j < b.N/(2*runtime.NumCPU()); j++ {
				<-c
				time.Sleep(30 * time.Millisecond)
			}
			wait.Done()
		}()
	}
	for i := 0; i < 2*runtime.NumCPU(); i++ {
		go func() {
			for j := 0; j < b.N/(2*runtime.NumCPU()); j++ {
				c <- j
			}
			wait.Done()
		}()
	}
	wait.Wait()
}

func BenchmarkSPSCMutex(b *testing.B) {
	l := sync.Mutex{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Lock()
		l.Unlock()
	}
}
func BenchmarkSPSCCAS(b *testing.B) {

}
