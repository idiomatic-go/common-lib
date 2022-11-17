package vhost

import (
	"fmt"
	"github.com/idiomatic-go/common-lib/eventing"
)

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
