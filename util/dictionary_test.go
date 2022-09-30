package util

import "fmt"

func ExampleInvertedDictionary() {
	dict := CreateInvertedDictionary(false)
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
}
