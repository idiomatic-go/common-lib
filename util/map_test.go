package util

import "fmt"

func ExampleSubset() {
	fmt.Printf("Panic : %v\n", MapSubset(nil, nil) != nil)
	fmt.Printf("Panic : %v\n", MapSubset(nil, make(map[string]string)) != nil)
	m := map[string]string{"key": "first key", "key2": "key two"}
	fmt.Printf("Subset : %v\n", MapSubset(nil, m))
	fmt.Printf("Subset : %v\n", MapSubset([]string{"key2"}, m))
	fmt.Printf("Subset : %v\n", MapSubset([]string{"key4"}, m))
	fmt.Printf("Subset : %v\n", MapSubset([]string{"key2", "key"}, m))

	//Output:
	// Panic : false
	// Panic : false
	// Subset : map[]
	// Subset : map[key2:key two]
	// Subset : map[]
	// Subset : map[key:first key key2:key two]
}
