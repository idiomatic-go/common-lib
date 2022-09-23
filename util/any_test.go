package util

import "fmt"

type testStruct struct {
	vers  string
	count int
}

func ExampleIsPointerl() {
	var i any
	var s string
	var data = testStruct{}
	var count int
	var bytes []byte

	fmt.Printf("any : %v\n", IsPointer(i))
	fmt.Printf("int : %v\n", IsPointer(count))
	fmt.Printf("int * : %v\n", IsPointer(&count))
	fmt.Printf("string : %v\n", IsPointer(s))
	fmt.Printf("string * : %v\n", IsPointer(&s))
	fmt.Printf("struct : %v\n", IsPointer(data))
	fmt.Printf("struct * : %v\n", IsPointer(&data))
	fmt.Printf("[]byte : %v\n", IsPointer(bytes))
	//fmt.Printf("Struct * : %v\n", IsPointer(&data))

	//Output:
	// any : false
	// int : false
	// int * : true
	// string : false
	// string * : true
	// struct : false
	// struct * : true
	// []byte : false

}

func _ExampleIsNil() {
	var i any

	fmt.Printf("Nil any : %v", IsNil(i))
	fmt.Printf("Wrapped nil Nil pointer : %v", IsNil(i))

	//Output:
	// Nil pointer : true
}
