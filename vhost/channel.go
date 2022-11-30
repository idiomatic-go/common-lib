package vhost

import "time"

type ClosableChannel struct {
	C chan struct{}
}

// CreateClosableChannel : create a closable channel type, with a minimum channel capacity of 1, otherwise,
//                        the routine will block waiting on a receive from the stop channel
func CreateClosableChannel() *ClosableChannel {
	c := ClosableChannel{C: make(chan struct{}, 1)}
	return &c
}

func (c *ClosableChannel) Close() {
	if c != nil {
		close(c.C)
	}
}

const (
	CommandChannelStart = iota
	CommandChannelStop
	CommandChannelPause
	CommandChannelResume
)

type CommandChannel struct {
	C    chan int
	Tick *time.Duration
}

func CreateCommandChannel(tick time.Duration) *CommandChannel {
	c := CommandChannel{C: make(chan int, 1), Tick: &tick}
	return &c
}

func (c *CommandChannel) NewTicker() *time.Ticker {
	if c != nil {
		return time.NewTicker(*c.Tick)
	}
	return nil
}

func (c *CommandChannel) send(cmd int) {
	if c != nil {
		switch cmd {
		case CommandChannelStart, CommandChannelResume, CommandChannelPause, CommandChannelStop:
			c.C <- cmd
		default:
		}
	}
}

func (c *CommandChannel) Start() {
	if c != nil {
		c.send(CommandChannelStart)
	}
}

func (c *CommandChannel) Stop() {
	if c != nil {
		c.send(CommandChannelStop)
	}
}

func (c *CommandChannel) Pause() {
	if c != nil {
		c.send(CommandChannelPause)
	}
}

func (c *CommandChannel) Resume() {
	if c != nil {
		c.send(CommandChannelResume)
	}
}

type ResponseChannel struct {
	C    chan any
	Tick *time.Duration
}

func CreateResponseChannel(tick time.Duration) *ResponseChannel {
	c := ResponseChannel{C: make(chan any, 1), Tick: &tick}
	return &c
}

func (c *ResponseChannel) NewTicker() *time.Ticker {
	if c != nil {
		return time.NewTicker(*c.Tick)
	}
	return nil
}
