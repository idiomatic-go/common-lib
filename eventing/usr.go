package eventing

import "time"

const (
	StartupEvent   = "event:startup"
	ShutdownEvent  = "event:shutdown"
	StartupPause   = "event:pause"
	ShutdownResume = "event:resume"
	ErrorEvent     = "event:error"
	RestartEvent   = "event:restart"
	FailoverEvent  = "event:failover"
	FailbackEvent  = "event:failback"
	PingEvent      = "event:ping"
	ProfileEvent   = "event:profile"
	VirtualHost    = "vhost"

	StatusNotProvided = -100
	StatusOk          = 0  // Need to sink with gRPC Ok
	StatusInternal    = 13 // Need to sink with gRPC Internal
	StatusInProgress  = -3 // Need to sink with vhost StatusInProgress
)

type Message struct {
	To       string
	From     string
	Event    string
	Content  any
	Status   int32
	CreateTS time.Time
}

type MessageHandler func(msg Message)
