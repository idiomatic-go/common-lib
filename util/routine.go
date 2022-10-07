package util

import "time"

// InvokeRoutine - a routine that invokes a handler based on a timer.
func InvokeRoutine(repeat bool, interval time.Duration, handler Func, close *ClosableChannel) {
	if handler == nil {
		return
	}
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		if close != nil {
			select {
			case _, open := <-close.C:
				if !open {
				}
				LogDebug("%s\n", "invoke-routine : stopped")
				return
			default:
			}
		}
		select {
		case <-ticker.C:
			LogDebug("%s\n", "invoke-routine : invoke")
			if handler != nil {
				handler()
			}
			if !repeat {
				LogDebug("%s\n", "invoke-routine : finished")
				return
			}
		default:
		}
	}
}

// ExchangeRoutine - A routine that implements channeled request-response semantics.
func ExchangeRoutine(handler FuncResponse, resp *ResponseChannel, close *ClosableChannel) {
	if resp == nil || handler == nil {
		return
	}
	ticker := resp.NewTicker()
	defer ticker.Stop()

	for {
		if close != nil {
			select {
			case _, open := <-close.C:
				if !open {
				}
				LogDebug("%v", "exchange-routine : stopped")
				return
			default:
			}
		}
		select {
		case <-ticker.C:
			if handler != nil {
				LogDebug("%v", "exchange-routine : invoke")
				resp.C <- handler()
			}
		default:
		}
	}
}
