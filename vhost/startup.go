package vhost

import (
	"errors"
	"fmt"
	"github.com/idiomatic-go/common-lib/util"
	"log"
	"time"
)

type currentWork func(sent util.List, toSend messageMap, current messageMap, dir *syncMap) error
type toSend func(sent util.List, entry *entry, dir *syncMap) (bool, error)

// Response methods
var resp chan Message
var q *Queue

func init() {
	resp = make(chan Message, 100)
	q = CreateQueue()
	go receive(q)
}

// Startup - virtual host startup
func Startup(ticks int, msgs envelopeMap) bool {
	packages := directory.count()
	if packages == 0 {
		return true
	}
	// Create the toSend and sent maps
	log.Printf("Startup begin with iteration seconds: %v", ticks)
	toSend := createToSend(msgs)
	sent := make(util.List)
	var count = 1
	item := struct{}{}
	current := make(messageMap)
	for {
		if count > MaxStartupIterations {
			log.Printf("Startup failure: max iterations excedded: %v", count)
			return false
		}
		log.Printf("Startup iteration: %v", count)
		if count == 1 || len(current) == 0 {
			err := getCurrentWork(sent, toSend, current, directory)
			// This is either a startup message with a directory entry that does not exist, or a startup failure
			// on a dependent package
			if err != nil {
				log.Printf("Startup failure: getting current work: %v", err)
				return false
			}
			// Did not find any messages to send, but there are still messages waiting in the to send map
			if len(current) == 0 && len(toSend) > 0 {
				log.Printf("Startup failure: %v", "unable to find items to work, verify cyclic dependencies")
				return false
			}
			// Process the current work map
			for k := range current {
				if !directory.setStatus(k, StatusInProgress) {
					log.Printf("Startup failure: unable to set package %v startup status", k)
					return false
				}
				sent[k] = item
				SendMessage(k, toSend[k])
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

var getCurrentWork = func(sent util.List, toSend messageMap, current messageMap, dir *syncMap) error {
	for k := range toSend {
		e := dir.get(k)
		if e == nil {
			return errors.New(fmt.Sprintf("directory entry does not exist for package uri: %v", k))
		}
		ok, err := validToSend(sent, e, dir)
		if err != nil {
			return err
		}
		if ok {
			current[k] = toSend[k]
			delete(toSend, k)
		}
	}
	return nil
}

var validToSend = func(sent util.List, entry *entry, dir *syncMap) (bool, error) {
	if entry == nil || sent == nil || dir == nil {
		return false, errors.New("invalid argument for validToSend() : one of list, entry or directory is nil")
	}
	// No dependencies, so can be sent
	if len(entry.dependents) == 0 {
		return true, nil
	}
	// Need to determine if all dependencies have been sent and are successful
	for _, uri := range entry.dependents {
		if !sent.Contains(uri) {
			return false, nil
		}
		status, ok := dir.getStatus(uri)
		if !ok {
			return false, errors.New(fmt.Sprintf("dependency not fufilled, package entry not found: %v", uri))
		}
		if status != StatusSuccessful {
			return false, errors.New(fmt.Sprintf("dependency not fufilled, startup has failed for package: %v", uri))
		}
	}
	return true, nil
}

func createToSend(msgs envelopeMap) messageMap {
	m := make(messageMap)
	for k := range directory.data() {
		if msgs != nil {
			env, ok := msgs[k]
			if ok {
				env.Msg.Event = StartupEvent
				env.Msg.From = HostFrom
				m[k] = env.Msg
				continue
			}
		}
		m[k] = CreateMessage(StartupEvent, HostFrom, 0, nil)
	}
	return m
}

func receive(q *Queue) {
	for {
		select {
		case msg, open := <-resp:
			// Exit on a closed channel
			if !open {
				return
			}
			if msg.Event == StartupEvent {
				if !directory.setStatus(msg.From, msg.Status) {
					log.Printf("failure to set startup status from package: %v", msg.From)
				}
			} else {
				q.Enqueue(msg)
			}
		}
	}
}
