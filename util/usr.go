package util

const (
	EnvTemplateVar = "{env}"
)

// VariableLookup - type used in template.go
type VariableLookup = func(name string) (value string, err error)
