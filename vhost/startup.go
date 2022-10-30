package vhost

import (
	"errors"
	"fmt"
	"github.com/idiomatic-go/common-lib/logxt"
	"github.com/idiomatic-go/common-lib/util"
)

var maxStartupIterations = DefaultMaxStartupIterations

type work func(sent util.List, toSend messageMap, current messageMap) error
type currentWork func(sent util.List, toSend messageMap, current messageMap) bool
type toSend func(sent util.List, entry *entry) bool

// Response methods
var resp chan Message
var q *Queue

func init() {
	resp = make(chan Message, 100)
	q = CreateQueue()
	go receive(q)
}

// Startup - virtual host startup
func Startup(ticks int, override messageMap) bool {
	packages := directory.count()
	if packages == 0 {
		return true
	}
	toSend := createToSend(override)
	err := validateToSend(toSend)
	if err != nil {
		logxt.LogPrintf("%v", err)
		return false
	}
	return startupProcess(ticks, toSend)
}

func createToSend(msgs messageMap) messageMap {
	m := make(messageMap)
	for k := range directory.data() {
		if msgs != nil {
			message, ok := msgs[k]
			if ok {
				message.Event = StartupEvent
				message.From = HostFrom
				message.Status = StatusEmpty
				m[k] = message
				continue
			}
		}
		e := directory.get(k)
		if e != nil {
			m[k] = CreateMessage(e.uri, StartupEvent, HostFrom, StatusEmpty, nil)
		} else {
			m[k] = CreateMessage("invalid:uri", StartupEvent, HostFrom, StatusEmpty, nil)
		}
	}
	return m
}

func validateToSend(toSend messageMap) error {
	for k := range toSend {
		e := directory.get(k)
		if e == nil {
			return errors.New(fmt.Sprintf("directory entry does not exist for package uri: %v", k))
		}
		for _, k2 := range e.dependents {
			e := directory.get(k2)
			if e == nil {
				return errors.New(fmt.Sprintf("directory entry does not exist for dependent package uri: %v", k2))
			}
		}
	}
	return nil
}

var getCurrentWork currentWork = func(sent util.List, toSend messageMap, current messageMap) bool {
	valid := false
	for k := range toSend {
		e := directory.get(k)
		if e == nil {
			continue
		}
		ok := validToSend(sent, e)
		if ok {
			valid = true
			current[k] = toSend[k]
			delete(toSend, k)
		}
	}
	return valid
}

var validToSend toSend = func(sent util.List, entry *entry) bool {
	// No dependencies, so can be sent
	if len(entry.dependents) == 0 {
		return true
	}
	// Need to determine if all dependencies have been sent and are successful
	for _, uri := range entry.dependents {
		if !sent.Contains(uri) {
			return false
		}
	}
	return true
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
					logxt.LogPrintf("failure to set startup status from package: %v", msg.From)
				}
			} else {
				q.Enqueue(msg)
			}
		}
	}
}
