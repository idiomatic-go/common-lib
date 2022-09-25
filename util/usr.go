package util

import "context"

// Dispatch types
type Dispatch func()
type DispatchStatus func() error

type DoPoll func(ctx context.Context) Response

type Response struct {
	content string
}

// Debug flag
var Debug = false

func ToggleDebug(v bool) {
	Debug = v
}

// DebugFmt Overridable
type DebugFmt func(specifier string, v ...any)

type DefaultFmt func(v ...any)

type ContextDefaultFmt func(ctx context.Context, v ...any)

type SpecifiedFmt func(specifier string, v ...any)

type ContextSpecifiedFmt func(ctx context.Context, specifier string, v ...any)
