package util

import (
	"fmt"
)

func ExampleList() {
	uri := "www.google.com/search"
	links := make(List)

	links.Add(uri)
	fmt.Printf("Test - {Added} : %v\n", links)

	fmt.Printf("Test - {Is Empty} : %v\n", links.IsEmpty())

	fmt.Printf("Test - {Item} : %v\n", links[uri])

	fmt.Printf("Test - {Contains} : %v\n", links.Contains(uri))

	fmt.Printf("Test - {Contains invalid} : %v\n", links.Contains("invalid"))

	links.Remove(uri)
	fmt.Printf("Test - {Remove item} : %v\n", links)

	fmt.Printf("Test - {Is Empty} : %v\n", links.IsEmpty())

	//Output:
	// Test - {Added} : map[www.google.com/search:{}]
	// Test - {Is Empty} : false
	// Test - {Item} : {}
	// Test - {Contains} : true
	// Test - {Contains invalid} : false
	// Test - {Remove item} : map[]
	// Test - {Is Empty} : true

}

func _ExampleSyncList() {
	uri := "www.google.com/search"
	links := new(SyncList)

	links.Add(uri)
	fmt.Printf("Test - {Added} : %v\n", links)
	fmt.Printf("Test - {Contains} : %v\n", links.Contains(uri))
	fmt.Printf("Test - {Contains invalid} : %v\n", links.Contains("invalid"))

	links.Remove(uri)
	fmt.Printf("Test - {Remove item} : %v\n", links)

	//Output:
	// Test - {Added} : map[www.google.com/search:{}]
	// Test - {Is Empty} : false
	// Test - {Item} : {}
	// Test - {Contains} : true
	// Test - {Contains invalid} : false
	// Test - {Remove item} : map[]
	// Test - {Is Empty} : true

}
