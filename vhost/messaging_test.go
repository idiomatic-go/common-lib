package vhost

import (
	"fmt"
)

func ExampleCreateMessage() {
	msg := CreateMessage("event", "sender", 0, "content")
	fmt.Printf("Content : %v", len(msg.Content))

	//Output:
	// Content : 1

}
