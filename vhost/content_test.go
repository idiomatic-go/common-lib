package vhost

import (
	"fmt"
)

func ExampleAccessCredentialsSuccess() {
	msg := CreateCredentialsMessage("to:uri", "event", "sender", Credentials(func() (username string, password string, err error) { return "", "", nil }))
	fmt.Printf("Credentials Fn : %v\n", AccessCredentials(&msg) != nil)

	//Output:
	// Credentials Fn : true
}

// Need to cast as adding content via any
func ExampleAccessCredentialsSlice() {
	msg := CreateMessage("to:uri", "event", "sender", 0, "first content")
	AddContent(&msg, Credentials(func() (username string, password string, err error) { return "", "", nil }))
	fmt.Printf("Credentials Fn : %v\n", AccessCredentials(&msg) != nil)

	//Output:
	// Credentials Fn : true
}
