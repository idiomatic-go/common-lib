package util

import "context"

// Timer
type Notify func()

type Do func(ctx context.Context) Response

type Response struct {
	content string
}
