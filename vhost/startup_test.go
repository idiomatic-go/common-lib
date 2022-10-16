package vhost

import (
	"errors"
	"fmt"
	"github.com/idiomatic-go/common-lib/util"
	"time"
)

func ExampleCreateToSend() {
	dir := createSyncMap()

	e := &entry{uri: "package:none", c: nil, dependents: nil, startupStatus: 0}
	dir.put(e)
	e = &entry{uri: "package:one", c: nil, dependents: nil, startupStatus: 0}
	dir.put(e)

	m := createToSend(nil, dir)
	fmt.Printf("Test {no override messages} : %v\n", m)

	em := messageMap{"package:one": {To: "package:one", Event: StartupEvent, From: "fromUri"}}
	m = createToSend(em, dir)
	fmt.Printf("Test {one override messages} : %v\n", m)

	//Output:
	// Test {no override messages} : map[package:none:{package:none event:startup vhost 0 []} package:one:{package:one event:startup vhost 0 []}]
	// Test {one override messages} : map[package:none:{package:none event:startup vhost 0 []} package:one:{package:one event:startup vhost 0 []}]
}

func ExampleStatusUpdate() {
	uri := "progresql:main"
	entry := &entry{uri: uri, c: nil, dependents: []string{"uri1", "uri2"}, startupStatus: 0}
	directory.put(entry)
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
	dir := createSyncMap()

	e := &entry{uri: "package:none", c: nil, dependents: nil, startupStatus: 0}
	dir.put(e)
	toSend := messageMap{"invalid": {Event: StartupEvent, From: HostFrom}}
	err := validateToSend(toSend, dir)
	fmt.Printf("Test - {invalid package uri in message} : %v\n", err)

	toSend = messageMap{"package:none": {Event: StartupEvent, From: HostFrom}}
	err = validateToSend(toSend, dir)
	fmt.Printf("Test - {valid package uri in message} : %v\n", err)

	e = &entry{uri: "package:one", c: nil, dependents: []string{"package:invalid"}, startupStatus: 0}
	dir.put(e)
	toSend = messageMap{"package:none": {Event: StartupEvent, From: HostFrom}, "package:one": {Event: StartupEvent, From: HostFrom}}
	err = validateToSend(toSend, dir)
	fmt.Printf("Test - {invalid package uri in dependent} : %v\n", err)

	e = &entry{uri: "package:one", c: nil, dependents: []string{"package:none"}, startupStatus: 0}
	dir.put(e)
	toSend = messageMap{"package:none": {Event: StartupEvent, From: HostFrom}, "package:one": {Event: StartupEvent, From: HostFrom}}
	err = validateToSend(toSend, dir)
	fmt.Printf("Test - {valid package uri in dependent} : %v\n", err)

	//Output:
	// Test - {invalid package uri in message} : directory entry does not exist for package uri: invalid
	// Test - {valid package uri in message} : <nil>
	// Test - {invalid package uri in dependent} : directory entry does not exist for dependent package uri: package:invalid
	// Test - {valid package uri in dependent} : <nil>
}

func ExampleValidToSend() {
	depUri := "test:dependent"
	uri := "progresql:main"
	sent := make(util.List)
	dir := createSyncMap()

	// Test entry with no dependents, should be able to send
	e := &entry{uri: uri, c: nil, dependents: nil, startupStatus: 0}
	ok, err := validToSend(sent, e, dir)
	fmt.Printf("Test - {Empty Dependents} : %v %v\n", ok, err)

	// Test entry with dependents not in sent list
	e = &entry{uri: uri, c: nil, dependents: []string{depUri, "test:uri2"}, startupStatus: 0}
	ok, err = validToSend(sent, e, dir)
	fmt.Printf("Test - {Dependents Not In Sent List} : %v %v\n", ok, err)

	// Test entry with one dependent in sent list, target package not found
	//e = &entry{uri: uri, c: nil, dependents: []string{depUri}, startupStatus: 0}
	//sent.Add(depUri)
	//ok, err = validToSend(sent, e, dir)
	//fmt.Printf("Test - {One Dependent In Sent List - Target Package Not Found} : %v %v\n", ok, err)

	// Start the target package
	e = &entry{uri: depUri, c: nil, dependents: nil, startupStatus: StatusEmpty}
	dir.put(e)

	// Test entry with one dependent in sent list, target package not started
	e = &entry{uri: uri, c: nil, dependents: []string{depUri}, startupStatus: 0}
	sent.Add("test:dependent")
	ok, err = validToSend(sent, e, dir)
	fmt.Printf("Test - {One Dependent In Sent List - Target Package Not Started} : %v %v\n", ok, err)

	// Make all dependents valid
	testUri := "test:uri2"
	dir.setStatus(depUri, StatusSuccessful)
	e = &entry{uri: testUri, c: nil, dependents: nil, startupStatus: StatusSuccessful}
	dir.put(e)

	// Test entry with all dependents in sent list
	e = &entry{uri: uri, c: nil, dependents: []string{depUri, testUri}, startupStatus: 0}
	sent.Add(depUri)
	sent.Add(testUri)
	ok, err = validToSend(sent, e, dir)
	fmt.Printf("Test - {All Dependents In Sent List And Startup Successful} : %v %v\n", ok, err)

	//Output:
	// Test - {Empty Dependents} : true <nil>
	// Test - {Dependents Not In Sent List} : false <nil>
	// Test - {One Dependent In Sent List - Target Package Not Started} : false dependency not fufilled, startup has failed for package: test:dependent
	// Test - {All Dependents In Sent List And Startup Successful} : true <nil>

}

func ExampleGetCurrentWorkError() {
	uri := "progresql:main"
	sent := make(util.List)
	dir := createSyncMap()
	toSend := messageMap{uri: {Event: "test", From: ""}}
	current := messageMap{}

	err := getCurrentWork(sent, toSend, current, dir)
	fmt.Printf("Test - {empty directory} : %v\n", err)

	e := &entry{uri: uri, c: nil, dependents: nil, startupStatus: 0}
	dir.put(e)
	validToSend = func(sent util.List, entry *entry, dir *syncMap) (bool, error) {
		return false, errors.New("validToSend error ")
	}

	err = getCurrentWork(sent, toSend, current, dir)
	fmt.Printf("Test - {validToSend error} : %v\n", err)

	//Output:
	// Test - {empty directory} : <nil>
	// Test - {validToSend error} : validToSend error

}

func ExampleGetCurrentWork() {
	uri := "progresql:main"
	uri2 := "awssql:main"
	sent := make(util.List)
	dir := createSyncMap()
	toSend := messageMap{uri: {Event: StartupEvent, From: HostFrom}, uri2: {Event: StartupEvent, From: HostFrom, Status: StatusSuccessful}}
	current := messageMap{}

	e := &entry{uri: uri, c: nil, dependents: nil, startupStatus: StatusEmpty}
	dir.put(e)
	e = &entry{uri: uri2, c: nil, dependents: nil, startupStatus: StatusEmpty}
	dir.put(e)

	validToSend = func(sent util.List, entry *entry, dir *syncMap) (bool, error) {
		return true, nil
	}
	fmt.Printf("Test - {valid}   : current : %v  toSend : %v\n", current, toSend)
	err := getCurrentWork(sent, toSend, current, dir)
	fmt.Printf("Test - {valid}   : current : %v  toSend : %v %v\n", current, toSend, err)

	toSend = messageMap{uri: {Event: StartupEvent, From: HostFrom}, uri2: {Event: StartupEvent, From: HostFrom, Status: StatusSuccessful}}
	current = messageMap{}
	validToSend = func(sent util.List, entry *entry, dir *syncMap) (bool, error) {
		return false, nil
	}
	fmt.Printf("Test - {invalid} : current : %v  toSend : %v\n", current, toSend)
	err = getCurrentWork(sent, toSend, current, dir)
	fmt.Printf("Test - {invalid} : current : %v  toSend : %v %v\n", current, toSend, err)

	//Output:
	// Test - {valid}   : current : map[]  toSend : map[awssql:main:{ event:startup vhost 2 []} progresql:main:{ event:startup vhost 0 []}]
	// Test - {valid}   : current : map[awssql:main:{ event:startup vhost 2 []} progresql:main:{ event:startup vhost 0 []}]  toSend : map[] <nil>
	// Test - {invalid} : current : map[]  toSend : map[awssql:main:{ event:startup vhost 2 []} progresql:main:{ event:startup vhost 0 []}]
	// Test - {invalid} : current : map[]  toSend : map[awssql:main:{ event:startup vhost 2 []} progresql:main:{ event:startup vhost 0 []}] <nil>

}

func _ExampleStartupInvalid() {

}
