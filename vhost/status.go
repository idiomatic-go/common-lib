package vhost

type status struct {
	errs  Errors
	errs2 Errors
	grpc  gRPCStatus
}

// String - fmt.Stringer interface implementation
func (s *status) String() string {
	str := s.Message()
	if str == "" {
		str = s.Error()
	}
	return str
}

// Ok - gRPC interface implementation
func (s *status) Ok() bool               { return s.grpc.Ok() }
func (s *status) InvalidArgument() bool  { return s.grpc.InvalidArgument() }
func (s *status) Unauthenticated() bool  { return s.grpc.Unauthenticated() }
func (s *status) PermissionDenied() bool { return s.grpc.PermissionDenied() }
func (s *status) NotFound() bool         { return s.grpc.NotFound() }
func (s *status) Internal() bool         { return s.grpc.Internal() }
func (s *status) Unavailable() bool      { return s.grpc.Unavailable() }
func (s *status) DeadlineExceeded() bool { return s.grpc.DeadlineExceeded() }
func (s *status) AlreadyExists() bool    { return s.grpc.AlreadyExists() }
func (s *status) Cancelled() bool        { return s.grpc.Cancelled() }

func (s *status) Code() int32     { return s.grpc.Code() }
func (s *status) Message() string { return s.grpc.Message() }

// Error - Errors interface implementation
func (s *status) Error() string   { return s.errs.Error() }
func (s *status) IsError() bool   { return s.errs.IsError() }
func (s *status) Errors() []error { return s.errs.Errors() }
func (s *status) Add(err error)   { s.errs.Add(err) }

// Handled - update to reflect that the errors have already been handled
func (s *status) Handled() Status {
	s.errs2 = s.errs
	return &status{errs: newErrors(), grpc: s.grpc}
}

func (s *status) HandledNewCode(code int32, msg string) Status {
	s.errs2 = s.errs
	return &status{errs: newErrors(), grpc: NewStatusCode(code, msg)}
}

// HttpStatus - convert gRPC -> Http
func (s *status) HttpStatus() int {
	code := 0
	switch s.grpc.Code() {
	case StatusOk:
		code = 200

	case StatusInvalidArgument:
		code = 400
	case StatusUnauthenticated:
		code = 401
	case StatusPermissionDenied:
		code = 403
	case StatusNotFound:
		code = 404

	case StatusInternal:
		code = 500
	case StatusUnavailable:
		code = 503
	case StatusDeadlineExceeded:
		code = 504
	case StatusInvalidContent,
		StatusNotProvided,
		StatusCancelled,
		StatusUnknown,
		StatusAlreadyExists,
		StatusResourceExhausted,
		StatusFailedPrecondition,
		StatusAborted,
		StatusOutOfRange,
		StatusUnimplemented,
		StatusDataLoss:
	}
	return code
}

func NewStatusOk() Status {
	s := status{errs: newErrors(), grpc: NewgRPCStatus(StatusOk, "")}
	return &s
}

func NewStatusCode(code int32, a any) Status {
	s := status{errs: NewErrorsAny(a), grpc: NewgRPCStatus(code, a)}
	return &s
}

func NewStatusInvalidArgument(a any) Status {
	return NewStatusCode(StatusInvalidArgument, a)
}

func NewStatusUnauthenticated(a any) Status {
	return NewStatusCode(StatusUnauthenticated, a)
}

func NewStatusPermissionDenied(a any) Status {
	return NewStatusCode(StatusPermissionDenied, a)
}

func NewStatusNotFound(a any) Status {
	return NewStatusCode(StatusNotFound, a)
}

func NewStatusInternal(a any) Status {
	return NewStatusCode(StatusInternal, a)
}

func NewStatusUnavailable(a any) Status {
	return NewStatusCode(StatusUnavailable, a)
}

func NewStatusDeadlineExceeded(a any) Status {
	return NewStatusCode(StatusDeadlineExceeded, a)
}

func NewStatusAlreadyExists(a any) Status {
	return NewStatusCode(StatusAlreadyExists, a)
}

func NewStatusCancelled(a any) Status {
	return NewStatusCode(StatusCancelled, a)
}

func NewStatusError(err ...error) Status {
	s := status{errs: NewErrorsList(err), grpc: nil}
	if len(s.errs.Errors()) > 0 {
		s.grpc = NewgRPCStatus(StatusInternal, "")
	} else {
		s.grpc = NewgRPCStatus(StatusOk, "")
	}
	return &s
}
