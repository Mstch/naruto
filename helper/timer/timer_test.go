package timer

import (
	"testing"
	"time"
)

func TestChangeTo(t *testing.T) {
	fuckTimer := &RTimerOption{
		name: "fuck",
		task: func() {
			println("fuck")
		},
		min: 1000,
		max: 2000,
	}
	fuckedTimer := &RTimerOption{
		name: "fucked",
		task: func() {
			println("fucked")
		},
		min: 100,
		max: 100,
	}
	Loop(fuckTimer)
	time.Sleep(5 * time.Second)
	ticker := time.NewTicker(100 * time.Millisecond)
	for {
		<-ticker.C
		Loop(fuckedTimer)
	}
}
