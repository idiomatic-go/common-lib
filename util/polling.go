package util

import (
	"time"
)

func Polling(resp chan any, respInterval time.Duration, stop chan struct{}, stopInterval time.Duration, handler NiladicResponse) {
	respTicker := time.NewTicker(respInterval)
	var stopTicker *time.Ticker

	defer respTicker.Stop()
	if stop != nil {
		stopTicker = time.NewTicker(stopInterval)
		defer stopTicker.Stop()
	}

	for {
		// Using separate select statements to avoid starvation of the close message.
		// Go documentation states that the selection of a case statement is indeterminate when more than one case
		// is asserted.
		if stop != nil {
			select {
			case <-stopTicker.C:
				select {
				case <-stop:
					LogDebug("%v", "polling : closed")
					return
				default:
				}
			default:
			}
		}
		select {
		case <-respTicker.C:
			if handler != nil {
				LogDebug("%v", "polling : invoke")
				resp <- handler()
			}
		default:
		}
	}
}
