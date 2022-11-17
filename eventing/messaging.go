package eventing

import (
	"time"
)

func CreateMessage(to, from, event string, status int32, content any) Message {
	return Message{To: to, From: from, Event: event, Status: status, Content: content, CreateTS: time.Now()}
}

func SendMessage(msg Message) error {
	return Directory.SendMessage(msg)
}

func SendResponse(msg Message) {
	resp <- msg
}

func SendErrorResponse(from string, status int32) {
	SendResponse(CreateMessage(VirtualHost, from, ErrorEvent, status, nil))
}
