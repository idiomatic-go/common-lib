package startup

import "fmt"

func ExampleSyncMapInit() {
	uri := "urn:test"

	fmt.Printf("Count : %v\n", directory.count())
	d2 := directory.get(uri)
	fmt.Printf("Entry : %v\n", d2)

	entry := &entry{uri: uri, c: nil, dependents: []string{"uri1", "uri2"}, startupStatus: 100}
	directory.put(entry)
	fmt.Printf("Count : %v\n", directory.count())
	d2 = directory.get(uri)
	fmt.Printf("Entry : %v\n", d2)

	//Output:
	// Count : 0
	// Entry : <nil>
	// Count : 1
	// Entry : &{urn:test <nil> [uri1 uri2] 100}

}

func ExampleSyncMapStatus() {
	uri := "urn:test"

	m := &syncMap{m: make(map[string]*entry)}

	entry := &entry{uri: uri}
	m.put(entry)
	e2 := m.get(uri)
	fmt.Printf("Entry : %v\n", e2)

	status := m.getStatus(uri)
	fmt.Printf("Get Status [%v]: %v\n", uri, status)

	status = m.getStatus("invalid")
	fmt.Printf("Get Status [%v]: %v\n", "invalid", status)

	ok := m.setStatus(uri, StatusFailure)
	fmt.Printf("Set Status [%v] : %v %v\n", uri, StatusFailure, ok)

	ok = m.setStatus("invalid", StatusFailure)
	fmt.Printf("Set Status [%v] : %v %v\n", "invalid", StatusFailure, ok)

	status = m.getStatus(uri)
	fmt.Printf("Get Status [%v]: %v\n", uri, status)

	ok = m.isStartupSuccessful(uri)
	fmt.Printf("Startup Successful [%v]: %v\n", uri, ok)

	ok = m.setStatus(uri, StatusSuccessful)
	fmt.Printf("Set Status [%v] : %v %v\n", uri, StatusSuccessful, ok)

	ok = m.isStartupSuccessful(uri)
	fmt.Printf("Startup Successful [%v]: %v\n", uri, ok)

	//Output:
	// Entry : &{urn:test <nil> [] 0}
	// Get Status [urn:test]: 0  true
	// Get Status [invalid]: 0  false
	// Set Status [urn:test] : 3 true
	// Set Status [invalid] : 3 false
	// Get Status [urn:test]: 3  true
	// Startup Successful [urn:test]: false
	// Set Status [urn:test] : 2 true
	// Startup Successful [urn:test]: true
}

func _ExampleSyncMapDirectoryStatus() {
	uri := "urn:test"

	m := &syncMap{m: make(map[string]*entry)}

	e := &entry{uri: uri, startupStatus: StatusSuccessful}
	m.put(e)
	e2 := m.get(uri)
	fmt.Printf("Entry [%v] : %v\n", uri, e2)

	e = &entry{uri: "urn:test2", startupStatus: StatusFailure}
	m.put(e)
	e2 = m.get("urn:test2")
	fmt.Printf("Entry [%v] : %v\n", "urn:test2", e2)

	e = &entry{uri: "urn:test3", startupStatus: StatusEmpty}
	m.put(e)
	e2 = m.get("urn:test3")
	fmt.Printf("Entry [%v] : %v\n", "urn:test3", e2)

	s := m.startupFailure()
	fmt.Printf("Startup Failure : %v\n", s)

	s = m.startupInProgress()
	fmt.Printf("Startup In Progress : %v\n", s)

	m.setStatus("urn:test2", StatusInProgress)
	s = m.startupInProgress()
	fmt.Printf("Startup In Progress : %v\n", s)

	//Output:
	// Entry [urn:test] : &{urn:test <nil> [] 2}
	// Entry [urn:test2] : &{urn:test <nil> [] 3}
	// Entry [urn:test3] : &{urn:test <nil> [] 0}
	// Startup Failure : urn:test2
	// Startup In Progress :
	// Startup In Progress : urn:test2

}
