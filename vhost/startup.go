package vhost

import (
	"errors"
	"fmt"
	"github.com/idiomatic-go/common-lib/logxt"
	"time"
)

var maxStartupIterations = DefaultMaxStartupIterations

// Response methods
var resp chan Message
var q *Queue

func init() {
	resp = make(chan Message, 100)
	q = CreateQueue()
	go receive(q)
}

// IsPackageStartupSuccessful - determine if a package was successfully started
func IsPackageStartupSuccessful(uri string) bool {
	count := 1
	for {
		if count > 3 {
			return false
		}
		count++
		status := directory.getStatus(uri)
		switch status {
		case StatusSuccessful:
			return true
		case StatusFailure:
			return false
		default:
			time.Sleep(time.Second * time.Duration(1))
		}
	}
}

// RegisterPackage - function to register a package uri
func RegisterPackage(uri string, c chan Message, dependents []string) error {
	if uri == "" {
		return errors.New("Startup RegisterPackage() error : uri is empty")
	}
	if c == nil {
		return errors.New("Startup RegisterPackage() error : channel is nil")
	}
	registerPackageUnchecked(uri, c, dependents)
	return nil
}

func registerPackageUnchecked(uri string, c chan Message, dependents []string) error {
	directory.put(&entry{uri: uri, c: c, dependents: dependents})
	return nil
}

// UnregisterPackage - function to unregister a package
func UnregisterPackage(uri string) {
	if uri == "" {
		return
	}
	entry := directory.get(uri)
	if entry != nil {
		if entry.c != nil {
			close(entry.c)
		}
		delete(directory.m, uri)
	}
}
func unregisterPackages() {

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
	err = sendMessages(toSend)
	if err != nil {
		logxt.LogPrintf("%v", err)
		unregisterPackages()
		return false
	}
	var count = 1
	for {
		if count > maxStartupIterations {
			logxt.LogPrintf("Startup failure %v, max iterations exceeded: %v", directory.notSuccessfulStatus(), count)
			unregisterPackages()
			return false
		}
		time.Sleep(time.Second * time.Duration(ticks))
		// Check the startup status of the directory, continue if a package is still in startup
		uri := directory.inProgress()
		if uri != "" {
			logxt.LogPrintf("Startup in progress: continuing: %v", uri)
			count++
			continue
		}
		// All the current messages have been sent, so lets check for failure.
		fail := directory.failure()
		if fail != "" {
			logxt.LogPrintf("Startup failure: status on: %v", fail)
			unregisterPackages()
			return false
		}
		logxt.LogPrintf("Startup successful: %v", count)
		break
	}
	return true
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
			return errors.New(fmt.Sprintf("Startup failure: directory entry does not exist for package uri: %v", k))
		}
	}
	return nil
}

func sendMessages(msgs messageMap) error {
	for k := range msgs {
		if !directory.setStatus(k, StatusInProgress) {
			return errors.New(fmt.Sprintf("Startup failure: unable to set package %v startup status", k))
		}
		SendMessage(msgs[k])
	}
	return nil
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
					logxt.LogPrintf("Startup failure: unable to set startup status from package: %v", msg.From)
				}
			} else {
				//q.Enqueue(msg)
				// All messages that are received must have valid processing, otherwise log an error
				logxt.LogPrintf("vhost message received error : unable to process message, no mapping for event : %v", msg)
			}
		}
	}
}
