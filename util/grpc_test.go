package util

import "fmt"

func ExampleGrpcStatus() {
	var s gRPCStatus

	s = creategRPCStatus(0, "")
	fmt.Printf("Code  : %v\n", s.Code())

	//Output:
	//fail
}
