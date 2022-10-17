package vhost

import (
	"fmt"
	"github.com/idiomatic-go/common-lib/util"
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

func ExampleProcessWorkDependencyDirWithoutChannels() {
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
	testUri := "test:uri"
	testUri2 := "test:uri2"
	sent := make(util.List)

	registerPackageUnchecked(testUri, nil, nil)
	registerPackageUnchecked(testUri2, nil, []string{testUri})

	toSend := messageMap{testUri: {To: testUri, Event: "StartupEvent", From: HostFrom, Status: 0}}
	current := messageMap{}

	// Test sending one message
	fmt.Printf("Test - {valid - dir before} : %v %v\n", directory.get(testUri), directory.get(testUri2))
	processWork(sent, toSend, current)
	fmt.Printf("Test - {valid - dir  after}  : %v %v\n", directory.get(testUri), directory.get(testUri2))

	// Test sending two messages, with the second message having a dependency
	sent = make(util.List)
	current = messageMap{}
	toSend = messageMap{testUri: {To: testUri, Event: "StartupEvent", From: HostFrom, Status: 0}, testUri2: {To: testUri2, Event: "StartupEvent", From: HostFrom, Status: 0}}

	fmt.Printf("Test - {valid - maps before} : %v %v\n", directory.get(testUri), directory.get(testUri2))
	processWork(sent, toSend, current)
	fmt.Printf("Test - {valid - maps after}  : %v %v\n", directory.get(testUri), directory.get(testUri2))

	//Output:
	// Test - {valid - maps before} : map[] map[test:uri:{test:uri StartupEvent vhost 0 []}] map[]
	// Test - {valid - maps after}  : map[test:uri:{}] map[] map[test:uri:{test:uri StartupEvent vhost 0 []}] <nil>
	// Test - {valid - maps before} : map[] map[test:uri:{test:uri StartupEvent vhost 0 []} test:uri2:{test:uri2 StartupEvent vhost 0 []}] map[]
	// Test - {valid - maps after}  : map[test:uri:{}] map[test:uri2:{test:uri2 StartupEvent vhost 0 []}] map[test:uri:{test:uri StartupEvent vhost 0 []}] <nil>

}
