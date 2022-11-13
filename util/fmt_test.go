package util

import (
	"fmt"
	"time"
)

func ExampleTrimDoubleSpace() {
	s := "' this is an example of do  uble spaces  '"

	fmt.Printf("String : %v\n", TrimDoubleSpace(s))

	//Output:
	//String : ' this is an example of do uble spaces '
}

func ExampleFmtTimestamp() {
	s := FmtTimestamp(time.Now())

	fmt.Printf("Timestamp : %v\n", s)

	//Output:
	//fail
}
