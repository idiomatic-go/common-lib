package vhost

import (
	"fmt"
	"github.com/idiomatic-go/common-lib/vhost/usr"
)

func ExampleAccessCredentialsSuccess() {
	msg := CreateCredentialsMessage("event", "sender", usr.Credentials(func() (username string, password string) { return "", "" }))
	fmt.Printf("Credentials Fn : %v\n", AccessCredentials(msg) != nil)

	//Output:
	// Credentials Fn : true
}

// Need to cast as adding content via any
func ExampleAccessCredentialsSlice() {
	msg := CreateMessage("event", "sender", "first content")
	AddContent(msg, usr.Credentials(func() (username string, password string) { return "", "" }))
	fmt.Printf("Credentials Fn : %v\n", AccessCredentials(msg) != nil)

	//Output:
	// Credentials Fn : true
}
