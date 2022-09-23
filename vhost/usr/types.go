package usr

import (
	"context"
)

// Environment
const (
	DEV_ENV    = "DEV"
	REVIEW_ENV = "REVIEW"
	TEST_ENV   = "TEST"
	STAGE_ENV  = "STAGE"
	PROD_ENV   = "PROD"
)

var RuntimeEnvKey string = "RUNTIME_ENV"

func OverrideRuntimeEnvKey(k string) {
	if k != "" {
		RuntimeEnvKey = k
	}
}

// Overridable
type EnvValid func() bool

var IsDevEnv EnvValid

func OverrideIsDevEnv(fn EnvValid) {
	if fn != nil {
		IsDevEnv = fn
	}
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

// Messaging
const (
	StartupEvent  = "event:startup"
	ShutdownEvent = "event:shutdown"
	ACKEvent      = "event:ack"
	ErrorEvent    = "event:error"
	HostSender    = "vhost"
)

type Message struct {
	Event   string
	Sender  string
	Content []any
}

type Credentials func() (username string, password string)

// Startup
type Envelope struct {
	Uri string
	Msg *Message
}

// Polling
type Do func(context.Context) *Response

type Response struct {
	Error   error
	Status  any
	Content any
}

// Timer
type Notify func()
