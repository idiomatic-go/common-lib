package startup

import (
	"fmt"
	"github.com/idiomatic-go/common-lib/logxt"
	"github.com/idiomatic-go/common-lib/util"
	"time"
)

func ExampleProcessWorkCyclicDependencyError() {
	testUri := "test:uri"
	testUri2 := "test:uri2"
	sent := make(util.List)

	registerPackageUnchecked(testUri, nil, []string{testUri2})
	registerPackageUnchecked(testUri2, nil, []string{testUri})

	toSend := messageMap{testUri: {Event: "test", From: ""}}
	current := messageMap{}

	err := processWork(sent, toSend, current)
	fmt.Printf("Test - {invalid - cyclic dependency} : %v\n", err)

	//Output:
	// Test - {invalid - cyclic dependency} : Startup failure: unable to find items to work, verify cyclic dependencies

}

func ExampleProcessWorkDependencyMapsWithoutChannels() {
	testUri := "test:uri"
	testUri2 := "test:uri2"
	sent := make(util.List)

	registerPackageUnchecked(testUri, nil, nil)
	registerPackageUnchecked(testUri2, nil, []string{testUri})

	toSend := messageMap{testUri: {To: testUri, Event: "StartupEvent", From: HostFrom, Status: 0}}
	current := messageMap{}

	// Test sending one message
	fmt.Printf("Test - {valid - maps before} : %v %v %v\n", sent, toSend, current)
	err := processWork(sent, toSend, current)
	fmt.Printf("Test - {valid - maps after}  : %v %v %v %v\n", sent, toSend, current, err)

	// Test sending two messages, with the second message having a dependency
	sent = make(util.List)
	current = messageMap{}
	toSend = messageMap{testUri: {To: testUri, Event: "StartupEvent", From: HostFrom, Status: 0}, testUri2: {To: testUri2, Event: "StartupEvent", From: HostFrom, Status: 0}}

	fmt.Printf("Test - {valid - maps before} : %v %v %v\n", sent, toSend, current)
	err = processWork(sent, toSend, current)
	fmt.Printf("Test - {valid - maps after}  : %v %v %v %v\n", sent, toSend, current, err)

	//Output:
	// Test - {valid - maps before} : map[] map[test:uri:{test:uri StartupEvent vhost 0 []}] map[]
	// Test - {valid - maps after}  : map[test:uri:{}] map[] map[test:uri:{test:uri StartupEvent vhost 0 []}] <nil>
	// Test - {valid - maps before} : map[] map[test:uri:{test:uri StartupEvent vhost 0 []} test:uri2:{test:uri2 StartupEvent vhost 0 []}] map[]
	// Test - {valid - maps after}  : map[test:uri:{}] map[test:uri2:{test:uri2 StartupEvent vhost 0 []}] map[test:uri:{test:uri StartupEvent vhost 0 []}] <nil>

}

func _ExampleProcessWorkDependencyDirWithoutChannels() {
	testUri := "test:uri"
	testUri2 := "test:uri2"
	sent := make(util.List)

	registerPackageUnchecked(testUri, nil, nil)
	registerPackageUnchecked(testUri2, nil, nil)

	toSend := messageMap{testUri: {To: testUri, Event: "StartupEvent", From: HostFrom, Status: 0}}
	current := messageMap{}

	// Test sending one message
	fmt.Printf("Test - {valid - dir before} : %v %v\n", directory.get(testUri), directory.get(testUri2))
	processWork(sent, toSend, current)
	fmt.Printf("Test - {valid - dir after}  : %v %v\n", directory.get(testUri), directory.get(testUri2))

	// Test sending two messages
	sent = make(util.List)
	current = messageMap{}
	toSend = messageMap{testUri: {To: testUri, Event: "StartupEvent", From: HostFrom, Status: 0}, testUri2: {To: testUri2, Event: "StartupEvent", From: HostFrom, Status: 0}}
	directory.setStatus(testUri, StatusEmpty)

	fmt.Printf("Test - {valid - dir before} : %v %v\n", directory.get(testUri), directory.get(testUri2))
	processWork(sent, toSend, current)
	fmt.Printf("Test - {valid - dir after}  : %v %v\n", directory.get(testUri), directory.get(testUri2))

	//Output:
	// Test - {valid - dir before} : &{test:uri <nil> [] 0} &{test:uri2 <nil> [] 0}
	// Test - {valid - dir after}  : &{test:uri <nil> [] 1} &{test:uri2 <nil> [] 0}
	// Test - {valid - dir before} : &{test:uri <nil> [] 0} &{test:uri2 <nil> [] 0}
	// Test - {valid - dir after}  : &{test:uri <nil> [] 1} &{test:uri2 <nil> [] 1}

}

func _ExampleProcessWorkDependencyWithChannels() {
	//util.ToggleDebug(true)
	testUri := "test:uri"
	testUri2 := "test:postgresql"
	sent := make(util.List)
	current := messageMap{}
	c := make(chan Message, 100)

	registerPackageUnchecked(testUri, c, nil)
	registerPackageUnchecked(testUri2, c, []string{testUri})
	go receiveTest(c)

	toSend := messageMap{testUri: {To: testUri, Event: "StartupEvent", From: HostFrom, Status: 0}, testUri2: {To: testUri2, Event: "StartupEvent", From: HostFrom, Status: 0}}
	e := directory.get(testUri)
	e2 := directory.get(testUri2)
	fmt.Printf("Test - {entries - before} : %v %v %v %v %v %v\n", e.uri, e.startupStatus, e.statusChangeTS.Format(time.RFC3339), e2.uri, e2.startupStatus, e2.statusChangeTS.Format(time.RFC3339))
	processWork(sent, toSend, current)
	fmt.Printf("Test - {entries - after send} : %v %v %v %v %v %v\n", e.uri, e.startupStatus, e.statusChangeTS.Format(time.RFC3339), e2.uri, e2.startupStatus, e2.statusChangeTS.Format(time.RFC3339))
	time.Sleep(time.Second * 2)
	fmt.Printf("Test - {entries - after recv} : %v %v %v %v %v %v\n", e.uri, e.startupStatus, e.statusChangeTS.Format(time.RFC3339), e2.uri, e2.startupStatus, e2.statusChangeTS.Format(time.RFC3339))

	empty(current)
	processWork(sent, toSend, current)
	fmt.Printf("Test - {entries - after send} : %v %v %v %v %v %v\n", e.uri, e.startupStatus, e.statusChangeTS.Format(time.RFC3339), e2.uri, e2.startupStatus, e2.statusChangeTS.Format(time.RFC3339))
	time.Sleep(time.Second * 2)
	fmt.Printf("Test - {entries - after recv} : %v %v %v %v %v %v\n", e.uri, e.startupStatus, e.statusChangeTS.Format(time.RFC3339), e2.uri, e2.startupStatus, e2.statusChangeTS.Format(time.RFC3339))

	close(c)

	fmt.Printf("Test - {all sent} : %v\n", toSend)

	//Output:
	// Test - {entries - before} : test:uri 0 0001-01-01T00:00:00Z test:postgresql 0 0001-01-01T00:00:00Z
	// Test - {entries - after send} : test:uri 2 2022-10-17T13:58:52-05:00 test:postgresql 0 0001-01-01T00:00:00Z
	// Test - {entries - after recv} : test:uri 2 2022-10-17T13:58:52-05:00 test:postgresql 0 0001-01-01T00:00:00Z
	// Test - {entries - after send} : test:uri 2 2022-10-17T13:58:52-05:00 test:postgresql 1 2022-10-17T13:58:54-05:00
	// Test - {entries - after recv} : test:uri 2 2022-10-17T13:58:52-05:00 test:postgresql 2 2022-10-17T13:58:54-05:00
	// Test - {all sent} : map[]

}

func receiveTest(c chan Message) {
	for {
		select {
		case msg, open := <-c:
			// Exit on a closed channel
			if !open {
				return
			}
			logxt.LogDebug("%v\n", msg)
			SendStartupSuccessfulResponse(msg.To)
		default:
		}
	}
}
