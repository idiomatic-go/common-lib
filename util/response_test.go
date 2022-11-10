package util

import "fmt"

func ExampleNewResponse() {
	var sc StatusCode
	resp := NewResponse(nil, sc)

	fmt.Printf("Nil         : %v\n", resp.IsContentNil())
	fmt.Printf("Serialized  : %v\n", resp.IsContentSerialized())
	fmt.Printf("Response    : %v\n", resp)

	resp = NewResponse(nil, "string content")
	fmt.Printf("Nil         : %v\n", resp.IsContentNil())
	fmt.Printf("Serialized  : %v\n", resp.IsContentSerialized())
	fmt.Printf("Response    : %v\n", resp)
	buf, ok := resp.ContentBytes()
	fmt.Printf("Bytes       : %v %v\n", ok, string(buf))

	//Output:
	//Nil         : true
	//Serialized  : false
	//Response    : &{ <nil>}
	//Nil         : false
	//Serialized  : true
	//Response    : &{ [115 116 114 105 110 103 32 99 111 110 116 101 110 116]}
	//Bytes       : true string content
}
