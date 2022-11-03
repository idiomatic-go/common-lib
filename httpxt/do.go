package httpxt

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"errors"
	"github.com/idiomatic-go/common-lib/fse"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/idiomatic-go/common-lib/httpxt/internal"
)

var client *http.Client

func init() {
	t, ok := http.DefaultTransport.(*http.Transport)
	if ok {
		// Used clone instead of assignment due to presence of sync.Mutex fields
		var transport = t.Clone()
		transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		transport.MaxIdleConns = 200
		transport.MaxIdleConnsPerHost = 100
		client = &http.Client{Transport: transport, Timeout: time.Second * 5}
	} else {
		client = &http.Client{Transport: http.DefaultTransport, Timeout: time.Second * 5}
	}
}

func DoStatus(req *http.Request) *ResponseStatus {
	status := ResponseStatus{}
	if req == nil {
		status.RequestErr = errors.New("invalid argument: Request is nil")
		return &status
	}
	status.Response, status.HttpErr = Do(req)
	return &status
}

func Do(req *http.Request) (resp *http.Response, err error) {
	if req == nil {
		return nil, errors.New("invalid argument: Request is nil")
	}
	var traceFinish HttpTraceFinish
	if traceStart != nil {
		traceFinish = traceStart(req)
	}
	if isEmbeddedContent(req) {
		resp, err = createEmbeddedResponse(fse.ContextContent(req.Context()), req)
	} else {
		switch req.URL.Scheme {
		case "echo":
			resp, err = createEchoResponse(req)
		default:
			resp, err = client.Do(req)
		}
	}
	if traceFinish != nil {
		traceFinish(resp, err)
	}
	return resp, err
}

func isEmbeddedContent(req *http.Request) bool {
	if req == nil || req.URL.RawQuery == "" {
		return false
	}
	return strings.Index(req.URL.RawQuery, EmbeddedContent) != -1
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

func createEmbeddedResponse(entry *fse.Entry, req *http.Request) (*http.Response, error) {
	if req == nil || entry == nil {
		return nil, errors.New("invalid argument: Request or Entry nil")
	}
	var path = entry.Name
	if isHttpResponseMessage(entry.Content) {
		return http.ReadResponse(bufio.NewReader(bytes.NewReader(entry.Content)), req)
	} else {
		resp := &http.Response{StatusCode: http.StatusOK, Header: make(http.Header), Body: &internal.ReaderCloser{Reader: bytes.NewReader(entry.Content), Err: nil}}
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
