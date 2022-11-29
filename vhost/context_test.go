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
	if IsContextContent(ctx) {
		fmt.Printf("Status : %v\n", ContextAnyContent(ctx))
	}

	//Output:
	//Status :
}
