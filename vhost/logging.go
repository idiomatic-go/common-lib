package vhost

import (
	"context"
	"github.com/idiomatic-go/common-lib/logxt"
)

func LogContextPrint(ctx context.Context, v ...any) {
	u := []any{ContextRequestId(ctx)}
	logxt.LogPrintf("%v\n", append(u, v))
}

func LogContextPrintf(ctx context.Context, specifier string, v ...any) {
	logxt.LogPrintf(specifier, v)
}
