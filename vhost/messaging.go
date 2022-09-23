package vhost

import (
	"fmt"
)

type entry struct {
	uri string
	c   chan Message
}

var directory = make(map[string]*entry)

func RegisterPackage(uri string, c chan Message) error {
	if uri == "" {
		return fmt.Errorf("invalid argument : uri is empty")
	}
	if c == nil {
		return fmt.Errorf("invalid argument : channel is nil")
	}
	directory[uri] = &entry{uri, c}
	return nil
}

func UnregisterPackage(uri string) {
	if uri == "" {
		return
	}
	entry := directory[uri]
	if entry != nil {
		close(entry.c)
		directory[uri] = nil
	}
}

func CreateMessage(event, sender string, content any) Message {
	msg := Message{Event: event, Sender: sender, Content: nil}
	if content != nil {
		AddContent(&msg, content)
	}
	return msg
}

func AddContent(msg *Message, content any) {
	if msg == nil || content == nil {
		return
	}
	msg.Content = append(msg.Content, content)
}

func SendMessage(uri string, msg Message) error {
	e := directory[uri]
	if e == nil {
		return fmt.Errorf("invalid argument : %v", uri)
	}
	e.c <- msg
	return nil
}

// Response processing
func SendResponse(msg Message) {
	resp <- msg
}

func SendErrorResponse(sender string) {
	SendResponse(Message{Event: ErrorEvent, Sender: sender, Content: nil})
}

func SendAckResponse(sender string) {
	SendResponse(Message{Event: ACKEvent, Sender: sender, Content: nil})
}
