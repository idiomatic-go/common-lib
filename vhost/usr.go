package vhost

import (
	"os"
)

// Environment
const (
	DevEnv         = "dev"
	EnvTemplateVar = "{env}"
	RuntimeEnvKey  = "RUNTIME_ENV"
)

// OverrideRuntimeEnvKey - allows configuration
func OverrideRuntimeEnvKey(k string) {
	if k != "" {
		runtimeKey = k
	}
}

// GetEnv - function to get the vhost runtime environment
func GetEnv() string {
	s := os.Getenv(runtimeKey)
	if s == "" {
		return DevEnv
	}
	return s
}

// SetEnv - function to set the vhost runtime environment
func SetEnv(s string) {
	os.Setenv(runtimeKey, s)
}

// EnvValid - type to allow override of dev environment determination
type EnvValid func() bool

// OverrideIsDevEnv - function to override dev environment determination
func OverrideIsDevEnv(fn EnvValid) {
	if fn != nil {
		IsDevEnv = fn
	}
}

const (
	DefaultMaxStartupIterations = 4

	StartupEvent  = "event:startup"
	ShutdownEvent = "event:shutdown"
	ACKEvent      = "event:ack"
	ErrorEvent    = "event:error"
	HostFrom      = "vhost"

	StatusEmpty      = int32(0)
	StatusInProgress = int32(1)
	StatusSuccessful = int32(2)
	StatusFailure    = int32(3)
)

func OverrideMaxStartupIterations(count int) {
	if count > 0 {
		maxStartupIterations = count
	}
}

type Message struct {
	To      string // Uri of the destination package
	Event   string
	From    string // Uri of package that is sending the message
	Status  int32
	Content []any
}

type Credentials func() (username string, password string, err error)

type MessageHandler func(msg Message)
