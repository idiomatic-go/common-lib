package httpxt

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

// Response status
type ResponseStatus struct {
	HttpErr   error
	Response  *http.Response
	RequestID string
}

func (c *ResponseStatus) IsError() bool {
	return c.HttpErr != nil
}

func (c *ResponseStatus) IsContent() bool {
	return c.Response != nil && c.Response.Body != nil && c.Response.StatusCode != http.StatusNoContent
}

func (c *ResponseStatus) IsSuccess() bool {
	return c.Response != nil && (c.Response.StatusCode >= http.StatusOK && c.Response.StatusCode < http.StatusMultipleChoices)
}

func (c *ResponseStatus) IsClientError() bool {
	return c.Response != nil && (c.Response.StatusCode >= http.StatusBadRequest && c.Response.StatusCode < http.StatusInternalServerError)
}

func (c *ResponseStatus) IsServerError() bool {
	return c.Response != nil && (c.Response.StatusCode >= http.StatusInternalServerError && c.Response.StatusCode <= 599)
}

func (c *ResponseStatus) StatusCode() int {
	if c.Response == nil {
		return 0
	}
	return c.Response.StatusCode
}

func (c *ResponseStatus) Url() *url.URL {
	if c.Response == nil || c.Response.Request == nil || c.Response.Request.URL == nil {
		return &url.URL{}
	}
	return c.Response.Request.URL
}

func (c *ResponseStatus) ReadBody() ([]byte, error) {
	if !c.IsContent() {
		return nil, errors.New("invalid argument: interface{} is nil")
	}
	defer c.Response.Body.Close()
	return io.ReadAll(c.Response.Body)
}

func (c *ResponseStatus) UnmarshalJson(i interface{}) (body []byte, err error) {
	if i == nil {
		return nil, errors.New("invalid argument: interface{} is nil")
	}
	body, err = c.ReadBody()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, i)
	if err != nil {
		return body, err
	}
	return nil, nil
}

// DecodeJson is used when you want to read directly from a stream. This should be used when
// the requested content may be very large in size, as this uses less memory
func (c *ResponseStatus) DecodeJson(i interface{}) error {
	if i == nil {
		return errors.New("invalid argument: interface{} is nil")
	}
	if !c.IsContent() {
		return errors.New("invalid argument: no content is available")
	}
	return json.NewDecoder(c.Response.Body).Decode(i)
}
