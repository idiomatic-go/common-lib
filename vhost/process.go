package vhost

import (
	"errors"
	"fmt"
	"github.com/idiomatic-go/common-lib/util"
	"log"
	"time"
)

func startupProcess(ticks int, toSend messageMap) bool {
	log.Printf("Startup begin with iteration seconds: %v", ticks)
	sent := make(util.List)
	var count = 1
	current := make(messageMap)
	for {
		if count > MaxStartupIterations {
			log.Printf("Startup failure: max iterations excedded: %v", count)
			return false
		}
		log.Printf("Startup iteration: %v", count)
		if count == 1 || len(current) == 0 {
			err := processWork(sent, toSend, current, directory)
			if err != nil {
				log.Printf("%v", err)
				return false
			}
		}
		time.Sleep(time.Second * time.Duration(ticks))
		// Check the startup status of the directory, continue if a package is still in startup
		uri := directory.startupInProgress()
		if uri != "" {
			log.Printf("Startup still in progress continuing: %v", uri)
			count++
			continue
		}
		// All the current messages have been sent, so lets check for failure.
		fail := directory.startupFailure()
		if fail != "" {
			log.Printf("Startup failure status on: %v", fail)
			return false
		}
		// Success so empty current work map and check for completion
		empty(current)
		if len(toSend) == 0 {
			log.Printf("Startup successful: %v", count)
			return true
		}
		count++
	}
	return true
}

var processWork work = func(sent util.List, toSend messageMap, current messageMap, dir *syncMap) error {
	getCurrentWork(sent, toSend, current, dir)
	// Did not find any messages to send, but there are still messages waiting in the to send map
	if len(current) == 0 && len(toSend) > 0 {
		return errors.New(fmt.Sprintf("Startup failure: %v", "unable to find items to work, verify cyclic dependencies"))
	}
	// Process the current work map
	for k := range current {
		if !dir.setStatus(k, StatusInProgress) {
			return errors.New(fmt.Sprintf("Startup failure: unable to set package %v startup status", k))
		}
		sent.Add(k)
		SendMessageWithDirectory(current[k], dir)
	}
	return nil
}
