package vhost

import (
	"context"
	"fmt"
	"github.com/idiomatic-go/common-lib/logxt"
	"log"
)

var debug = false

func init() {
	//logxt.SetFlags(logxt.Ldate | logxt.Ltime | logxt.Lmicroseconds | logxt.Llongfile | logxt.LUTC)
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.LUTC)
}

var LogDebug DefaultFmt = func(v ...any) {
	if debug {
		fmt.Print(v)
	}
}

var LogDebugf SpecifiedFmt = func(specifier string, v ...any) {
	if debug {
		fmt.Printf(specifier, v)
	}
}

// ContextDefaultFmt - logging function type
type ContextDefaultFmt func(ctx context.Context, v ...any)

// ContextSpecifiedFmt - logging function type
type ContextSpecifiedFmt func(ctx context.Context, specifier string, v ...any)

var LogContextPrint ContextDefaultFmt = func(ctx context.Context, v ...any) {
	u := []any{ContextRequestId(ctx)}
	logxt.LogPrintf("%v", append(u, v))
}

var LogContextPrintf ContextSpecifiedFmt = func(ctx context.Context, specifier string, v ...any) {
	logxt.LogPrintf(specifier, v)
}
