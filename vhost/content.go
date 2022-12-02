package vhost

import (
	"context"
	"errors"
	"fmt"
	"github.com/idiomatic-go/common-lib/eventing"
	"github.com/idiomatic-go/common-lib/fse"
	"github.com/idiomatic-go/common-lib/logxt"
	"github.com/idiomatic-go/common-lib/util"
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

func ProcessContent[T any](content any) (T, Status) {
	var t T
	if util.IsNil(content) {
		status := NewStatusInvalidArgument(errors.New("vhost.ProcessContent internal error : no content available"))
		logxt.Debug(status)
		return t, status
	}
	if buf, ok := content.(fse.Entry); ok {
		t1, err := fse.ProcessEntry[T](&buf)
		if err != nil {
			return t, NewStatusError(err)
		}
		return t1, NewStatusOk()
	}
	if status, ok := content.(Status); ok {
		return t, status
	}
	// Code for err must be after Status as Status is an error
	if err, ok := content.(error); ok {
		return t, NewStatusError(err)
	}
	if grpc, ok := content.(gRPCStatus); ok {
		return t, NewStatusGRPC(grpc)
	}
	if t1, ok := content.(T); ok {
		return t1, NewStatusOk()
	}
	// TODO : update to reflect contained type.
	status := NewStatusInvalidArgument(errors.New(fmt.Sprintf("vhost.ProcessContent internal error : invalid content type : %v", reflect.TypeOf(content))))
	logxt.Debug(status)
	return t, status
}

func ProcessContextContent[T any](ctx context.Context) (T, Status) {
	var t T
	if ctx == nil {
		status := NewStatusInvalidArgument(errors.New(fmt.Sprintf("vhost.ProcessContent internal error : context is nil")))
		logxt.Debug(status)
		return t, status
	}
	i := ctx.Value(contentKey)
	if util.IsNil(i) {
		status := NewStatusInvalidArgument(errors.New(fmt.Sprintf("vhost.ProcessContent internal error : no content available")))
		logxt.Debug(status)
		return t, status
	}
	return ProcessContent[T](i)
}
