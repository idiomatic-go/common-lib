package startup

import (
	"fmt"
	"github.com/idiomatic-go/common-lib/util"
	"github.com/idiomatic-go/common-lib/vhost"
	"time"
)

func ExampleCreateToSend() {
	uriNone := "package:none"
	uriOne := "package:one"

	vhost.registerPackageUnchecked(uriNone, nil, nil)
	vhost.registerPackageUnchecked(uriOne, nil, nil)

	m := createToSend(nil)
	fmt.Printf("Test {no override messages} : %v\n", m)

	em := vhost.messageMap{"package:one": {To: "package:one", Event: vhost.StartupEvent, From: "fromUri"}}
	m = createToSend(em)
	fmt.Printf("Test {one override messages} : %v\n", m)

	//Output:
	// Test {no override messages} : map[package:none:{package:none event:startup vhost 0 []} package:one:{package:one event:startup vhost 0 []}]
	// Test {one override messages} : map[package:none:{package:none event:startup vhost 0 []} package:one:{package:one event:startup vhost 0 []}]
}

func ExampleStatusUpdate() {
	uri := "progresql:main"

	vhost.registerPackageUnchecked(uri, nil, []string{"uri1", "uri2"})
	e := vhost.directory.get(uri)
	fmt.Printf("Entry : %v %v\n", e.uri, e.startupStatus) //, e.statusChangeTS.Format(time.RFC3339))

	SendStartupSuccessfulResponse(uri)
	time.Sleep(time.Nanosecond * 1)
	e = vhost.directory.get(uri)
	fmt.Printf("Entry : %v %v\n", e.uri, e.startupStatus) //e.statusChangeTS.Format(time.RFC3339))

	//Output:
	// Entry : progresql:main 0
	// Entry : progresql:main 2

}

func ExampleValidateToSend() {
	uri := "package:none"

	vhost.registerPackageUnchecked(uri, nil, nil)

	toSend := vhost.messageMap{"invalid": {Event: vhost.StartupEvent, From: vhost.HostFrom}}
	err := validateToSend(toSend)
	fmt.Printf("Test - {invalid package uri in message} : %v\n", err)

	toSend = vhost.messageMap{uri: {Event: vhost.StartupEvent, From: vhost.HostFrom}}
	err = validateToSend(toSend)
	fmt.Printf("Test - {valid package uri in message} : %v\n", err)

	uri2 := "package:one"
	vhost.registerPackageUnchecked(uri2, nil, []string{"package:invalid"})

	toSend = vhost.messageMap{uri: {Event: vhost.StartupEvent, From: vhost.HostFrom}, uri2: {Event: vhost.StartupEvent, From: vhost.HostFrom}}
	err = validateToSend(toSend)
	fmt.Printf("Test - {invalid package uri in dependent} : %v\n", err)

	UnregisterPackage(uri2)
	vhost.registerPackageUnchecked(uri2, nil, []string{"package:none"})

	toSend = vhost.messageMap{"package:none": {Event: vhost.StartupEvent, From: vhost.HostFrom}, "package:one": {Event: vhost.StartupEvent, From: vhost.HostFrom}}
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
	e := &vhost.entry{uri: uri, c: nil, dependents: nil, startupStatus: 0}
	ok := validToSend(sent, e)
	fmt.Printf("Test - {Empty Dependents} : %v\n", ok)

	// Test entry with dependents not in sent list
	e = &vhost.entry{uri: uri, c: nil, dependents: []string{depUri, "test:uri2"}, startupStatus: 0}
	ok = validToSend(sent, e)
	fmt.Printf("Test - {Dependents Not In Sent List} : %v\n", ok)

	// Test entry with all dependents in sent list
	e = &vhost.entry{uri: uri, c: nil, dependents: []string{depUri, testUri}, startupStatus: 0}
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
	toSend := vhost.messageMap{uri: {Event: "test", From: ""}}
	current := vhost.messageMap{}

	ok := getCurrentWork(sent, toSend, current)
	fmt.Printf("Test - {empty directory} : %v\n", ok)

	vhost.registerPackageUnchecked(uri, nil, nil)
	//e := &entry{uri: uri, c: nil, dependents: nil, startupStatus: 0}
	//dir.put(e)
	validToSend = func(sent util.List, entry *vhost.entry) bool {
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
	toSend := vhost.messageMap{uri: {Event: vhost.StartupEvent, From: vhost.HostFrom}, uri2: {Event: vhost.StartupEvent, From: vhost.HostFrom, Status: vhost.StatusSuccessful}}
	current := vhost.messageMap{}

	vhost.registerPackageUnchecked(uri, nil, nil)
	vhost.registerPackageUnchecked(uri2, nil, nil)

	validToSend = func(sent util.List, entry *vhost.entry) bool {
		return true
	}
	fmt.Printf("Test - {valid}   : current : %v  toSend : %v\n", current, toSend)
	ok := getCurrentWork(sent, toSend, current)
	fmt.Printf("Test - {valid}   : current : %v  toSend : %v %v\n", current, toSend, ok)

	toSend = vhost.messageMap{uri: {Event: vhost.StartupEvent, From: vhost.HostFrom}, uri2: {Event: vhost.StartupEvent, From: vhost.HostFrom, Status: vhost.StatusSuccessful}}
	current = vhost.messageMap{}
	validToSend = func(sent util.List, entry *vhost.entry) bool {
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
