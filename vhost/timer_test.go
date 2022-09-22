package vhost

import (
	"github.com/idiomatic-go/common-lib/vhost/usr"
	"time"
)

func ExampleOneTickNoClose() {
	usr.ToggleDebugLogging(true)
	go Timer(false, time.Millisecond*500, nil, nil)
	time.Sleep(time.Second * time.Duration(2))

	//Output:
	// [timer : notify]
	// [timer : finished]
}

func _ExampleMultiTicksAndClose() {
	usr.ToggleDebugLogging(true)
	c := NewStopChannel()
	go Timer(true, time.Millisecond*500, c, nil)
	time.Sleep(time.Second * time.Duration(3))
	StopTimer(c)
	time.Sleep(time.Second * time.Duration(1))

	//Output:
	// [timer : notify]
	// [timer : notify]
	// [timer : notify]
	// [timer : notify]
	// [timer : notify]
	// [timer : notify]
	// [timer : stopped]
}

func _ExampleMultiTicksAndCloseWithEcho() {
	usr.ToggleDebugLogging(true)
	c := NewStopChannel()
	go Timer(true, time.Millisecond*500, c, func() { LogDebug("%s\n", "notify") })
	time.Sleep(time.Second * time.Duration(2))
	StopTimer(c)
	time.Sleep(time.Second * time.Duration(1))

	//Output:
	// [timer : notify]
	// [notify]
	// [timer : notify]
	// [notify]
	// [timer : notify]
	// [notify]
	// [timer : notify]
	// [notify]
	// [timer : stopped]
}
