package vhost

import (
	"errors"
	"fmt"
)

func NilEmpty(s string) string {
	if s == "" {
		return "<nil>"
	}
	return s
}

func ExampleStatusOk() {
	sc := NewStatusOk()

	fmt.Printf("Status.Ok()      : %v\n", sc.Ok())
	fmt.Printf("Status.IsError() : %v\n", sc.IsError())
	fmt.Printf("Status.InvalidArgument() : %v\n", sc.InvalidArgument())
	fmt.Printf("Status.Message() : %v\n", NilEmpty(sc.Message()))

	//Output:
	//Status.Ok()      : true
	//Status.IsError() : false
	//Status.InvalidArgument() : false
	//Status.Message() : <nil>
}

func ExampleStatusCode() {
	sc := NewStatusCode(StatusDeadlineExceeded, "request timed out")

	fmt.Printf("Status.Ok()      : %v\n", sc.Ok())
	fmt.Printf("Status.IsError() : %v\n", sc.IsError())
	fmt.Printf("Status.DeadlineExceeded() : %v\n", sc.DeadlineExceeded())
	fmt.Printf("Status.Message() : %v\n", sc.Message())
	fmt.Printf("Status.Error()   : %v\n", NilEmpty(sc.Error()))

	sc = NewStatusCode(StatusDeadlineExceeded, errors.New("request timed out"))

	fmt.Printf("Status.Ok()      : %v\n", sc.Ok())
	fmt.Printf("Status.IsError() : %v\n", sc.IsError())
	fmt.Printf("Status.DeadlineExceeded() : %v\n", sc.DeadlineExceeded())
	fmt.Printf("Status.Message() : %v\n", NilEmpty(sc.Message()))
	fmt.Printf("Status.Error()   : %v\n", NilEmpty(sc.Error()))

	//Output:
	//Status.Ok()      : false
	//Status.IsError() : false
	//Status.DeadlineExceeded() : true
	//Status.Message() : request timed out
	//Status.Error()   : <nil>
	//Status.Ok()      : false
	//Status.IsError() : true
	//Status.DeadlineExceeded() : true
	//Status.Message() : <nil>
	//Status.Error()   : request timed out
}

func ExampleStatusError() {
	sc := NewStatusError(errors.New("this is the FIRST error message"), errors.New("this is the SECOND error message"))
	fmt.Printf("Status.Ok()      : %v\n", sc.Ok())
	fmt.Printf("Status.IsError() : %v\n", sc.IsError())
	fmt.Printf("Status.Message() : %v\n", NilEmpty(sc.Message()))
	fmt.Printf("Status.Errors()  : %v\n", sc.Errors())
	//	fmt.Printf("Status.Cat()     : %v\n", sc.Cat())

	err, ok := sc.(error)
	fmt.Printf("Status(.error)   : [%v] [%v]\n", err, ok)

	sc = NewStatusError(nil)
	fmt.Printf("Status.Ok()      : %v\n", sc.Ok())
	fmt.Printf("Status.IsError() : %v\n", sc.IsError())
	fmt.Printf("Status.Message() : %v\n", NilEmpty(sc.Message()))
	fmt.Printf("Status.Errors()  : %v\n", sc.Errors())
	//fmt.Printf("Status.Cat()     : %v\n", NilEmpty(sc.Cat()))

	//Output:
	//Status.Ok()      : false
	//Status.IsError() : true
	//Status.Message() : <nil>
	//Status.Errors()  : [this is the FIRST error message this is the SECOND error message]
	//Status(.error)   : [this is the FIRST error message : this is the SECOND error message] [true]
	//Status.Ok()      : true
	//Status.IsError() : false
	//Status.Message() : <nil>
	//Status.Errors()  : []
}

func ExampleStatusErrorHandled() {
	sc := NewStatusError(errors.New("this is the FIRST error message"), errors.New("this is the SECOND error message"))

	fmt.Printf("Status.Ok()       : %v\n", sc.Ok())
	fmt.Printf("Status.Internal() : %v\n", sc.Internal())
	fmt.Printf("Status.IsError()  : %v\n", sc.IsError())
	fmt.Printf("Status.Errors()   : %v\n", sc.Errors())

	s2 := sc.Handled()
	fmt.Printf("Status.Ok()       : %v\n", s2.Ok())
	fmt.Printf("Status.Internal() : %v\n", s2.Internal())
	fmt.Printf("Status.IsError()  : %v\n", s2.IsError())
	fmt.Printf("Status.Errors()   : %v\n", s2.Errors())

	//Output:
	//Status.Ok()       : false
	//Status.Internal() : true
	//Status.IsError()  : true
	//Status.Errors()   : [this is the FIRST error message this is the SECOND error message]
	//Status.Ok()       : false
	//Status.Internal() : true
	//Status.IsError()  : false
	//Status.Errors()   : []

}

func ExampleStatusErrorHandledNotFound() {
	sc := NewStatusNotFound(errors.New("resource NOT FOUND error message"))

	fmt.Printf("Status.Ok()       : %v\n", sc.Ok())
	fmt.Printf("Status.Internal() : %v\n", sc.Internal())
	fmt.Printf("Status.NotFound() : %v\n", sc.NotFound())
	fmt.Printf("Status.IsError()  : %v\n", sc.IsError())
	fmt.Printf("Status.Errors()   : %v\n", sc.Errors())

	s2 := sc.Handled()
	fmt.Printf("Status.Ok()       : %v\n", s2.Ok())
	fmt.Printf("Status.Internal() : %v\n", s2.Internal())
	fmt.Printf("Status.NotFound() : %v\n", sc.NotFound())
	fmt.Printf("Status.IsError()  : %v\n", s2.IsError())
	fmt.Printf("Status.Errors()   : %v\n", s2.Errors())

	//Output:
	//Status.Ok()       : false
	//Status.Internal() : false
	//Status.NotFound() : true
	//Status.IsError()  : true
	//Status.Errors()   : [resource NOT FOUND error message]
	//Status.Ok()       : false
	//Status.Internal() : false
	//Status.NotFound() : true
	//Status.IsError()  : false
	//Status.Errors()   : []

}
