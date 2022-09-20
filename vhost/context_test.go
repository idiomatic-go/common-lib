package vhost

import (
	"context"
	"fmt"
)

func ExampleRequestId() {
	ctx := ContextWithRequestId(context.Background(), "requestId")
	fmt.Printf("RequestId : %v\n", ContextRequestId(ctx))

	//Output:
	// RequestId : requestId

}
