package util

import "context"

// Niladic - type for functions with no parameters
type Niladic func()

// NiladicStatus - type for functions with no parameters and a return status
type NiladicStatus func() error

// NiladicResponse - type for functions with no parameters and a return response
type NiladicResponse func() any

// Debug - variable
var debug = false

// ToggleDebug - function to toggle the debug flag
func ToggleDebug(v bool) {
	debug = v
}

// DebugFmt - logging function type
type DebugFmt func(specifier string, v ...any)

// DefaultFmt - logging function type
type DefaultFmt func(v ...any)

// ContextDefaultFmt - logging function type
type ContextDefaultFmt func(ctx context.Context, v ...any)

// SpecifiedFmt - logging function type
type SpecifiedFmt func(specifier string, v ...any)

// ContextSpecifiedFmt - logging function type
type ContextSpecifiedFmt func(ctx context.Context, specifier string, v ...any)
