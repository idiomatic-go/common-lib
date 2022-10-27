package vhost

import (
	"context"
	"github.com/idiomatic-go/common-lib/logxt"
)

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
