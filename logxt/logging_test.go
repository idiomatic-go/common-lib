package logxt_test

import (
	"github.com/idiomatic-go/common-lib/logxt"
)

func ExamplePrint() {
	logxt.Print("test string")
	logxt.Printf("using a specifier : %v", true)

	// Output:
}

func ExampleDebug() {
	logxt.Debug("test string")

	// Output:
}

func ExampleDebugEnabled() {
	logxt.ToggleDebug(true)

	logxt.Debug("test string")
	logxt.Debugf("this is a specified fmt : %v", 45)

	//Output:
	//test string
	//this is a specified fmt : 45
}
