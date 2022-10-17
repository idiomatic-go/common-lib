package vhost

import (
	"fmt"
	"github.com/idiomatic-go/common-lib/util"
)

func _ExampleProcessWorkCyclicDependencyError() {
	testUri := "test:uri"
	testUri2 := "test:uri2"
	sent := make(util.List)
	dir := createSyncMap()

	e := &entry{uri: testUri, c: nil, dependents: []string{testUri2}, startupStatus: 0}
	dir.put(e)

	e = &entry{uri: testUri2, c: nil, dependents: []string{testUri}, startupStatus: 0}
	dir.put(e)

	toSend := messageMap{testUri: {Event: "test", From: ""}}
	current := messageMap{}

	err := processWork(sent, toSend, current, dir)
	fmt.Printf("Test - {invalid - cyclic dependency} : %v\n", err)

	//Output:
	// Test - {invalid - cyclic dependency} : Startup failure: unable to find items to work, verify cyclic dependencies

}

func _ExampleProcessWorkDependencyWithoutChannels() {
	testUri := "test:uri"
	testUri2 := "test:uri2"
	sent := make(util.List)
	dir := createSyncMap()
 
	e := &entry{uri: testUri, c: nil, dependents: nil, startupStatus: 0}
	dir.put(e)

	e = &entry{uri: testUri2, c: nil, dependents: []string{testUri}, startupStatus: 0}
	dir.put(e)

	toSend := messageMap{testUri: {To: testUri, Event: "StartupEvent", From: HostFrom, Status: 0}}
	current := messageMap{}

	// Test sending one message
	fmt.Printf("Test - {valid - before} : %v %v %v\n", sent, toSend, current)
	err := processWork(sent, toSend, current, dir)
	fmt.Printf("Test - {valid - after}  : %v %v %v %v\n", sent, toSend, current, err)

	// Test sending two messages, with the second message having a dependency
	sent = make(util.List)
	current = messageMap{}
	toSend = messageMap{testUri: {To: testUri, Event: "StartupEvent", From: HostFrom, Status: 0}, testUri2: {To: testUri2, Event: "StartupEvent", From: HostFrom, Status: 0}}

	fmt.Printf("Test - {valid - before} : %v %v %v\n", sent, toSend, current)
	err = processWork(sent, toSend, current, dir)
	fmt.Printf("Test - {valid - after}  : %v %v %v %v\n", sent, toSend, current, err)

	//Output:
	// Test - {valid - before} : map[] map[test:uri:{test:uri StartupEvent vhost 0 []}] map[]
	// Test - {valid - after}  : map[test:uri:{}] map[] map[test:uri:{test:uri StartupEvent vhost 0 []}] <nil>
	// Test - {valid - before} : map[] map[test:uri:{test:uri StartupEvent vhost 0 []} test:uri2:{test:uri2 StartupEvent vhost 0 []}] map[]
	// Test - {valid - after}  : map[test:uri:{}] map[test:uri2:{test:uri2 StartupEvent vhost 0 []}] map[test:uri:{test:uri StartupEvent vhost 0 []}] <nil>

}

func ExampleProcessWorkDependency() {
	testUri := "test:uri"
	testUri2 := "test:uri2"
	sent := make(util.List)
	dir := createSyncMap()

	e := &entry{uri: testUri, c: nil, dependents: nil, startupStatus: 0}
	dir.put(e)

	e = &entry{uri: testUri2, c: nil, dependents: []string{testUri}, startupStatus: 0}
	dir.put(e)

	toSend := messageMap{testUri: {To: testUri, Event: "StartupEvent", From: HostFrom, Status: 0}}
	current := messageMap{}

	// Test sending one message
	fmt.Printf("Test - {valid - before} : %v %v %v\n", sent, toSend, current)
	err := processWork(sent, toSend, current, dir)
	fmt.Printf("Test - {valid - after}  : %v %v %v %v\n", sent, toSend, current, err)

	// Test sending two messages, with the second message having a dependency
	sent = make(util.List)
	current = messageMap{}
	toSend = messageMap{testUri: {To: testUri, Event: "StartupEvent", From: HostFrom, Status: 0}, testUri2: {To: testUri2, Event: "StartupEvent", From: HostFrom, Status: 0}}

	fmt.Printf("Test - {valid - before} : %v %v %v\n", sent, toSend, current)
	err = processWork(sent, toSend, current, dir)
	fmt.Printf("Test - {valid - after}  : %v %v %v %v\n", sent, toSend, current, err)

	//Output:
	// Test - {valid - before} : map[] map[test:uri:{test:uri StartupEvent vhost 0 []}] map[]
	// Test - {valid - after}  : map[test:uri:{}] map[] map[test:uri:{test:uri StartupEvent vhost 0 []}] <nil>
	// Test - {valid - before} : map[] map[test:uri:{test:uri StartupEvent vhost 0 []} test:uri2:{test:uri2 StartupEvent vhost 0 []}] map[]
	// Test - {valid - after}  : map[test:uri:{}] map[test:uri2:{test:uri2 StartupEvent vhost 0 []}] map[test:uri:{test:uri StartupEvent vhost 0 []}] <nil>

}

