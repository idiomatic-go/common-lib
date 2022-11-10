package util

type statusCode struct {
	code int32
	errs []error
	msg  string
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
	return sc.errs != nil
}

func (sc *statusCode) Errors() []error {
	return sc.errs
}

func (sc *statusCode) Message() string {
	if sc.IsError() {
		return sc.errs[0].Error()
	}
	return sc.msg
}

func (sc *statusCode) Code() int32 {
	return sc.code
}

func (sc *statusCode) Error() string {
	if sc.IsError() {
		return sc.errs[0].Error()
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
	return &statusCode{code: StatusOk}
}

func NewStatusOptionalNotFound(isNull bool, msg string) StatusCode {
	if isNull {
		return NewStatusNotFound(msg)
	}
	return &statusCode{code: StatusOk}
}

func NewStatusNotFound(msg string) StatusCode {
	return &statusCode{code: StatusNotFound, msg: msg}
}

func NewStatusError(err ...error) StatusCode {
	//if err == nil || (len(err) == 1 && err[0] == nil) {
	//	return NewStatusOk()
	//}
	var sc = statusCode{code: StatusNotProvided}
	for _, e := range err {
		if e != nil {
			sc.errs = append(sc.errs, e)
		}
	}
	if len(sc.errs) == 0 {
		sc.code = StatusOk
	}
	return &sc
}

func NewStatusInvalidArgument(err error) StatusCode {
	if err == nil {
		return &statusCode{code: StatusInvalidArgument, errs: nil}
	}
	return &statusCode{code: StatusInvalidArgument, errs: []error{err}}
}

func NewStatusDeadlineExceeded(err error) StatusCode {
	if err == nil {
		return &statusCode{code: StatusDeadlineExceeded, errs: nil}
	}
	return &statusCode{code: StatusDeadlineExceeded, errs: []error{err}}
}

func NewStatusAlreadyExists(err error) StatusCode {
	if err == nil {
		return &statusCode{code: StatusAlreadyExists, errs: nil}
	}
	return &statusCode{code: StatusAlreadyExists, errs: []error{err}}
}
