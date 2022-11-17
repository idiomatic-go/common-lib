package vhost

import "github.com/idiomatic-go/common-lib/eventing"

// CreateCredentialsMessage - functions
func CreateCredentialsMessage(to, from, event string, fn Credentials) eventing.Message {
	return eventing.CreateMessage(to, from, event, eventing.StatusNotProvided, fn)
}

func AccessCredentials(msg *eventing.Message) Credentials {
	if msg == nil || msg.Content == nil {
		return nil
	}
	fn, ok := msg.Content.(Credentials)
	if ok {
		return fn
	}
	return nil
}
