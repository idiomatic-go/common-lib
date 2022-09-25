package util

import (
	"errors"
	"github.com/idiomatic-go/common-lib/httpxt"
	"net/http"
)

// HttpGeto - process a get request with error handling
func HttpGet(url string) *httpxt.ResponseStatus {
	req, err := http.NewRequest("", url, nil)
	if err != nil {
		LogPrintf("%v", err)
		return &httpxt.ResponseStatus{RequestErr: err}
	}
	status := httpxt.DoStatus(req)
	if !status.IsSuccess() && status.FirstError() != nil {
		LogPrintf("%v", status.FirstError())
	}
	return status
}

// HttpGetContent - processes a simple get request, unmarshalling content, and handling errors
func HttpGetContent(url string, content any) *httpxt.ResponseStatus {
	if content == nil {
		err0 := errors.New("invalid argument: content interface{} is nil")
		LogDebug("%v", err0)
		return &httpxt.ResponseStatus{RequestErr: err0}
	}
	status := HttpGet(url)
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
