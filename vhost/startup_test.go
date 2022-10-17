package vhost

import (
	"fmt"
	"github.com/idiomatic-go/common-lib/util"
	"time"
)

func ExampleCreateToSend() {
	uriNone := "package:none"
	uriOne := "package:one"

	registerPackageUnchecked(uriNone, nil, nil)
	registerPackageUnchecked(uriOne, nil, nil)

	m := createToSend(nil)
	fmt.Printf("Test {no override messages} : %v\n", m)

	em := messageMap{"package:one": {To: "package:one", Event: StartupEvent, From: "fromUri"}}
	m = createToSend(em)
	fmt.Printf("Test {one override messages} : %v\n", m)

	//Output:
	// Test {no override messages} : map[package:none:{package:none event:startup vhost 0 []} package:one:{package:one event:startup vhost 0 []}]
	// Test {one override messages} : map[package:none:{package:none event:startup vhost 0 []} package:one:{package:one event:startup vhost 0 []}]
}

func ExampleStatusUpdate() {
	uri := "progresql:main"

	registerPackageUnchecked(uri, nil, []string{"uri1", "uri2"})
	e := directory.get(uri)
	fmt.Printf("Entry : %v\n", e)

	SendStartupSuccessfulResponse(uri)
	time.Sleep(time.Nanosecond * 1)
	e = directory.get(uri)
	fmt.Printf("Entry : %v\n", e)

	//Output:
	// Entry : &{progresql:main <nil> [uri1 uri2] 0}
	// Entry : &{progresql:main <nil> [uri1 uri2] 2}
}

func ExampleValidateToSend() {
	uri := "package:none"

	registerPackageUnchecked(uri, nil, nil)

	toSend := messageMap{"invalid": {Event: StartupEvent, From: HostFrom}}
	err := validateToSend(toSend)
	fmt.Printf("Test - {invalid package uri in message} : %v\n", err)

	toSend = messageMap{uri: {Event: StartupEvent, From: HostFrom}}
	err = validateToSend(toSend)
	fmt.Printf("Test - {valid package uri in message} : %v\n", err)

	uri2 := "package:one"
	registerPackageUnchecked(uri2, nil, []string{"package:invalid"})

	toSend = messageMap{uri: {Event: StartupEvent, From: HostFrom}, uri2: {Event: StartupEvent, From: HostFrom}}
	err = validateToSend(toSend)
	fmt.Printf("Test - {invalid package uri in dependent} : %v\n", err)

	UnregisterPackage(uri2)
	registerPackageUnchecked(uri2, nil, []string{"package:none"})

	toSend = messageMap{"package:none": {Event: StartupEvent, From: HostFrom}, "package:one": {Event: StartupEvent, From: HostFrom}}
	err = validateToSend(toSend)
	fmt.Printf("Test - {valid package uri in dependent} : %v\n", err)

	//Output:
	// Test - {invalid package uri in message} : directory entry does not exist for package uri: invalid
	// Test - {valid package uri in message} : <nil>
	// Test - {invalid package uri in dependent} : directory entry does not exist for dependent package uri: package:invalid
	// Test - {valid package uri in dependent} : <nil>
}

func ExampleValidToSend() {
	depUri := "test:dependent"
	testUri := "test:uri2"
	uri := "progresql:main"
	sent := make(util.List)

	// Test entry with no dependents, should be able to send
	e := &entry{uri: uri, c: nil, dependents: nil, startupStatus: 0}
	ok := validToSend(sent, e)
	fmt.Printf("Test - {Empty Dependents} : %v\n", ok)

	// Test entry with dependents not in sent list
	e = &entry{uri: uri, c: nil, dependents: []string{depUri, "test:uri2"}, startupStatus: 0}
	ok = validToSend(sent, e)
	fmt.Printf("Test - {Dependents Not In Sent List} : %v\n", ok)

	// Test entry with all dependents in sent list
	e = &entry{uri: uri, c: nil, dependents: []string{depUri, testUri}, startupStatus: 0}
	sent.Add(depUri)
	sent.Add(testUri)
	ok = validToSend(sent, e)
	fmt.Printf("Test - {All Dependents In Sent List} : %v\n", ok)

	//Output:
	// Test - {Empty Dependents} : true
	// Test - {Dependents Not In Sent List} : false
	// Test - {All Dependents In Sent List} : true

}

func ExampleGetCurrentWorkError() {
	uri := "progresql:main"
	sent := make(util.List)
	toSend := messageMap{uri: {Event: "test", From: ""}}
	current := messageMap{}

	ok := getCurrentWork(sent, toSend, current)
	fmt.Printf("Test - {empty directory} : %v\n", ok)

	registerPackageUnchecked(uri, nil, nil)
	//e := &entry{uri: uri, c: nil, dependents: nil, startupStatus: 0}
	//dir.put(e)
	validToSend = func(sent util.List, entry *entry) bool {
		return false
	}

	ok = getCurrentWork(sent, toSend, current)
	fmt.Printf("Test - {validToSend error} : %v\n", ok)

	//Output:
	// Test - {empty directory} : false
	// Test - {validToSend error} : false

}

func ExampleGetCurrentWork() {
	uri := "progresql:main"
	uri2 := "awssql:main"
	sent := make(util.List)
	toSend := messageMap{uri: {Event: StartupEvent, From: HostFrom}, uri2: {Event: StartupEvent, From: HostFrom, Status: StatusSuccessful}}
	current := messageMap{}

	registerPackageUnchecked(uri, nil, nil)
	registerPackageUnchecked(uri2, nil, nil)

	validToSend = func(sent util.List, entry *entry) bool {
		return true
	}
	fmt.Printf("Test - {valid}   : current : %v  toSend : %v\n", current, toSend)
	ok := getCurrentWork(sent, toSend, current)
	fmt.Printf("Test - {valid}   : current : %v  toSend : %v %v\n", current, toSend, ok)

	toSend = messageMap{uri: {Event: StartupEvent, From: HostFrom}, uri2: {Event: StartupEvent, From: HostFrom, Status: StatusSuccessful}}
	current = messageMap{}
	validToSend = func(sent util.List, entry *entry) bool {
		return false
	}
	fmt.Printf("Test - {invalid} : current : %v  toSend : %v\n", current, toSend)
	ok = getCurrentWork(sent, toSend, current)
	fmt.Printf("Test - {invalid} : current : %v  toSend : %v %v\n", current, toSend, ok)

	//Output:
	// Test - {valid}   : current : map[]  toSend : map[awssql:main:{ event:startup vhost 2 []} progresql:main:{ event:startup vhost 0 []}]
	// Test - {valid}   : current : map[awssql:main:{ event:startup vhost 2 []} progresql:main:{ event:startup vhost 0 []}]  toSend : map[] true
	// Test - {invalid} : current : map[]  toSend : map[awssql:main:{ event:startup vhost 2 []} progresql:main:{ event:startup vhost 0 []}]
	// Test - {invalid} : current : map[]  toSend : map[awssql:main:{ event:startup vhost 2 []} progresql:main:{ event:startup vhost 0 []}] false

}

func _ExampleStartupInvalid() {

}
