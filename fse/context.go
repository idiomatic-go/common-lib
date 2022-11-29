package fse

import (
	"context"
	"github.com/idiomatic-go/common-lib/logxt"
	"io/fs"
	"strings"
)

type entryid struct{}

var entrykey entryid

// ContextWithContent - creates a new Context with content
func ContextWithContent(ctx context.Context, fs fs.FS, name string) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	if fs == nil {
		logxt.LogDebugf("%v", "file system is nil")
		return ctx
	}
	buf, err := ReadFile(fs, name)
	if err != nil {
		logxt.LogDebugf("file system read error : %v", err)
		return ctx
	}
	return context.WithValue(ctx, entrykey, Entry{Name: strings.ToLower(name), Content: buf})
}

// ContextContent - return the content
func ContextContent(ctx context.Context) *Entry {
	if ctx == nil {
		return nil
	}
	i := ctx.Value(entrykey)
	if buf, ok := i.(Entry); ok {
		return &buf
	}
	return nil
}
