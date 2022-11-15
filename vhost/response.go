package vhost

import "reflect"

type response struct {
	status  Status
	content any
	headers any
}

func (r *response) IsContentNil() bool {
	return IsNil(r.content)
}

func (r *response) IsContentSerialized() bool {
	if r.IsContentNil() {
		return false
	}
	if _, ok := r.content.([]byte); ok {
		return ok
	}
	return false
}

func (r *response) ContentBytes() (buf []byte, ok bool) {
	if !r.IsContentSerialized() {
		return nil, false
	}
	if buf, ok = r.content.([]byte); ok {
		return buf, ok
	}
	return nil, false
}

func (r *response) Content() any {
	return r.content
}

func (r *response) Headers() any {
	return r.headers
}

func (r *response) Status() Status {
	return r.status
}

func NewResponse(status Status) Response {
	sc := status
	if status == nil {
		sc = NewStatusOk()
	}
	return &response{status: sc, content: nil, headers: nil}
}

func NewResponseContent[T any](status Status, content T) Response {
	sc := status
	if status == nil {
		sc = NewStatusOk()
	}
	i := any(content)
	if i == nil {
		return &response{status: sc, content: content}
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.String:
		if s, ok := i.(string); ok {
			buf := []byte(s)
			return &response{status: sc, content: buf}
		}
	default:
	}
	return &response{status: sc, content: content}
}

func NewResponseHeaders[T any](status Status, headers any) Response {
	return &response{status: status, content: nil, headers: headers}
}
