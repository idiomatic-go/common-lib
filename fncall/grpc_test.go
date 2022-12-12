package fncall

import (
	"fmt"
)

func ExampleGrpcStatus() {
	var s gRPCStatus

	s = NewgRPCStatus(0, "")
	fmt.Printf("Code    : %v\n", s.Code())
	fmt.Printf("Message : %v\n", NilEmpty(s.Message()))

	s = NewgRPCStatus(StatusNotFound, "test message")
	fmt.Printf("Code    : %v\n", s.Code())
	fmt.Printf("Message : %v\n", NilEmpty(s.Message()))

	s = NewgRPCStatus(StatusAlreadyExists, 1001)
	fmt.Printf("Code    : %v\n", s.Code())
	fmt.Printf("Message : %v\n", NilEmpty(s.Message()))

	//Output:
	//Code    : 0
	//Message : <nil>
	//Code    : 5
	//Message : test message
	//Code    : 6
	//Message : <nil>
}
