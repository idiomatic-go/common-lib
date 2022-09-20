package vhost

import (
	"github.com/idiomatic-go/common-lib/vhost/usr"
)

// Credentials function methds
func CreateCredentialsMessage(event, sender string, fn usr.Credentials) *usr.Message {
	return CreateMessage(event, sender, fn)
}

func AccessCredentials(msg *usr.Message) usr.Credentials {
	if msg == nil || msg.Content == nil || len(msg.Content) == 0 {
		return nil
	}
	for _, c := range msg.Content {
		fn, ok := c.(usr.Credentials)
		if ok {
			return fn
		}
	}
	return nil
}
