package vhost

import (
	"errors"
	"github.com/idiomatic-go/common-lib/eventing"
	"github.com/idiomatic-go/common-lib/logxt"
	"time"
)

type MessageMap map[string]eventing.Message

var maxStartupIterations = 4 //DefaultMaxStartupIterations

// IsPackageStartupSuccessful - determine if a package was successfully started
func IsPackageStartupSuccessful(uri string) bool {
	count := 1
	for {
		if count > 3 {
			return false
		}
		count++
		status := eventing.Directory.GetStatus(uri, eventing.StartupEvent)
		switch status {
		case StatusOk:
			return true
		case StatusInternal:
			return false
		default:
			time.Sleep(time.Second * time.Duration(1))
		}
	}
}

// RegisterPackage - function to register a package uri
func RegisterPackage(uri string, c chan eventing.Message) error {
	if uri == "" {
		return errors.New("startup RegisterPackage() error : uri is empty")
	}
	if c == nil {
		return errors.New("startup RegisterPackage() error : channel is nil")
	}
	registerPackageUnchecked(uri, c)
	return nil
}

func registerPackageUnchecked(uri string, c chan eventing.Message) error {
	eventing.Directory.Put(uri, c)
	return nil
}

// UnregisterPackage - function to unregister a package
/*
func unregisterPackage(uri string) {
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

*/

// Shutdown - virtual host shutdown
func Shutdown() {
	eventing.Directory.Broadcast(eventing.ShutdownEvent, eventing.StatusNotProvided)
	eventing.Directory.Shutdown()
}

// Startup - virtual host startup
func Startup(ticks int, override MessageMap) bool {
	packages := eventing.Directory.Count()
	if packages == 0 {
		return true
	}
	toSend := createToSend(override)
	sendMessages(toSend)
	var count = 1
	for {
		if count > maxStartupIterations {
			//LogPrintf("startup failure %v, max iterations exceeded: %v", directory.notSuccessfulStatus(StartupEvent), count)
			Shutdown()
			return false
		}
		time.Sleep(time.Second * time.Duration(ticks))
		// Check the startup status of the directory, continue if a package is still in startup
		uri := eventing.Directory.FindStatus(eventing.StartupEvent, StatusInProgress)
		if uri != "" {
			logxt.LogPrintf("startup in progress: continuing: %v\n", uri)
			count++
			continue
		}
		// All the current messages have been sent, so lets check for failure.
		fail := eventing.Directory.FindStatus(eventing.StartupEvent, StatusInternal)
		if fail != "" {
			logxt.LogPrintf("startup failure: status on: %v\n", fail)
			Shutdown()
			return false
		}
		logxt.LogPrintf("startup successful: %v\n", count)
		break
	}
	return true
}

func createToSend(msgs MessageMap) MessageMap {
	m := make(MessageMap)
	for _, k := range eventing.Directory.Uri() {
		if msgs != nil {
			message, ok := msgs[k]
			if ok {
				message.Event = eventing.StartupEvent
				message.From = eventing.VirtualHost
				message.Status = eventing.StatusNotProvided
				m[k] = message
				continue
			}
		}
		m[k] = eventing.CreateMessage(k, eventing.VirtualHost, eventing.StartupEvent, eventing.StatusNotProvided, nil)
	}
	return m
}

func sendMessages(msgs MessageMap) {
	for k := range msgs {
		eventing.Directory.SendMessage(msgs[k])
		eventing.Directory.Add(k, eventing.CreateMessage(eventing.VirtualHost, eventing.VirtualHost, eventing.StartupEvent, StatusInProgress, nil))
	}
}

func SendStartupSuccessfulResponse(from string) {
	eventing.SendResponse(eventing.CreateMessage(eventing.VirtualHost, from, eventing.StartupEvent, StatusOk, nil))
}

func SendStartupFailureResponse(from string) {
	eventing.SendResponse(eventing.CreateMessage(eventing.VirtualHost, from, eventing.StartupEvent, StatusInternal, nil))
}
