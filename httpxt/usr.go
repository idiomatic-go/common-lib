package httpxt

import "net/http"

// OverrideHttpClient - change client implementation
func OverrideHttpClient(c *http.Client) {
	if c != nil {
		client = c
	}
}

type HttpTraceStart func(req *http.Request) HttpTraceFinish

type HttpTraceFinish func(resp *http.Response, err error)

// OverrideHttpTracing - Enable Http tracing
func OverrideHttpTracing(fn HttpTraceStart) {
	traceStart = fn
}

// ResponseStatus - status from a Http exchange
type ResponseStatus struct {
	BodyIOErr    error
	UnmarshalErr error
	RequestErr   error
	HttpErr      error
	Response     *http.Response
	RequestID    string
}
