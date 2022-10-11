package util

import "context"

// Func - type for niladic functions, functions with no parameters
type Func func()

type FuncBool func() bool

// FuncStatus - type for functions with no parameters and a return status
type FuncStatus func() error

// FuncResponse - type for functions with no parameters and a return response
type FuncResponse func() any

type FuncValues func() ([]any, error)

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
