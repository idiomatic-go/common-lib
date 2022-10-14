package vhost

import (
	"fmt"
	"time"
)

func _ExampleStatusUpdate() {
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

func ExampleValidToSend() {
	//item := struct{}{}
	depUri := "test:dependent"
	uri := "progresql:main"
	sent := make(list)
	dir := createSyncMap()

	// Test nil inputs
	ok, err := validToSend(nil, nil, nil)
	fmt.Printf("Test - {nil} : %v %v\n", ok, err)

	// Test entry with no dependents, should be able to send
	e := &entry{uri: uri, c: nil, dependents: nil, startupStatus: 0}
	ok, err = validToSend(sent, e, dir)
	fmt.Printf("Test - {Empty Dependents} : %v %v\n", ok, err)

	// Test entry with dependents not in sent list
	e = &entry{uri: uri, c: nil, dependents: []string{depUri, "test:uri2"}, startupStatus: 0}
	ok, err = validToSend(sent, e, dir)
	fmt.Printf("Test - {Dependents Not In Sent List} : %v %v\n", ok, err)

	// Test entry with one dependent in sent list, target package not found
	e = &entry{uri: uri, c: nil, dependents: []string{depUri}, startupStatus: 0}
	sent.Add(depUri)
	ok, err = validToSend(sent, e, dir)
	fmt.Printf("Test - {One Dependent In Sent List - Target Package Not Found} : %v %v\n", ok, err)

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
	// ToSend : false invalid argument for validToSend() : one of list, entry or directory is nil
	// ToSend [Empty Dependents]: true <nil>
	// ToSend [Dependents Not In Sent List]: false <nil>
	// ToSend [One Dependent In Sent List - Target Package Not Found]: false dependency not fufilled, package entry not found: test:dependent
	// ToSend [One Dependent In Sent List - Target Package Not Started]: false dependency not fufilled, startup has failed for package: test:dependent
	// ToSend [All Dependents In Sent List And Startup Successful]: true <nil>
	// ToSend [One Dependent In Sent List]: false dependency not fufilled, startup has failed for package: test:dependent

}
