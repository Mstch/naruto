package timer

import (
	"testing"
	"time"
)

func TestChangeTo(t *testing.T) {
	fuckTimer := &RTimerOption{
		Name: "fuck",
		Task: func() {
			println("fuck")
		},
		Min: 1000,
		Max: 2000,
	}
	fuckedTimer := &RTimerOption{
		Name: "fucked",
		Task: func() {
			println("fucked")
		},
		Min: 100,
		Max: 100,
	}
	Loop(fuckTimer)
	time.Sleep(5 * time.Second)
	ticker := time.NewTicker(100 * time.Millisecond)
	for {
		<-ticker.C
		Loop(fuckedTimer)
	}
}
