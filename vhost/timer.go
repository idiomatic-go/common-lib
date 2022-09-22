package vhost

import (
	"github.com/idiomatic-go/common-lib/vhost/usr"
	"time"
)

func NewStopChannel() chan struct{} {
	return make(chan struct{}, 1)
}

func StopTimer(c chan struct{}) {
	if c != nil {
		close(c)
	}
}

// Timer - a simple timer with notification.
// Note.: Create a stop channel with a minimum capacity of 1, otherwise, the Timer will block waiting on
//        the stop channel
func Timer(repeat bool, interval time.Duration, stop chan struct{}, fn usr.Notify) {
	ticker := time.NewTicker(interval)

	for {
		if stop != nil {
			select {
			case <-stop:
				LogDebug("%s\n", "timer : stopped")
				return
			default:
			}
		}
		select {
		case <-ticker.C:
			LogDebug("%s\n", "timer : notify")
			if fn != nil {
				fn()
			}
			if !repeat {
				LogDebug("%s\n", "timer : finished")
				return
			}
		default:
		}
	}
}