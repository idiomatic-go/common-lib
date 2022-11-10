package util

import "reflect"

func (r *Response) IsContentNil() bool {
	return IsNil(r.Content)
}

func (r *Response) IsContentSerialized() bool {
	if r.IsContentNil() {
		return false
	}
	if _, ok := r.Content.([]byte); ok {
		return ok
	}
	return false
}

func (r *Response) ContentBytes() (buf []byte, ok bool) {
	if !r.IsContentSerialized() {
		return nil, false
	}
	if buf, ok = r.Content.([]byte); ok {
		return buf, ok
	}
	return nil, false
}

func NewResponse[T any](status StatusCode, content T) *Response {
	sc := status
	if status == nil {
		sc = NewStatusOk()
	}
	i := any(content)
	if i == nil {
		return &Response{Status: sc, Content: content}
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.String:
		if s, ok := i.(string); ok {
			buf := []byte(s)
			return &Response{Status: sc, Content: buf}
		}
	default:
	}
	return &Response{Status: sc, Content: content}
}

func NewResponseHeaders[T any](status StatusCode, content T, headers any) *Response {
	resp := NewResponse(status, content)
	resp.Headers = headers
	return resp
}
