package util

import (
	"errors"
	"fmt"
)

func ExampleStatusCodeOk() {
	sc := NewStatusOk()

	fmt.Printf("StatusCode.Ok() : %v\n", sc.Ok())
	fmt.Printf("StatusCode.InvalidArgument() : %v\n", sc.InvalidArgument())
	fmt.Printf("StatusCode.IsError() : %v\n", sc.IsError())
	fmt.Printf("StatusCode.Message() : %v\n", sc.Message())

	//Output:
	//StatusCode.Ok() : true
	//StatusCode.InvalidArgument() : false
	//StatusCode.IsError() : false
	//StatusCode.Message() :
}

func ExampleStatusCodeInvalidArgument() {
	sc := NewStatusInvalidArgument("this is an invalid argument message")
	fmt.Printf("StatusCode.Ok() : %v\n", sc.Ok())
	fmt.Printf("StatusCode.InvalidArgument() : %v\n", sc.InvalidArgument())
	fmt.Printf("StatusCode.IsError() : %v\n", sc.IsError())
	fmt.Printf("StatusCode.Message() : %v\n", sc.Message())

	//Output:
	//StatusCode.Ok() : false
	//StatusCode.InvalidArgument() : true
	//StatusCode.IsError() : false
	//StatusCode.Message() : this is an invalid argument message
}

func ExampleStatusCodeError() {
	sc := NewStatusError(errors.New("this is an error message"))
	fmt.Printf("StatusCode.Ok() : %v\n", sc.Ok())
	fmt.Printf("StatusCode.IsError() : %v\n", sc.IsError())
	fmt.Printf("StatusCode.Message() : %v\n", sc.Message())

	err, ok := sc.(error)
	fmt.Printf("StatusCode(.error) : [%v] [%v]\n", err, ok)

	//Output:
	//StatusCode.Ok() : false
	//StatusCode.IsError() : true
	//StatusCode.Message() : status code not provided for errors
	//StatusCode(.error) : [this is an error message] [true]
}
