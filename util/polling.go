package util

import (
	"github.com/idiomatic-go/common-lib/vhost"
	"time"
)

func PollingDo(resp chan Response, respPolling time.Duration, stop chan struct{}, stopPolling time.Duration, fn Do) {
	stopTick := time.NewTicker(stopPolling)
	respTick := time.NewTicker(respPolling)

	// One call before waiting
	resp <- fn(nil)
	for {
		// Using seperate select statements because if both of the ticks ocurred at the same time, one tick would be
		// ignored possibly leading to starvation of the close message.
		// Go documentation states that the selection of a case statement is indeterminate when both cases are asserted at
		// the same time.
		if stop != nil {
			select {
			case <-stopTick.C:
				select {
				case <-stop:
					vhost.LogDebug("polling : closed")
					return
				default:
				}
			default:
			}
		}
		select {
		case <-respTick.C:
			if fn != nil {
				resp <- fn(nil)
			}
		default:
		}
	}
}
