package fse

import (
	"context"
	"io/fs"
)

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

type contentid struct{}

var contentkey contentid

// ContextWithEmbeddedContent - creates a new Context with an embedded content
func ContextWithEmbeddedContent(ctx context.Context, name string) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, contentkey, name)
}

// ContextEmbeddedContent - return the embedded content
func ContextEmbeddedContent(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	i := ctx.Value(contentkey)
	if name, ok := i.(string); ok {
		return name
	}
	return ""
}
