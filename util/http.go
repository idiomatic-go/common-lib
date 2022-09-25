package util

import (
	"context"
	"errors"
	"github.com/idiomatic-go/common-lib/httpxt"
	"net/http"
)

// HttpGet - process a get request with error handling
func HttpGet(ctx context.Context, url string, header http.Header) *httpxt.ResponseStatus {
	if ctx == nil {
		ctx = context.Background()
	}
	req, err := http.NewRequestWithContext(ctx, "", url, nil)
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

// HttpGetContent - processes a get request, unmarshalling content, and handling errors
func HttpGetContent(ctx context.Context, url string, header http.Header, content any) *httpxt.ResponseStatus {
	if content == nil {
		err0 := errors.New("invalid argument: content interface{} is nil")
		LogDebug("%v", err0)
		return &httpxt.ResponseStatus{RequestErr: err0}
	}
	status := HttpGet(ctx, url, header)
	if status.IsError() {
		return status
	}
	body, _ := status.UnmarshalJson(content)
	if status.FirstError() != nil {
		LogPrintf("%v", status.FirstError())
		if body != nil {
			LogDebug("%v", string(body))
		}
	}
	return status
}
