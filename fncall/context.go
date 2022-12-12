package fncall

import (
	"context"
	"time"
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

// ContextWithContent - creates a new Context with content
func ContextWithContent(ctx context.Context, content any) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	if content == nil {
		Debugf("%v\n", "content is nil")
		return ctx
	}
	return &valueCtx{ctx, contentKey, content}
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
	if ctx == nil {
		return false
	}
	for {
		switch ctx.(type) {
		case *valueCtx:
			return true
		default:
			return false
		}
	}
}

type valueCtx struct {
	ctx      context.Context
	key, val any
}

func (*valueCtx) Deadline() (deadline time.Time, ok bool) {
	return
}

func (*valueCtx) Done() <-chan struct{} {
	return nil
}

func (*valueCtx) Err() error {
	return nil
}

func (v *valueCtx) Value(key any) any {
	return v.val
}
