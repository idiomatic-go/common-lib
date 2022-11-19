package vhost

import (
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

// https://github.com/gorilla/handlers/blob/master/logging.go
// LogFormatterParams is the structure any formatter will be handed when time to log comes
type LogFormatterParams struct {
	Request    *http.Request
	URL        url.URL
	TimeStamp  time.Time
	StatusCode int
	Size       int
}

// LogFormatter gives the signature of the formatter function passed to CustomLoggingHandler
type LogFormatter func(writer io.Writer, params LogFormatterParams)

// loggingHandler is the http.Handler implementation for LoggingHandlerTo and its
// friends

type loggingHandler struct {
	writer    io.Writer
	handler   http.Handler
	formatter LogFormatter
}

func (h loggingHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
}

func CustomLoggingHandler(out io.Writer, h http.Handler, f LogFormatter) http.Handler {
	return loggingHandler{out, h, f}
}

type stderrWriter struct{}

func (c *stderrWriter) Write(p []byte) (n int, err error) {
	return os.Stderr.Write(p)
}

type anonymousWrite func(p []byte) (n int, err error)

func TestDriver() string {
	custom := &stderrWriter{}
	testWriter(custom)
	//testWriter(custom.Write(nil))

	std := os.Stderr.Write
	testWriter(os.Stderr)

	if out, ok := any(std).(io.Writer); ok {
		testWriter(out)
		return "valid"
	}

	testWriter(testInterface())

	testWriter(func(w anonymousWrite) io.Writer { return os.Stderr }(nil))

	return "invalid"
}

func testWriter(out io.Writer) {

}

func testInterface() io.Writer {
	//type test struct{}

	return &stderrWriter{}
}
