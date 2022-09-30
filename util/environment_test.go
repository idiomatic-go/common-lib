package util

import (
	"fmt"
	"os"
)

func ExampleGetenv() {
	os.Setenv("key", "value")
	os.Setenv("key-1", "value-1")
	os.Setenv("key-2", "$key-1")
	fmt.Printf("os.Getenv [key]   : %v\n", os.Getenv("key"))
	fmt.Printf("os.Getenv [key-1] : %v\n", os.Getenv("key-1"))
	fmt.Printf("os.Getenv [key-2] : %v\n\n", os.Getenv("key-2"))

	fmt.Printf("Getenv [key]   : %v\n", Getenv("key"))
	fmt.Printf("Getenv [key-1] : %v\n", Getenv("key-1"))
	fmt.Printf("Getenv [key-2] : %v\n", Getenv("key-2"))

	//Output:
	// os.Getenv [key]   : value
	// os.Getenv [key-1] : value-1
	// os.Getenv [key-2] : $key-1
	//
	// Getenv [key]   : value
	// Getenv [key-1] : value-1
	// Getenv [key-2] : value-1
}
