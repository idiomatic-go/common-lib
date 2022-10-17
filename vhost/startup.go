package vhost

import (
	"errors"
	"fmt"
	"github.com/idiomatic-go/common-lib/util"
	"log"
)

type work func(sent util.List, toSend messageMap, current messageMap, dir *syncMap) error
type currentWork func(sent util.List, toSend messageMap, current messageMap, dir *syncMap) bool
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
	toSend := createToSend(override, directory)
	err := validateToSend(toSend, directory)
	if err != nil {
		log.Printf("%v", err)
		return false
	}
	return startupProcess(ticks, toSend)
}

func createToSend(msgs messageMap, dir *syncMap) messageMap {
	m := make(messageMap)
	for k := range dir.data() {
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
		e := dir.get(k)
		if e != nil {
			m[k] = CreateMessage(e.uri, StartupEvent, HostFrom, StatusEmpty, nil)
		} else {
			m[k] = CreateMessage("invalid:uri", StartupEvent, HostFrom, StatusEmpty, nil)
		}
	}
	return m
}

func validateToSend(toSend messageMap, dir *syncMap) error {
	for k := range toSend {
		e := dir.get(k)
		if e == nil {
			return errors.New(fmt.Sprintf("directory entry does not exist for package uri: %v", k))
		}
		for _, k2 := range e.dependents {
			e := dir.get(k2)
			if e == nil {
				return errors.New(fmt.Sprintf("directory entry does not exist for dependent package uri: %v", k2))
			}
		}
	}
	return nil
}

var getCurrentWork currentWork = func(sent util.List, toSend messageMap, current messageMap, dir *syncMap) bool {
	valid := false
	for k := range toSend {
		e := dir.get(k)
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
					log.Printf("failure to set startup status from package: %v", msg.From)
				}
			} else {
				q.Enqueue(msg)
			}
		}
	}
}
