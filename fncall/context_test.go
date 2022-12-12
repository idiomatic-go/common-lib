package fncall

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

func ExampleContent() {
	status := NewStatusOk()
	ctx := ContextWithContent(context.Background(), status)
	if IsContextContent(ctx) {
		fmt.Printf("Status : %v\n", status) //ContextContent(ctx))
	}

	//Output:
	//Status :
}
