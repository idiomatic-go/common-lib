package vhost

import (
	"os"
)

// Environment
const (
	DevEnv         = "DEV"
	EnvTemplateVar = "{env}"
	RuntimeEnvKey  = "RUNTIME_ENV"
)

// OverrideRuntimeEnvKey - allows configuration
func OverrideRuntimeEnvKey(k string) {
	if k != "" {
		runtimeKey = k
	}
}

// GetEnv - function to get the runtime environment
func GetEnv() string {
	return os.Getenv(runtimeKey)
}

// EnvValid - type to allow override of dev environment determination
type EnvValid func() bool

// OverrideIsDevEnv - function to override dev environment determination
func OverrideIsDevEnv(fn EnvValid) {
	if fn != nil {
		IsDevEnv = fn
	}
}
