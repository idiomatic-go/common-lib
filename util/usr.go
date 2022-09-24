package util

import "context"

// Dispatch types
type Dispatch func()
type DispatchStatus func() error

type Do func(ctx context.Context) Response

type Response struct {
	content string
}
