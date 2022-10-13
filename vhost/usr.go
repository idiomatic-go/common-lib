package vhost

// Environment
const (
	DEV_ENV          = "DEV"
	REVIEW_ENV       = "REVIEW"
	TEST_ENV         = "TEST"
	STAGE_ENV        = "STAGE"
	PROD_ENV         = "PROD"
	ENV_TEMPLATE_VAR = "{env}"
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

// Messaging
var MaxStartupIterations = 4

const (
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

type Message struct {
	Event   string
	From    string // Uri of package that is sending the message
	Status  int32
	Content []any
}

type Credentials func() (username string, password string, err error)

// Envelope - struct for startup
type Envelope struct {
	Uri string
	Msg Message
}
