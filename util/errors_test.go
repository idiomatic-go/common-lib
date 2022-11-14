package util

import (
	"errors"
	"fmt"
)

func ExampleNewErrors() {
	errs := NewErrors(errors.New("first error"), errors.New("second error"))

	fmt.Printf("IsError: %v\n", errs.IsError())
	fmt.Printf("Error  : %v\n", errs)
	fmt.Printf("Errors : %v\n", errs.Errors())
	errs.Add(errors.New("third error"))
	fmt.Printf("Errors : %v\n", errs.Errors())
	//fmt.Printf("Cat    : %v\n", errs.Cat())

	//Output:
	//IsError: true
	//Error  : first error : second error
	//Errors : [first error second error]
	//Errors : [first error second error third error]
}

func ExampleNewErrorsList() {
	errs := NewErrorsList([]error{errors.New("first error"), errors.New("second error")})

	fmt.Printf("IsError: %v\n", errs.IsError())
	fmt.Printf("Error  : %v\n", errs)
	fmt.Printf("Errors : %v\n", errs.Errors())
	//fmt.Printf("Cat    : %v\n", errs.Cat())

	//Output:
	//IsError: true
	//Error  : first error : second error
	//Errors : [first error second error]

}

func ExampleNewErrorsHandled() {
	errs := NewErrorsList([]error{errors.New("first error"), errors.New("second error")})

	fmt.Printf("IsError: %v\n", errs.IsError())
	fmt.Printf("Error  : %v\n", errs)
	fmt.Printf("Errors : %v\n", errs.Errors())

	errs.Handled()
	fmt.Printf("IsError: %v\n", errs.IsError())
	//fmt.Printf("Error  : %v\n", errs)
	fmt.Printf("Errors : %v\n", errs.Errors())

	//Output:
	//IsError: true
	//Error  : first error : second error
	//Errors : [first error second error]
	//IsError: false
	//Errors : []

}

func ExampleNewErrorsAny() {
	errs := NewErrorsAny("should not be an error")

	fmt.Printf("IsError: %v\n", errs.IsError())
	fmt.Printf("Errors : %v\n", errs.Errors())

	errs = NewErrorsAny(errors.New("this should be an error"))
	fmt.Printf("IsError: %v\n", errs.IsError())
	fmt.Printf("Errors : %v\n", errs.Errors())

	//Output:
	//IsError: false
	//Errors : []
	//IsError: true
	//Errors : [this should be an error]
}
