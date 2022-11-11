package util

import (
	"github.com/idiomatic-go/common-lib/logxt"
	"strconv"
)

type statusCode struct {
	code  int32
	errs  map[string]error
	msg   string
	index int
	first string
}

func createStatusCode(code int32) *statusCode {
	return &statusCode{code: code, errs: make(map[string]error, 1)}
}

func createStatusCodeError(code int32, name string, err error) *statusCode {
	sc := createStatusCode(code)
	addError(sc, name, err)
	return sc
}

func addError(sc *statusCode, name string, err error) {
	if sc == nil || err == nil {
		return
	}
	if name == "" {
		name = strconv.Itoa(sc.index)
		sc.index++
	}
	if sc.first == "" {
		sc.first = name
	}
	sc.errs[name] = err
}

func (sc *statusCode) AddError(name string, err error) {
	addError(sc, name, err)
}

func (sc *statusCode) Ok() bool {
	return sc.code == StatusOk
}

func (sc *statusCode) InvalidArgument() bool {
	return sc.code == StatusInvalidArgument
}

func (sc *statusCode) NotFound() bool {
	return sc.code == StatusNotFound
}

func (sc *statusCode) DeadlineExceeded() bool {
	return sc.code == StatusDeadlineExceeded
}

func (sc *statusCode) AlreadyExists() bool {
	return sc.code == StatusAlreadyExists
}

func (sc *statusCode) IsError() bool {
	return len(sc.errs) > 0
}

func (sc *statusCode) Errors() map[string]error {
	return sc.errs
}

func (sc *statusCode) Message() string {
	if sc.IsError() {
		return sc.Error()
	}
	return sc.msg
}

func (sc *statusCode) Code() int32 {
	return sc.code
}

func (sc *statusCode) Error() string {
	if sc.IsError() {
		return sc.errs[sc.first].Error()
	}
	return ""
}

func (sc *statusCode) String() string {
	if sc.IsError() {
		return sc.Error()
	}
	return sc.Message()
}

func NewStatusOk() StatusCode {
	return createStatusCode(StatusOk)
}

func NewStatusOptionalNotFound(isNull bool, msg string) StatusCode {
	if isNull {
		return NewStatusNotFound(msg)
	}
	return NewStatusOk()
}

func NewStatusNotFound(msg string) StatusCode {
	sc := createStatusCode(StatusNotFound)
	sc.msg = msg
	return sc
}

func NewStatusError(errs ...error) StatusCode {
	var sc = createStatusCode(StatusNotProvided)
	for _, e := range errs {
		if e == nil {
			continue
		}
		addError(sc, "", e)
	}
	if len(sc.errs) == 0 {
		sc.code = StatusOk
	}
	return sc
}

func NewStatusErrorAttrs(errs ...Attr) StatusCode {
	var sc = createStatusCode(StatusNotProvided)
	for _, attr := range errs {
		if attr.Val == nil {
			continue
		}
		if err, ok := attr.Val.(error); ok {
			addError(sc, attr.Name, err)
		} else {
			logxt.LogDebug("invalid attribute Val %v", attr.Val)
		}
	}
	if !sc.IsError() {
		sc.code = StatusOk
	}
	return sc
}

func NewStatusInvalidArgument(err error) StatusCode {
	return createStatusCodeError(StatusInvalidArgument, "", err)
}

func NewStatusDeadlineExceeded(err error) StatusCode {
	return createStatusCodeError(StatusDeadlineExceeded, "", err)
}

func NewStatusAlreadyExists(err error) StatusCode {
	return createStatusCodeError(StatusAlreadyExists, "", err)
}
