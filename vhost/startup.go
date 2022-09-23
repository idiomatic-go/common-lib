package vhost

import (
	"time"
)

// Response methods
var resp chan Message
var q *Queue

func init() {
	resp = make(chan Message, 100)
	q = CreateQueue()
	go receive(q)
}

// Virtual host startup
func Startup(timeout int, msgs map[string]Envelope) bool {
	packages := len(directory)
	if packages == 0 {
		return true
	}
	// Send startup messages
	for k := range directory {
		if msgs != nil {
			e, ok := msgs[k]
			if ok {
				e.Msg.Event = StartupEvent
				e.Msg.Sender = HostSender
				SendMessage(k, e.Msg)
				continue
			}
		}
		SendMessage(k, CreateMessage(StartupEvent, HostSender, nil))
	}
	time.Sleep(time.Second * time.Duration(timeout))
	if q.IsErrorEvent() {
		q.Empty()
		return false
	}
	valid := true
	if q.Count() != packages {
		valid = false
		for k := range directory {
			if !q.Exists(k) {
				LogPrintf("Missing startup message response from : %v", k)
			}
		}
	}
	q.Empty()
	return valid
}

func receive(q *Queue) {
	for {
		select {
		case msg, open := <-resp:
			// Exit on a closed channel
			if !open {
				return
			}
			q.Enqueue(msg)
		}
	}
}
