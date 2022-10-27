package logxt

// ToggleDebug - function to toggle the debug flag
func ToggleDebug(v bool) {
	debug = v
}

// DefaultFmt - logging function type
type DefaultFmt func(v ...any)

// SpecifiedFmt - logging function type
type SpecifiedFmt func(specifier string, v ...any)
