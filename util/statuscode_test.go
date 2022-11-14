package util

import (
	"errors"
	"fmt"
)

func ExampleStatusCodeOk() {
	sc := NewStatusCodeOk()

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

func ExampleStatusCodeOptionalNotFound() {
	sc := NewStatusCodeOptionalNotFound(false, "not found string")

	fmt.Printf("StatusCode.Ok()       : %v\n", sc.Ok())
	fmt.Printf("StatusCode.NotFound() : %v\n", sc.NotFound())
	fmt.Printf("StatusCode.IsError()  : %v\n", sc.IsError())
	fmt.Printf("StatusCode.Message()  : %v\n", NilEmpty(sc.Message()))

	sc = NewStatusCodeOptionalNotFound(true, "not found string")

	fmt.Printf("StatusCode.Ok()       : %v\n", sc.Ok())
	fmt.Printf("StatusCode.NotFound() : %v\n", sc.NotFound())
	fmt.Printf("StatusCode.IsError()  : %v\n", sc.IsError())
	fmt.Printf("StatusCode.Message()  : %v\n", NilEmpty(sc.Message()))

	//Output:
	//StatusCode.Ok()       : true
	//StatusCode.NotFound() : false
	//StatusCode.IsError()  : false
	//StatusCode.Message()  : <nil>
	//StatusCode.Ok()       : false
	//StatusCode.NotFound() : true
	//StatusCode.IsError()  : false
	//StatusCode.Message()  : not found string
}

func ExampleStatusCodeNotFound() {
	sc := NewStatusCodeNotFound("database row was not found")

	fmt.Printf("StatusCode.Ok()       : %v\n", sc.Ok())
	fmt.Printf("StatusCode.NotFound() : %v\n", sc.NotFound())
	fmt.Printf("StatusCode.IsError()  : %v\n", sc.IsError())
	fmt.Printf("StatusCode.Message()  : %v\n", sc.Message())

	//Output:
	//StatusCode.Ok()       : false
	//StatusCode.NotFound() : true
	//StatusCode.IsError()  : false
	//StatusCode.Message()  : database row was not found
}

func ExampleStatusCodeError() {
	sc := NewStatusCodeError(errors.New("this is an error message"))
	fmt.Printf("StatusCode.Ok()      : %v\n", sc.Ok())
	fmt.Printf("StatusCode.IsError() : %v\n", sc.IsError())
	fmt.Printf("StatusCode.Message() : %v\n", NilEmpty(sc.Message()))

	err, ok := sc.(error)
	fmt.Printf("StatusCode(.error)   : [%v] [%v]\n", err, ok)

	sc = NewStatusCodeError(nil)
	fmt.Printf("StatusCode.Ok()      : %v\n", sc.Ok())
	fmt.Printf("StatusCode.IsError() : %v\n", sc.IsError())
	fmt.Printf("StatusCode.Message() : %v\n", NilEmpty(sc.Message()))

	err, ok = sc.(error)
	fmt.Printf("StatusCode(.error)   : [%v] [%v]\n", err, ok)

	//Output:
	//StatusCode.Ok()      : false
	//StatusCode.IsError() : true
	//StatusCode.Message() : this is an error message
	//StatusCode(.error)   : [this is an error message] [true]
	//StatusCode.Ok()      : true
	//StatusCode.IsError() : false
	//StatusCode.Message() : <nil>
	//StatusCode(.error)   : [] [true]

}
func ExampleStatusCodeErrors() {
	sc := NewStatusCodeError(errors.New("this is the FIRST error message"), errors.New("this is the SECOND error message"))
	fmt.Printf("StatusCode.Ok()      : %v\n", sc.Ok())
	fmt.Printf("StatusCode.IsError() : %v\n", sc.IsError())
	fmt.Printf("StatusCode.Message() : %v\n", NilEmpty(sc.Message()))
	fmt.Printf("StatusCode.Errors()  : %v\n", sc.Errors())
	fmt.Printf("StatusCode.Cat()     : %v\n", sc.CatErrors())
	//fmt.Printf("StatusCode.Cat()     : %v\n", sc.CatErrors(" "))
	//fmt.Printf("StatusCode.Cat()     : %v\n", sc.CatErrors("|"))

	err, ok := sc.(error)
	fmt.Printf("StatusCode(.error)   : [%v] [%v]\n", err, ok)

	//Output:
	//StatusCode.Ok()      : false
	//StatusCode.IsError() : true
	//StatusCode.Message() : this is the FIRST error message
	//StatusCode.Errors()  : map[0:this is the FIRST error message 1:this is the SECOND error message]
	//StatusCode.Cat()     : 0:this is the FIRST error message 1:this is the SECOND error message
	//StatusCode(.error)   : [this is the FIRST error message] [true]

}

func ExampleStatusErrorsAttrs() {
	sc := NewStatusCodeErrorAttrs(Attr{"first", errors.New("this is the FIRST error message")}, Attr{"second", errors.New("this is the SECOND error message")})
	fmt.Printf("StatusCode.Ok()      : %v\n", sc.Ok())
	fmt.Printf("StatusCode.IsError() : %v\n", sc.IsError())
	fmt.Printf("StatusCode.Error()   : %v\n", NilEmpty(sc.Error()))
	fmt.Printf("StatusCode.Errors()  : %v\n", sc.Errors())

	//err, ok := sc.(error)
	//fmt.Printf("StatusCode(.error)   : [%v] [%v]\n", err, ok)
	sc.AddError("", errors.New("this is the THIRD error message"))
	sc.AddError("fourth", errors.New("this is the FOURTH error message"))

	fmt.Printf("StatusCode.Error()   : %v\n", NilEmpty(sc.Error()))
	fmt.Printf("StatusCode.Errors()  : %v\n", sc.Errors())

	fmt.Printf("StatusCode.GetError(): %v\n", sc.Errors()["second"])

	//Output:
	//StatusCode.Ok()      : false
	//StatusCode.IsError() : true
	//StatusCode.Error()   : this is the FIRST error message
	//StatusCode.Errors()  : map[first:this is the FIRST error message second:this is the SECOND error message]
	//StatusCode.Error()   : this is the FIRST error message
	//StatusCode.Errors()  : map[0:this is the THIRD error message first:this is the FIRST error message fourth:this is the FOURTH error message second:this is the SECOND error message]
	//StatusCode.GetError(): this is the SECOND error message
}

func ExampleStatusInvalidArgument() {
	sc := NewStatusCodeInvalidArgument(errors.New("this is an invalid argument error"))
	fmt.Printf("StatusCode.Ok() : %v\n", sc.Ok())
	fmt.Printf("StatusCode.InvalidArgument() : %v\n", sc.InvalidArgument())
	fmt.Printf("StatusCode.IsError() : %v\n", sc.IsError())
	fmt.Printf("StatusCode.Message() : %v\n", sc.Message())

	err, ok := sc.(error)
	fmt.Printf("StatusCode(.error) : [%v] [%v]\n", err, ok)

	//Output:
	//StatusCode.Ok() : false
	//StatusCode.InvalidArgument() : true
	//StatusCode.IsError() : true
	//StatusCode.Message() : this is an invalid argument error
	//StatusCode(.error) : [this is an invalid argument error] [true]

}

func ExampleStatusDeadlineExceeded() {
	sc := NewStatusCodeDeadlineExceeded(errors.New("this is a deadline exceeded ERROR"))
	fmt.Printf("StatusCode.Ok() : %v\n", sc.Ok())
	fmt.Printf("StatusCode.DeadlineExceeded() : %v\n", sc.DeadlineExceeded())
	fmt.Printf("StatusCode.IsError() : %v\n", sc.IsError())
	fmt.Printf("StatusCode.Message() : %v\n", sc.Message())

	err, ok := sc.(error)
	fmt.Printf("StatusCode(.error) : [%v] [%v]\n", err, ok)

	//Output:
	//StatusCode.Ok() : false
	//StatusCode.DeadlineExceeded() : true
	//StatusCode.IsError() : true
	//StatusCode.Message() : this is a deadline exceeded ERROR
	//StatusCode(.error) : [this is a deadline exceeded ERROR] [true]

}

/*
func ExampleStatusCode() {
	sc := NewStatusCoder(errors.New("this is an error message"))
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


*/
