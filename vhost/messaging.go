package vhost

import (
	"fmt"

	"github.com/idiomatic-go/common-lib/vhost/usr"
)

type entry struct {
	uri string
	c   chan *usr.Message
}

var directory = make(map[string]*entry)

func RegisterPackage(uri string, c chan *usr.Message) error {
	if uri == "" {
		return fmt.Errorf("invalid argument : uri is empty")
	}
	if c == nil {
		return fmt.Errorf("invalid argument : channel is nil")
	}
	directory[uri] = &entry{uri, c}
	return nil
}

func CreateMessage(event, sender string, content any) *usr.Message {
	msg := &usr.Message{Event: event, Sender: sender, Content: nil}
	if content != nil {
		AddContent(msg, content)
	}
	return msg
}

func AddContent(msg *usr.Message, content any) {
	if msg == nil || content == nil {
		return
	}
	msg.Content = append(msg.Content, content)
}

func SendMessage(uri string, msg *usr.Message) error {
	e := directory[uri]
	if e == nil {
		return fmt.Errorf("invalid argument : %v", uri)
	}
	e.c <- msg
	return nil
}

// Response processing
func SendResponse(msg *usr.Message) {
	resp <- msg
}

func SendErrorResponse(sender string) {
	SendResponse(&usr.Message{Event: usr.ErrorEvent, Sender: sender, Content: nil})
}

func SendAckResponse(sender string) {
	SendResponse(&usr.Message{Event: usr.ACKEvent, Sender: sender, Content: nil})
}
