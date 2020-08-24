package timer

import (
	"github.com/Mstch/naruto/helper/logger"
	"math/rand"
	"time"
)

var (
	timer        = time.NewTimer(0)
	currentTimer *RTimerOption
)

type RTimerOption struct {
	Name string
	Task func()
	Min  int
	Max  int
}



func Loop(t *RTimerOption) {
	if currentTimer == t {
		timer.Reset(calTimeout(t.Min, t.Max))
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
			timer.Reset(calTimeout(t.Min, t.Max))
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
