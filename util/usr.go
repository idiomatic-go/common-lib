package util

import "context"

// Niladic - types for functions that have no parameters
type Niladic func()
type NiladicStatus func() error
type NiladicResponse func() any

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
