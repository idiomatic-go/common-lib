package util

import (
	"fmt"
	"time"
)

func ExampleFmtTimestamp() {
	s := FmtTimestamp(time.Now())

	fmt.Printf("Timestamp : %v\n", s)

	//Output:
	//fail
}
