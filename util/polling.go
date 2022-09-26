package util

import (
	"time"
)

func Polling(resp chan any, respPolling time.Duration, stop chan struct{}, stopPolling time.Duration, handler NiladicResponse) {
	stopTick := time.NewTicker(stopPolling)
	respTick := time.NewTicker(respPolling)

	// Warmup
	resp <- handler()
	for {
		// Using separate select statements to avoid starvation of the close message.
		// Go documentation states that the selection of a case statement is indeterminate when more than one case
		// is asserted.
		if stop != nil {
			select {
			case <-stopTick.C:
				select {
				case <-stop:
					LogDebug("polling : closed")
					return
				default:
				}
			default:
			}
		}
		select {
		case <-respTick.C:
			if handler != nil {
				resp <- handler()
			}
		default:
		}
	}
}
