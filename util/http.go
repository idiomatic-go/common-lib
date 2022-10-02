package util

import (
	"context"
	"errors"
	"github.com/idiomatic-go/common-lib/httpxt"
	"io"
	"net/http"
)

// HttpDo - process a http request with error handling
func HttpDo(ctx context.Context, method, url string, header http.Header, body io.Reader) *httpxt.ResponseStatus {
	if ctx == nil {
		ctx = context.Background()
	}
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		LogPrintf("%v", err)
		return &httpxt.ResponseStatus{RequestErr: err}
	}
	httpxt.AddHeaders(req, header)
	status := httpxt.DoStatus(req)
	if !status.IsSuccess() && status.FirstError() != nil {
		LogPrintf("%v", status.FirstError())
	}
	return status
}

// HttpDoContent - process a http request with error handling
func HttpDoContent(ctx context.Context, method, url string, header http.Header, body io.Reader, content any) *httpxt.ResponseStatus {
	if content == nil {
		err0 := errors.New("invalid argument: content interface{} is nil")
		LogDebug("%v", err0)
		return &httpxt.ResponseStatus{RequestErr: err0}
	}
	status := HttpDo(ctx, method, url, header, body)
	if status.IsError() || !status.IsContent() {
		return status
	}
	entity, _ := status.UnmarshalJson(content)
	if status.FirstError() != nil {
		LogPrintf("%v", status.FirstError())
		if body != nil {
			LogDebug("%v", string(entity))
		}
	}
	return status
}

// HttpGet - process a get request with error handling
func HttpGet(ctx context.Context, url string, header http.Header) *httpxt.ResponseStatus {
	return HttpDo(ctx, "", url, header, nil)
}

// HttpGetContent - processes a get request, unmarshalling content, and handling errors
func HttpGetContent(ctx context.Context, url string, header http.Header, content any) *httpxt.ResponseStatus {
	return HttpDoContent(ctx, "", url, header, nil, content)
}

// HttpPost - process a post request with error handling
func HttpPost(ctx context.Context, url string, header http.Header, body io.Reader) *httpxt.ResponseStatus {
	return HttpDo(ctx, http.MethodPost, url, header, body)
}

// HttpPostContent - process a post request with error handling
func HttpPostContent(ctx context.Context, url string, header http.Header, body io.Reader, content any) *httpxt.ResponseStatus {
	return HttpDoContent(ctx, http.MethodPost, url, header, body, content)
}
