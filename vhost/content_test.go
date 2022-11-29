package vhost

import (
	"context"
	"errors"
	"fmt"
	"github.com/idiomatic-go/common-lib/eventing"
)

type address struct {
	Name  string
	Email string
	Cell  string
}

func ExampleAccessCredentialsSuccess() {
	msg := CreateCredentialsMessage("to:uri", "event", "sender", Credentials(func() (username string, password string, err error) { return "", "", nil }))
	fmt.Printf("Credentials Fn : %v\n", AccessCredentials(&msg) != nil)

	//Output:
	// Credentials Fn : true
}

// Need to cast as adding content via any
func ExampleAccessCredentialsSlice() {
	content := Credentials(func() (username string, password string, err error) { return "", "", nil })

	msg := eventing.CreateMessage("to:uri", eventing.VirtualHost, "event", 0, content)
	//AddContent(&msg, Credentials(func() (username string, password string, err error) { return "", "", nil }))
	fmt.Printf("Credentials Fn : %v\n", AccessCredentials(&msg) != nil)

	//Output:
	// Credentials Fn : true
}

func ExampleProcessContentStatus() {
	status := NewStatusOk()
	ctx := ContextWithAnyContent(context.Background(), status)
	t, s := ProcessContent[address](ctx)
	if t.Cell != "" {
	}
	fmt.Printf("Status : %v\n", s.Ok())

	//Output:
	//Status : true
}

func ExampleProcessContentError() {
	err := errors.New("this is a test error")
	ctx := ContextWithAnyContent(context.Background(), err)
	_, s := ProcessContent[address](ctx)
	fmt.Printf("Error : %v\n", s.Error())

	//Output:
	//Error : this is a test error
}

func ExampleProcessContentType() {
	addr := address{Name: "Mark", Email: "mark@gmail.com", Cell: "123-456-7891"}
	ctx := ContextWithAnyContent(context.Background(), addr)
	t, s := ProcessContent[address](ctx)
	if t.Cell != "" {
	}
	fmt.Printf("Address : %v\n", t)
	fmt.Printf("Status  : %v\n", s.Ok())

	//Output:
	//Address : {Mark mark@gmail.com 123-456-7891}
	//Status  : true
}
