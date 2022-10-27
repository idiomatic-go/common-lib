package startup

import (
	"errors"
	"fmt"
	"github.com/idiomatic-go/common-lib/logxt"
	"github.com/idiomatic-go/common-lib/util"
	"time"
)

func startupProcess(ticks int, toSend messageMap) bool {
	logxt.LogPrintf("Startup begin with iteration seconds: %v", ticks)
	sent := make(util.List)
	var count = 1
	current := make(messageMap)
	for {
		if count > maxStartupIterations {
			logxt.LogPrintf("Startup failure: max iterations excedded: %v", count)
			return false
		}
		logxt.LogPrintf("Startup iteration: %v", count)
		if count == 1 || len(current) == 0 {
			err := processWork(sent, toSend, current)
			if err != nil {
				logxt.LogPrintf("%v", err)
				return false
			}
		}
		time.Sleep(time.Second * time.Duration(ticks))
		// Check the startup status of the directory, continue if a package is still in startup
		uri := directory.startupInProgress()
		if uri != "" {
			logxt.LogPrintf("Startup still in progress continuing: %v", uri)
			count++
			continue
		}
		// All the current messages have been sent, so lets check for failure.
		fail := directory.startupFailure()
		if fail != "" {
			logxt.LogPrintf("Startup failure status on: %v", fail)
			return false
		}
		// Success so empty current work map and check for completion
		empty(current)
		if len(toSend) == 0 {
			logxt.LogPrintf("Startup successful: %v", count)
			return true
		}
		count++
	}
	return true
}

var processWork work = func(sent util.List, toSend messageMap, current messageMap) error {
	getCurrentWork(sent, toSend, current)
	// Did not find any messages to send, but there are still messages waiting in the to send map
	if len(current) == 0 && len(toSend) > 0 {
		return errors.New(fmt.Sprintf("Startup failure: %v", "unable to find items to work, verify cyclic dependencies"))
	}
	// Process the current work map
	for k := range current {
		if !directory.setStatus(k, StatusInProgress) {
			return errors.New(fmt.Sprintf("Startup failure: unable to set package %v startup status", k))
		}
		sent.Add(k)
		SendMessage(current[k])
	}
	return nil
}
