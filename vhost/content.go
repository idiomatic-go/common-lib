package vhost

import (
	"context"
	"errors"
	"fmt"
	"github.com/idiomatic-go/common-lib/eventing"
	"github.com/idiomatic-go/common-lib/fse"
	"reflect"
)

// CreateCredentialsMessage - functions
func CreateCredentialsMessage(to, from, event string, fn Credentials) eventing.Message {
	return eventing.CreateMessage(to, from, event, eventing.StatusNotProvided, fn)
}

func AccessCredentials(msg *eventing.Message) Credentials {
	if msg == nil || msg.Content == nil {
		return nil
	}
	fn, ok := msg.Content.(Credentials)
	if ok {
		return fn
	}
	return nil
}

func ProcessContent[T any](ctx context.Context) (T, Status) {
	var t T
	if ctx == nil {
		return t, NewStatusError(errors.New(fmt.Sprintf("vhost.ProcessContent internal error : context is nil")))
	}
	i := ctx.Value(contentKey)
	if IsNil(i) {
		return t, NewStatusError(errors.New(fmt.Sprintf("vhost.ProcessContent internal error : no content available")))
	}
	if buf, ok := i.(fse.Entry); ok {
		t1, err := fse.ProcessEntry[T](&buf)
		if err != nil {
			return t, NewStatusError(err)
		}
		return t1, NewStatusOk()
	}
	if status, ok := i.(Status); ok {
		return t, status
	}
	// Code for err must be after Status as Status is an error
	if err, ok := i.(error); ok {
		return t, NewStatusError(err)
	}
	if t1, ok := i.(T); ok {
		return t1, NewStatusOk()
	}
	return t, NewStatusError(errors.New(fmt.Sprintf("vhost.ProcessContent internal error : invalid content type : %v", reflect.TypeOf(i))))
}
