package vhost

import (
	"context"
	"errors"
	"fmt"
	"github.com/idiomatic-go/common-lib/eventing"
)

type location interface {
	Location() string
}

type address struct {
	Name  string
	Email string
	Cell  string
}

func (a *address) Location() string {
	return a.Name
}

type address2 struct {
	Name  string
	Email string
	Cell  string
	Zip   string
}

func init() {
	//logxt.ToggleDebug(true)
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

func ExampleProcessContextContentStatus() {
	status := NewStatusOk()
	ctx := ContextWithAnyContent(context.Background(), status)
	t, s := ProcessContextContent[address](ctx)
	if t.Cell != "" {
	}
	fmt.Printf("Status : %v\n", s.Ok())

	//Output:
	//Status : true
}

func ExampleProcessContextContentError() {
	err := errors.New("this is a test error")
	ctx := ContextWithAnyContent(context.Background(), err)
	_, s := ProcessContextContent[address](ctx)
	fmt.Printf("Error : %v\n", s.Error())

	//Output:
	//Error : this is a test error
}

func ExampleProcessContextContentType() {
	addr := address{Name: "Mark", Email: "mark@gmail.com", Cell: "123-456-7891"}
	ctx := ContextWithAnyContent(context.Background(), addr)
	t, s := ProcessContextContent[address](ctx)
	if t.Cell != "" {
	}
	fmt.Printf("Address : %v\n", t)
	fmt.Printf("Status  : %v\n", s.Ok())

	//Output:
	//Address : {Mark mark@gmail.com 123-456-7891}
	//Status  : true
}

func ExampleProcessContentInterface() {
	addr := address{Name: "Mark", Email: "mark@gmail.com", Cell: "123-456-7891"}
	var loc location = &addr

	ctx := ContextWithAnyContent(context.Background(), loc)
	l, s := ProcessContent[location](ContextAnyContent(ctx))
	fmt.Printf("Address : %v\n", l.Location())
	fmt.Printf("Status  : %v\n", s.Ok())

	//Output:
	//Address : Mark
	//Status  : true
}

func ExampleProcessContentErrors() {
	addr := address2{Name: "Mark", Email: "mark@gmail.com", Cell: "123-456-7891", Zip: "50436"}

	ContextWithAnyContent(context.Background(), addr)
	l, s := ProcessContent[address](nil)
	fmt.Printf("Address : %v\n", l)
	fmt.Printf("Ok      : %v\n", s.Ok())

	ctx := ContextWithAnyContent(context.Background(), addr)
	l, s = ProcessContent[address](ctx)
	fmt.Printf("Address : %v\n", l)
	fmt.Printf("Ok      : %v\n", s.Ok())

	//Output:
	//Address : {  }
	//Ok      : false
	//Address : {  }
	//Ok      : false
}
