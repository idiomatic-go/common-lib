package httpxt

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

func (r *ResponseStatus) IsError() bool {
	return r.HttpErr != nil || r.RequestErr != nil || r.UnmarshalErr != nil || r.BodyIOErr != nil
}

func (r *ResponseStatus) IsContent() bool {
	return r.Response != nil && r.Response.Body != nil && r.Response.StatusCode != http.StatusNoContent
}

func (r *ResponseStatus) IsSuccess() bool {
	return r.Response != nil && (r.Response.StatusCode >= http.StatusOK && r.Response.StatusCode < http.StatusMultipleChoices)
}

func (r *ResponseStatus) IsClientError() bool {
	return r.Response != nil && (r.Response.StatusCode >= http.StatusBadRequest && r.Response.StatusCode < http.StatusInternalServerError)
}

func (r *ResponseStatus) IsServerError() bool {
	return r.Response != nil && (r.Response.StatusCode >= http.StatusInternalServerError && r.Response.StatusCode <= 599)
}

func (r *ResponseStatus) StatusCode() int {
	if r.Response == nil {
		return 0
	}
	return r.Response.StatusCode
}

func (r *ResponseStatus) Url() *url.URL {
	if r.Response == nil || r.Response.Request == nil || r.Response.Request.URL == nil {
		return &url.URL{}
	}
	return r.Response.Request.URL
}

func (r *ResponseStatus) ReadBody() []byte {
	if !r.IsContent() {
		return nil
	}
	defer r.Response.Body.Close()
	var bytes []byte
	bytes, r.BodyIOErr = io.ReadAll(r.Response.Body)
	return bytes
}

func (r *ResponseStatus) UnmarshalJson(i interface{}) ([]byte, error) {
	if i == nil {
		return nil, errors.New("invalid argument: interface{} is nil")
	}
	body := r.ReadBody()
	if r.BodyIOErr != nil {
		return nil, nil
	}
	r.UnmarshalErr = json.Unmarshal(body, i)
	if r.UnmarshalErr != nil {
		return body, nil
	}
	return nil, nil
}

// DecodeJson is used when you want to read directly from a stream. This should be used when
// the requested content may be very large in size, as this uses less memory
func (r *ResponseStatus) DecodeJson(i interface{}) error {
	if i == nil {
		return errors.New("invalid argument: interface{} is nil")
	}
	if !r.IsContent() {
		return nil
	}
	r.UnmarshalErr = json.NewDecoder(r.Response.Body).Decode(i)
	return nil
}
