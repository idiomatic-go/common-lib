package vhost

import (
	"fmt"
)

func CreateMessage(to, event, from string, status int32, content any) Message {
	msg := Message{To: to, Event: event, From: from, Status: int32(status), Content: nil}
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

func SendMessage(msg Message) error {
	e := directory.get(msg.To)
	if e == nil {
		return fmt.Errorf("invalid argument : to uri invalid %v", msg.To)
	}
	if e.c == nil {
		return fmt.Errorf("invalid initialization : channel is nil %v", msg.To)
	}
	e.c <- msg
	return nil
}

func SendStartupMessage(to, from string) error {
	return SendMessage(Message{To: to, Event: StartupEvent, From: from})
}

func SendShutdownMessage(to, from string) error {
	return SendMessage(Message{To: to, Event: ShutdownEvent, From: from})
}

// SendResponse - response processing
func SendResponse(msg Message) {
	resp <- msg
}

// SendErrorResponse -
func SendErrorResponse(from string) {
	SendResponse(Message{Event: ErrorEvent, From: from, Content: nil})
}

func SendAckResponse(from string) {
	SendResponse(Message{Event: ACKEvent, From: from, Content: nil})
}

func SendStartupSuccessfulResponse(from string) {
	SendResponse(Message{Event: StartupEvent, From: from, Status: StatusSuccessful, Content: nil})
}

func SendStartupFailureResponse(from string) {
	SendResponse(Message{Event: StartupEvent, From: from, Status: StatusFailure, Content: nil})
}
