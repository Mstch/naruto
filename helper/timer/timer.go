package timer

import (
	"github.com/Mstch/raft/helper/logger"
	"math/rand"
	"time"
)

var (
	timer        = time.NewTimer(0)
	currentTimer *RTimerOption
)

type RTimerOption struct {
	name string
	task func()
	min  int
	max  int
}

func NewRTimerOption(name string, task func(), min int, max int) *RTimerOption {
	return &RTimerOption{name: name, task: task, min: min, max: max}
}

func Loop(t *RTimerOption) {
	if currentTimer == t {
		timer.Reset(calTimeout(t.min, t.max))
		return
	}
	currentTimer = t
	logger.Info("定时器切换为%s", t.name)
	if !timer.Stop() {
		<-timer.C
	}
	timer = time.NewTimer(calTimeout(t.min, t.max))
	go func() {
		for {
			<-timer.C
			go t.task()
			timer.Reset(calTimeout(t.min, t.max))
		}
	}()
}

func calTimeout(min, max int) time.Duration {
	timeout := min
	if max > min {
		timeout += rand.Intn(max - min)
	}
	return time.Duration(timeout) * time.Millisecond
}
