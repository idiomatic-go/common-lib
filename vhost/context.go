package vhost

import (
	"context"
	"github.com/idiomatic-go/common-lib/fse"
	"github.com/idiomatic-go/common-lib/logxt"
	"io/fs"
	"strings"
)

type requestid struct{}

var requestKey requestid

// ContextWithRequestId - creates a new Context with a request id
func ContextWithRequestId(ctx context.Context, requestId string) context.Context {
	return context.WithValue(ctx, requestKey, requestId)
}

// ContextRequestId - return the requestId from a Context
func ContextRequestId(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	i := ctx.Value(requestKey)
	if requestId, ok := i.(string); ok {
		return requestId
	}
	return ""
}

type contentid struct{}

var contentKey contentid

// ContextWithFSContent - creates a new Context with FS content
func ContextWithFSContent(ctx context.Context, fs fs.FS, name string) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	if fs == nil {
		logxt.LogDebugf("%v\n", "file system is nil")
		return ctx
	}
	buf, err := fse.ReadFile(fs, name)
	if err != nil {
		logxt.LogDebugf("file system read error : %v\n", err)
		return ctx
	}
	return context.WithValue(ctx, contentKey, fse.Entry{Name: strings.ToLower(name), Content: buf})
}

// ContextWithContent - creates a new Context with content
func ContextWithContent(ctx context.Context, content any) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	if content == nil {
		logxt.LogDebugf("%v\n", "content is nil")
		return ctx
	}
	return context.WithValue(ctx, contentKey, content)
}

func ContextContent(ctx context.Context) any {
	if ctx == nil {
		return nil
	}
	i := ctx.Value(contentKey)
	if IsNil(i) {
		return nil
	}
	return i
}

func IsContextContent(ctx context.Context) bool {
	return IsDevEnvironment() && ContextContent(ctx) != nil
}
