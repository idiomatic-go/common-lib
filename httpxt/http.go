package httpxt

import (
	"context"
	"errors"
	"github.com/idiomatic-go/common-lib/logxt"
	"io"
	"net/http"
)

// HttpDo - process a http request with error handling
func HttpDo(ctx context.Context, method, url string, header http.Header, body io.Reader) *ResponseStatus {
	if ctx == nil {
		ctx = context.Background()
	}
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		logxt.Debug(err)
		return &ResponseStatus{RequestErr: err}
	}
	AddHeaders(req, header)
	status := DoStatus(req)
	if !status.IsSuccess() && status.FirstError() != nil {
		logxt.Printf("%v", status.FirstError())
	}
	return status
}

// HttpDoContent - process a http request with error handling
func HttpDoContent(ctx context.Context, method, url string, header http.Header, body io.Reader, content any) *ResponseStatus {
	if content == nil {
		err0 := errors.New("invalid argument: content interface{} is nil")
		logxt.Debug(err0)
		return &ResponseStatus{RequestErr: err0}
	}
	status := HttpDo(ctx, method, url, header, body)
	if status.IsError() || !status.IsContent() {
		return status
	}
	entity, _ := status.UnmarshalJson(content)
	if status.FirstError() != nil {
		logxt.Printf("%v", status.FirstError())
		if body != nil {
			logxt.Debug(string(entity))
		}
	}
	return status
}

// HttpGet - process a get request with error handling
func HttpGet(ctx context.Context, url string, header http.Header) *ResponseStatus {
	return HttpDo(ctx, "", url, header, nil)
}

// HttpGetContent - processes a get request, unmarshalling content, and handling errors
func HttpGetContent(ctx context.Context, url string, header http.Header, content any) *ResponseStatus {
	return HttpDoContent(ctx, "", url, header, nil, content)
}

// HttpPost - process a post request with error handling
func HttpPost(ctx context.Context, url string, header http.Header, body io.Reader) *ResponseStatus {
	return HttpDo(ctx, http.MethodPost, url, header, body)
}

// HttpPostContent - process a post request with error handling
func HttpPostContent(ctx context.Context, url string, header http.Header, body io.Reader, content any) *ResponseStatus {
	return HttpDoContent(ctx, http.MethodPost, url, header, body, content)
}
