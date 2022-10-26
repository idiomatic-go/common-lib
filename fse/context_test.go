package fse

import (
	"context"
	"fmt"
)

func ExampleEmbeddedFS() {
	ctx := ContextWithEmbeddedFS(context.Background(), fsys)
	fmt.Printf("Embedded FS : %v\n", ContextEmbeddedFS(ctx))

	//Output:
	// Embedded FS : {0x71d7a0}

}
