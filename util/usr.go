package util

const (
	QbeNid = "qbe"
)

// Func - type for niladic functions, functions with no parameters
type Func func()

type FuncBool func() bool

// FuncStatus - type for functions with no parameters and a return status
type FuncStatus func() error

// FuncResponse - type for functions with no parameters and a return response
type FuncResponse func() any

type FuncValues func() ([]any, error)

type QbeCell struct {
	Field    string
	Criteria any
}

type URN struct {
	Nid      string
	Nss      string
	RawQuery string
	QbeGrid  []QbeCell
	Err      error
}

// VariableLookup - type used in template.go
type VariableLookup = func(name string) (value string, err error)
