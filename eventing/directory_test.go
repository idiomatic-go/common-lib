package eventing

import "fmt"

func createTestEntry(uri string, status int32) *entry {
	entry := createEntry(uri, nil)
	entry.msgs.add(CreateMessage(VirtualHost, VirtualHost, StartupEvent, status, nil))
	return entry
}

func ExampleSyncMapInit() {
	uri := "urn:test"

	fmt.Printf("Count : %v\n", Directory.Count())
	d2 := Directory.get(uri)
	fmt.Printf("Entry : %v\n", d2)

	entry := createTestEntry(uri, statusInProgress)
	Directory.put(entry)
	fmt.Printf("Count : %v\n", Directory.Count())
	d2 = Directory.get(uri)
	fmt.Printf("Entry : %v\n", d2)

	//Output:
	//Count : 0
	//Entry : <nil>
	//Count : 1
	//Entry : &{urn:test <nil> [uri1 uri2] 100 {0 0 <nil>} {0 0 <nil>}}

}
