package util

import (
	"github.com/idiomatic-go/common-lib/vhost"
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
func Timer(repeat bool, interval time.Duration, stop chan struct{}, fn Dispatch) {
	ticker := time.NewTicker(interval)

	for {
		if stop != nil {
			select {
			case <-stop:
				vhost.LogDebug("%s\n", "timer : stopped")
				return
			default:
			}
		}
		select {
		case <-ticker.C:
			vhost.LogDebug("%s\n", "timer : notify")
			if fn != nil {
				fn()
			}
			if !repeat {
				vhost.LogDebug("%s\n", "timer : finished")
				return
			}
		default:
		}
	}
}
