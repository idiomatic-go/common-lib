package vhost

import (
	"github.com/idiomatic-go/common-lib/vhost/usr"
	"time"
)

// Timer - a simple timer with a notification.
// Note.: Create a stop channel with a minimum capacity of 1, otherwise, the Timer will block waiting on
//        the stop channel
func Timer(repeat bool, interval time.Duration, stop chan struct{}, fn usr.Notify) {
	ticker := time.NewTicker(interval)

	for {
		if stop != nil {
			select {
			case <-stop:
				LogDebug("timer : stopped")
				return
			default:
			}
		}
		select {
		case <-ticker.C:
			LogDebug("timer : notify")
			if fn != nil {
				fn()
			}
			if !repeat {
				LogDebug("timer : finished")
				return
			}
		default:
		}
	}
}
