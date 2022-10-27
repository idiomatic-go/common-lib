package vhost

import "context"

type requestid struct{}

var key requestid

// ContextWithRequestId - creates a new Context with a request id
func ContextWithRequestId(ctx context.Context, requestId string) context.Context {
	return context.WithValue(ctx, key, requestId)
}

// ContextRequestId - return the requestId from a Context
func ContextRequestId(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	i := ctx.Value(key)
	if requestId, ok := i.(string); ok {
		return requestId
	}
	return ""
}
