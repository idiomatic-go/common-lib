package fse

import (
	"context"
	"github.com/idiomatic-go/common-lib/logxt"
	"io/fs"
)

/*
type fsid struct{}

var fskey fsid

// ContextWithEmbeddedFS - creates a new Context with an embedded FS
func ContextWithEmbeddedFS(ctx context.Context, fsys fs.FS) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, fskey, fsys)
}

// ContextEmbeddedFS - return the embedded FS from a Context
func ContextEmbeddedFS(ctx context.Context) fs.FS {
	if ctx == nil {
		return nil
	}
	i := ctx.Value(fskey)
	if fsys, ok := i.(fs.FS); ok {
		return fsys
	}
	return nil
}
*/

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
	return context.WithValue(ctx, entrykey, Entry{Name: name, Content: buf})
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
