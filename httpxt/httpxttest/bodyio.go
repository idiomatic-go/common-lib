package httpxttest

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
)

func ReadBody(r interface{}) (body []byte, err error) {
	if r == nil {
		return nil, nil
	}
	var reader io.ReadCloser
	if req, ok := r.(*http.Request); ok {
		reader = req.Body
	} else {
		if resp, ok := r.(*http.Response); ok {
			reader = resp.Body
		} else {
			if recorder, ok := r.(*httptest.ResponseRecorder); ok {
				return recorder.Body.Bytes(), nil
			}
		}
	}
	if reader == nil {
		return nil, nil
	}
	defer reader.Close()
	return io.ReadAll(reader)
}

func ReadBodyString(i interface{}) string {
	body, err := ReadBody(i)
	if body == nil {
		return ""
	}
	if err != nil {
		return err.Error()
	}
	return strings.TrimSpace(string(body))
}
