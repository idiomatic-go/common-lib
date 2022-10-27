package httpxt

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
)

// Tracing examples
type pseudoSpan struct {
	SetTag    func(name, value interface{})
	LogFields func(err error)
	Finish    func()
}

func startSpanFromContext(ctx context.Context, operationName string) (pseudoSpan, context.Context) {
	fmt.Printf("startSpanFromContext() : [%v]\n", operationName)
	return pseudoSpan{
			SetTag: func(name, value interface{}) {
				fmt.Printf("SetTag() : [%v] [%v]\n", name, value)
			}, LogFields: func(err error) {
				fmt.Printf("LogFields() : [%v]\n", err)
			}, Finish: func() {
				fmt.Printf("Finish()\n")
			}},
		ctx
}

func inject(s pseudoSpan, req *http.Request) {
	fmt.Printf("inject()\n")
}

func tracingStart(req *http.Request) HttpTraceFinish {
	span, _ := startSpanFromContext(req.Context(), fmt.Sprintf("getRequest : %v", req.URL.Path))
	span.SetTag("http.url", req.URL.String())
	span.SetTag("http.method", req.Method)
	inject(span, req)
	return func(resp *http.Response, err error) {
		if err != nil {
			span.SetTag("error", true)
			span.LogFields(err)
		} else {
			span.SetTag("http.status_code", strconv.Itoa(resp.StatusCode))
		}
		span.Finish()
	}
}

func ExampleTraceOutput() {
	req, _ := http.NewRequest("", "http://google.com/search", nil)
	span, _ := startSpanFromContext(req.Context(), fmt.Sprintf("getRequest : %v", req.URL.Path))
	span.SetTag("http.url", req.URL.String())
	span.SetTag("http.method", req.Method)
	inject(span, req)
	span.Finish()

	//Output:
	// startSpanFromContext() : [getRequest : /search]
	// SetTag() : [http.url] [http://google.com/search]
	// SetTag() : [http.method] [GET]
	// inject()
	// Finish()

}

func ExampleTracingSuccess() {
	OverrideHttpTracing(tracingStart)
	req, _ := http.NewRequest("", "echo://www.google.com/search", nil)
	Do(req)
	OverrideHttpTracing(nil)

	//Output:
	// startSpanFromContext() : [getRequest : /search]
	// SetTag() : [http.url] [echo://www.google.com/search]
	// SetTag() : [http.method] [GET]
	// inject()
	// SetTag() : [http.status_code] [200]
	// Finish()

}

func ExampleTracingErrors() {
	OverrideHttpTracing(tracingStart)
	req, _ := http.NewRequest("", "echo://www.google.com?httpError=true", nil)
	Do(req)
	OverrideHttpTracing(nil)

	//Output:
	// startSpanFromContext() : [getRequest : ]
	// SetTag() : [http.url] [echo://www.google.com?httpError=true]
	// SetTag() : [http.method] [GET]
	// inject()
	// SetTag() : [error] [true]
	// LogFields() : [http: connection has been hijacked]
	// Finish()

}
