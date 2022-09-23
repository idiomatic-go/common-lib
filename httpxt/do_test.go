package httpxt

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func isInvalidArgument(err error) bool {
	return err != nil && strings.HasPrefix(err.Error(), "invalid argument:")
}

func ExampleDoNilRequest() {
	_, err := Do(nil)
	fmt.Println(isInvalidArgument(err))
	fmt.Println(err)

	//Output:
	// true
	// invalid argument: Request is nil
}

func ExampleDoInvalidScheme() {
	req, _ := http.NewRequest(http.MethodGet, "ftp:test", nil)
	_, err := Do(req)
	fmt.Println(isInvalidArgument(err))
	fmt.Println(err)

	//Output:
	// true
	// invalid argument: URL scheme is not supported [ftp]
}

func ExampleDoHttpError() {
	req, _ := http.NewRequest(http.MethodGet, "echo://www.somestupidname.com?httpError=true", nil)
	resp, err := Do(req)
	fmt.Printf("Error : %v\n", err)
	fmt.Printf("Response : %v\n", resp != nil)

	//Output:
	// Error : http: connection has been hijacked
	// Response : false

}

func ExampleDoIOError() {
	req, _ := http.NewRequest(http.MethodGet, "echo://www.somestupidname.com?ioError=true", nil)
	resp, err := Do(req)
	fmt.Printf("Error : %v\n", err)
	fmt.Printf("Response : %v\n", resp != nil)
	fmt.Printf("Status Code : %v\n", resp.StatusCode)
	fmt.Printf("Body : %v\n", resp.Body != nil)

	defer resp.Body.Close()
	buf, ioerr := io.ReadAll(resp.Body)
	fmt.Printf("IOError : %v\n", ioerr)
	fmt.Printf("Body : %v\n", string(buf))

	//Output:
	// Error : <nil>
	// Response : true
	// Status Code : 200
	// Body : true
	// IOError : unexpected EOF
	// Body :
}

func ExampleDoSuccess() {
	var uri = "echo://www.somestupidname.com"
	uri += "?content-type=text/html"
	uri += "&content-length=1234"
	uri += "&body=<html><body><h1>Hello, World</h1></body></html>"
	req, err0 := http.NewRequest(http.MethodGet, uri, nil)
	if err0 != nil {
		fmt.Println("failure")
	}
	resp, err := Do(req)
	fmt.Printf("Error : %v\n", err)
	fmt.Printf("Response : %v\n", resp != nil)
	fmt.Printf("Status Code : %v\n", resp.StatusCode)
	fmt.Printf("Content-Type : %v\n", resp.Header.Get("content-type"))
	fmt.Printf("Content-Length : %v\n", resp.Header.Get("content-length"))
	fmt.Printf("Body : %v\n", resp.Body != nil)

	defer resp.Body.Close()
	buf, ioerr := io.ReadAll(resp.Body)
	fmt.Printf("IOError : %v\n", ioerr)
	fmt.Printf("Body : %v\n", string(buf))

	//Output:
	// Error : <nil>
	// Response : true
	// Status Code : 200
	// Content-Type : text/html
	// Content-Length : 1234
	// Body : true
	// IOError : <nil>
	// Body : <html><body><h1>Hello, World</h1></body></html>

}

func ExampleEchoNil() {
	_, err := createEchoResponse(nil)
	fmt.Println(isInvalidArgument(err))
	fmt.Println(err)

	//Output:
	// true
	// invalid argument: Request is nil
}

func ExampleEchoNoArgs() {
	req, _ := http.NewRequest("", "http://www.somestupidname.com", nil)
	resp, err := createEchoResponse(req)
	fmt.Printf("Error : %v\n", err)
	fmt.Printf("Request : %v\n", req != nil)
	fmt.Printf("Status Code : %v\n", resp.StatusCode)

	//Output:
	// Error : <nil>
	// Request : true
	// Status Code : 200

}

func ExampleEchoHttpError() {
	req, _ := http.NewRequest("", "http://somestupidname.com?httpError=true", nil)
	resp, err := createEchoResponse(req)
	fmt.Printf("Error : %v\n", err)
	fmt.Printf("Response : %v\n", resp != nil)

	//Output:
	// Error : http: connection has been hijacked
	// Response : false
}

func ExampleEchoHttp404() {
	req, _ := http.NewRequest("", "http://somestupidname.com?status=404", nil)
	resp, err := createEchoResponse(req)
	fmt.Printf("Error : %v\n", err)
	fmt.Printf("Response : %v\n", resp != nil)
	fmt.Printf("Status Code : %v\n", resp.StatusCode)

	//Output:
	// Error : <nil>
	// Response : true
	// Status Code : 404

}

func ExampleEchoHeaders() {
	req, _ := http.NewRequest("", "http://somestupidname.com?content-type=application/json&content-length=1234", nil)
	resp, err := createEchoResponse(req)
	fmt.Printf("Error : %v\n", err)
	fmt.Printf("Response : %v\n", resp != nil)
	fmt.Printf("Status Code : %v\n", resp.StatusCode)
	fmt.Printf("Content-Type : %v\n", resp.Header.Get("content-type"))
	fmt.Printf("Content-Length : %v\n", resp.Header.Get("content-length"))

	//Output:
	// Error : <nil>
	// Response : true
	// Status Code : 200
	// Content-Type : application/json
	// Content-Length : 1234

}

func ExampleEchoBody() {
	req, _ := http.NewRequest(http.MethodGet, "http://somestupidname.com?body=this is body content", nil)
	resp, err := createEchoResponse(req)
	fmt.Printf("Error : %v\n", err)
	fmt.Printf("Response : %v\n", resp != nil)
	fmt.Printf("Status Code : %v\n", resp.StatusCode)
	fmt.Printf("Body : %v\n", resp.Body != nil)

	defer resp.Body.Close()
	buf, ioerr := io.ReadAll(resp.Body)
	fmt.Printf("IOError : %v\n", ioerr)
	fmt.Printf("Body : %v\n", string(buf))

	//Output:
	// Error : <nil>
	// Response : true
	// Status Code : 200
	// Body : true
	// IOError : <nil>
	// Body : this is body content

}

func ExampleEchoIOError() {
	req, _ := http.NewRequest(http.MethodGet, "http://somestupidname.com?ioError=true", nil)
	resp, err := createEchoResponse(req)
	fmt.Printf("Error : %v\n", err)
	fmt.Printf("Response : %v\n", resp != nil)
	fmt.Printf("Status Code : %v\n", resp.StatusCode)
	fmt.Printf("Body : %v\n", resp.Body != nil)

	defer resp.Body.Close()
	buf, ioerr := io.ReadAll(resp.Body)
	fmt.Printf("IOError : %v\n", ioerr)
	fmt.Printf("Body : %v\n", string(buf))

	//Output:
	// Error : <nil>
	// Response : true
	// Status Code : 200
	// Body : true
	// IOError : unexpected EOF
	// Body :

}

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
