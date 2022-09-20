package usr

import (
	"net/http"
)

// Http
var Client *http.Client

func OverrideHttpClient(c *http.Client) {
	if c != nil {
		Client = c
	}
}

type HttpTraceStart func(req *http.Request) HttpTraceFinish

type HttpTraceFinish func(resp *http.Response, err error)

var TraceStart HttpTraceStart

func OverrideHttpTracing(fn HttpTraceStart) {
	TraceStart = fn
}
