package timer

import (
	"github.com/Mstch/naruto/helper/logger"
	"math/rand"
	"sync"
	"time"
)

var (
	timer        = time.NewTimer(0)
	currentTimer *RTimerOption
	timerLock    = sync.Mutex{}
)

type RTimerOption struct {
	Name string
	Task func()
	Min  int
	Max  int
}

func Loop(t *RTimerOption) {
	if currentTimer == t {
		reset(timer, calTimeout(t.Min, t.Max))
		return
	}
	currentTimer = t
	logger.Info("定时器切换为%s", t.Name)
	if !timer.Stop() {
		<-timer.C
	}
	timer = time.NewTimer(calTimeout(t.Min, t.Max))
	go func() {
		for {
			<-timer.C
			go t.Task()
			reset(timer, calTimeout(t.Min, t.Max))
		}
	}()
}
func reset(timer *time.Timer, timeout time.Duration) {
	timerLock.Lock()
	timer.Reset(timeout)
	timerLock.Unlock()
}

func calTimeout(min, max int) time.Duration {
	timeout := min
	if max > min {
		timeout += rand.Intn(max - min)
	}
	return time.Duration(timeout) * time.Millisecond
}
