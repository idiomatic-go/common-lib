package util

import (
	"github.com/idiomatic-go/common-lib/vhost"
	"github.com/idiomatic-go/common-lib/vhost/usr"
	"time"
)

func ExampleOneTickNoClose() {
	usr.ToggleDebug(true)
	go Timer(false, time.Millisecond*500, nil, nil)
	time.Sleep(time.Second * time.Duration(2))

	//Output:
	// [timer : notify]
	// [timer : finished]
}

func _ExampleMultiTicksAndClose() {
	usr.ToggleDebug(true)
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
	usr.ToggleDebug(true)
	c := NewStopChannel()
	go Timer(true, time.Millisecond*500, c, func() { vhost.LogDebug("%s\n", "notify") })
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
