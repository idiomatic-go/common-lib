package util

import "net/url"

// Func - type for niladic functions, functions with no parameters
type Func func()

type FuncBool func() bool

// FuncStatus - type for functions with no parameters and a return status
type FuncStatus func() error

// FuncResponse - type for functions with no parameters and a return response
type FuncResponse func() any

type FuncValues func() ([]any, error)

type UrnComponents struct {
	NSID   string
	NSS    string
	Values url.Values
}
