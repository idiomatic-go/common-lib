package util

import (
	"errors"
	"fmt"
)

func ExampleNewErrors() {
	errs := NewErrors(errors.New("first error"), errors.New("second error"))

	fmt.Printf("Error  : %v\n", errs)
	fmt.Printf("Errors : %v\n", errs.Errors())
	errs.Add(errors.New("third error"))
	fmt.Printf("Errors : %v\n", errs.Errors())

	//Output:
	//fail
}
