package vhost

import (
	"time"

	"github.com/idiomatic-go/common-lib/vhost/internal"
	"github.com/idiomatic-go/common-lib/vhost/usr"
)

// Response methods
var resp chan *usr.Message
var q *internal.Queue

func init() {
	resp = make(chan *usr.Message, 100)
	q = internal.CreateQueue()
	go receive(q)
}

// Virtual host startup
func Startup(timeout int, msgs map[string]usr.Envelope) bool {
	packages := len(directory)
	if packages == 0 {
		return true
	}
	// Send startup messages
	for k := range directory {
		if msgs != nil {
			e, ok := msgs[k]
			if ok {
				e.Msg.Event = usr.StartupEvent
				e.Msg.Sender = usr.HostSender
				SendMessage(k, e.Msg)
				continue
			}
		}
		SendMessage(k, CreateMessage(usr.StartupEvent, usr.HostSender, nil))
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

func receive(q *internal.Queue) {
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
