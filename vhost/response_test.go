package vhost

import "fmt"

func ExampleNewResponse() {
	var sc Status
	resp := NewResponseContent(nil, sc)

	fmt.Printf("Nil         : %v\n", resp.IsContentNil())
	fmt.Printf("Serialized  : %v\n", resp.IsContentSerialized())
	//fmt.Printf("Response    : %v\n", resp)

	resp = NewResponseContent(nil, "string content")
	fmt.Printf("Nil         : %v\n", resp.IsContentNil())
	fmt.Printf("Serialized  : %v\n", resp.IsContentSerialized())
	//fmt.Printf("Response    : %v\n", resp)
	buf, ok := resp.ContentBytes()
	fmt.Printf("Bytes       : %v %v\n", ok, string(buf))

	//Output:
	//Nil         : true
	//Serialized  : false
	//Nil         : false
	//Serialized  : true
	//Bytes       : true string content
}
