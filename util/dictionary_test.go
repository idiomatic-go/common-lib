package util

import "fmt"

func ExampleInvertedDictionary() {
	dict := CreateInvertedDictionary()
	fmt.Printf("Empty : %v\n", dict.IsEmpty())
	val := dict.Lookup("")
	fmt.Printf("Lookup [''] : %v\n\n", val)

	val = dict.Add("")
	fmt.Printf("Add [''] : %v\n", val)
	fmt.Printf("Empty : %v\n", dict.IsEmpty())
	val = dict.Lookup("")
	fmt.Printf("Lookup [''] : %v\n\n", val)

	val = dict.Add("new-key")
	fmt.Printf("Add [new-key] : %v\n", val)
	val = dict.Lookup("new-key")
	fmt.Printf("Lookup [new-key] : %v\n\n", val)

	val = dict.Add("new-key2")
	fmt.Printf("Add [new-key2] : %v\n", val)
	val = dict.Lookup("new-key2")
	fmt.Printf("Lookup [new-key2] : %v\n\n", val)

	var found = false
	var val2 string
	val2, found = dict.InverseLookup(1)
	fmt.Printf("LookupKey [1] : %v %v\n\n", val2, found)

	val2, found = dict.InverseLookup(2)
	fmt.Printf("LookupKey [2] : %v %v\n\n", val2, found)

	val2, found = dict.InverseLookup(3)
	fmt.Printf("LookupKey [3] : %v %v\n\n", val2, found)

	//Output:
	// Empty : true
	// Lookup [''] : -1
	//
	// Add [''] : 0
	// Empty : false
	// Lookup [''] : 0
	//
	// Add [new-key] : 1
	// Lookup [new-key] : 1
	//
	// Add [new-key2] : 2
	// Lookup [new-key2] : 2
	//
	// LookupKey [1] : new-key true
	//
	// LookupKey [2] : new-key2 true
	//
	// LookupKey [3] :  false
}
