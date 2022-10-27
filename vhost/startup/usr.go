package startup

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
