package httpxt

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"strconv"
	"strings"
	"time"

	internal "github.com/idiomatic-go/common-lib/httpxt/internal"
)

func init() {
	t, ok := http.DefaultTransport.(*http.Transport)
	if ok {
		// Used clone instead of assignment due to presence of sync.Mutex fields
		var transport = t.Clone()
		transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		transport.MaxIdleConns = 200
		transport.MaxIdleConnsPerHost = 100
		Client = &http.Client{Transport: transport, Timeout: time.Second * 5}
	} else {
		Client = &http.Client{Transport: http.DefaultTransport, Timeout: time.Second * 5}
	}
}

func DoWithStatus(req *http.Request) *ResponseStatus {
	status := ResponseStatus{}
	status.Response, status.HttpErr = Do(req)
	return &status
}

func Do(req *http.Request) (resp *http.Response, err error) {
	if req == nil {
		return nil, errors.New("invalid argument: Request is nil")
	}
	var traceFinish HttpTraceFinish
	if TraceStart != nil {
		traceFinish = TraceStart(req)
	}
	switch req.URL.Scheme {
	case "http", "https":
		resp, err = Client.Do(req)
	case "file":
		if fsys == nil {
			return nil, fmt.Errorf("no file system mounted")
		}
		resp, err = createFileResponse(req)
	case "echo":
		resp, err = createEchoResponse(req)
	default:
		resp, err = nil, fmt.Errorf("invalid argument: URL scheme is not supported [%v]", req.URL.Scheme)
	}
	if traceFinish != nil {
		traceFinish(resp, err)
	}
	return resp, err
}

var http11Bytes = []byte("HTTP/1.1")
var http12Bytes = []byte("HTTP/1.2")
var http20Bytes = []byte("HTTP/2.0")

func isHttpResponseMessage(buf []byte) bool {
	if buf == nil {
		return false
	}
	l := len(buf)
	if bytes.Equal(buf[0:l], http11Bytes) {
		return true
	}
	if bytes.Equal(buf[0:l], http12Bytes) {
		return true
	}
	if bytes.Equal(buf[0:l], http20Bytes) {
		return true
	}
	return false
}

func createFileResponse(req *http.Request) (*http.Response, error) {
	path := req.URL.Path
	path = strings.TrimPrefix(path, "/")
	buf, err := fs.ReadFile(fsys, path)
	if err != nil {
		if strings.Contains(err.Error(), "file does not exist") {
			return &http.Response{StatusCode: http.StatusNotFound}, nil
		}
		return &http.Response{StatusCode: http.StatusInternalServerError}, nil
	}
	if isHttpResponseMessage(buf) {
		return http.ReadResponse(bufio.NewReader(bytes.NewReader(buf)), req)
	} else {
		resp := &http.Response{StatusCode: http.StatusOK, Header: make(http.Header), Body: &internal.ReaderCloser{Reader: bytes.NewReader(buf), Err: nil}}
		if strings.HasSuffix(path, ".json") {
			resp.Header.Add("Content-Type", "application/json")
		} else {
			resp.Header.Add("Content-Type", "text/plain")
		}
		return resp, nil
	}
}

func createEchoResponse(req *http.Request) (*http.Response, error) {
	if req == nil {
		return nil, errors.New("invalid argument: Request is nil")
	}
	var resp = http.Response{StatusCode: http.StatusOK, Header: make(http.Header), Request: req}
	for key, element := range req.URL.Query() {
		switch key {
		case "httpError":
			return nil, http.ErrHijacked
		case "status":
			sc, err := strconv.Atoi(element[0])
			if err == nil {
				resp.StatusCode = sc
			} else {
				resp.StatusCode = http.StatusInternalServerError
			}
		case "body":
			if len(element[0]) > 0 && resp.Body == nil {
				// Handle escaped path? See notes on the url.URL struct
				resp.Body = &internal.ReaderCloser{Reader: strings.NewReader(element[0]), Err: nil}
			}
		case "ioError":
			// resp.StatusCode = http.StatusInternalServerError
			resp.Body = &internal.ReaderCloser{Reader: nil, Err: io.ErrUnexpectedEOF}
		default:
			if len(element[0]) > 0 {
				resp.Header.Add(key, element[0])
			}
		}
	}
	return &resp, nil
}
