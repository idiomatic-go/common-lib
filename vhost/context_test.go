package vhost

import (
	"context"
	"fmt"
)

func ExampleRequestId() {
	ctx := ContextWithRequestId(context.Background(), "requestId")
	fmt.Printf("RequestId : %v\n", ContextRequestId(ctx))

	//Output:
	//RequestId : requestId

}

func ExampleAnyContent() {
	status := NewStatusOk()
	ctx := ContextWithAnyContent(context.Background(), status)
	s := ContextAnyContent(ctx)
	fmt.Printf("Status : %v\n", s)

	//Output:
	//Status :
}
