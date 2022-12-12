package fncall

import "fmt"

type Status interface {
	fmt.Stringer
	gRPCStatus
	Errors
	HttpStatus() int
	Handled() Status
}

type statusT struct {
	errs  Errors
	errs2 Errors
	grpc  gRPCStatus
}

// String - fmt.Stringer interface implementation
func (s *statusT) String() string {
	if s.errs2 != nil {
		return s.errs2.Error()
	}
	if s.IsError() {
		return s.Error()
	}
	return s.Message()
}

// Ok - gRPC interface implementation
func (s *statusT) Ok() bool               { return s.grpc.Ok() }
func (s *statusT) InvalidArgument() bool  { return s.grpc.InvalidArgument() }
func (s *statusT) Unauthenticated() bool  { return s.grpc.Unauthenticated() }
func (s *statusT) PermissionDenied() bool { return s.grpc.PermissionDenied() }
func (s *statusT) NotFound() bool         { return s.grpc.NotFound() }
func (s *statusT) Internal() bool         { return s.grpc.Internal() }
func (s *statusT) Unavailable() bool      { return s.grpc.Unavailable() }
func (s *statusT) DeadlineExceeded() bool { return s.grpc.DeadlineExceeded() }
func (s *statusT) AlreadyExists() bool    { return s.grpc.AlreadyExists() }
func (s *statusT) Cancelled() bool        { return s.grpc.Cancelled() }

func (s *statusT) Code() int32     { return s.grpc.Code() }
func (s *statusT) IsMessage() bool { return s.grpc.Message() != "" }
func (s *statusT) Message() string { return s.grpc.Message() }

// Error - Errors interface implementation
func (s *statusT) Error() string   { return s.errs.Error() }
func (s *statusT) IsError() bool   { return s.errs.IsError() }
func (s *statusT) Errors() []error { return s.errs.Errors() }
func (s *statusT) Add(err error)   { s.errs.Add(err) }

// Handled - update to reflect that the errors have already been handled
func (s *statusT) Handled() Status {
	s.errs2 = s.errs
	return &statusT{errs: newErrors(), grpc: s.grpc}
}

/*
func (s *statusT) HandledCode(code int32) Status {
	s.errs2 = s.errs
	return &statusT{errs: newErrors(), grpc: NewStatusCode(code,"")}
}

*/

// HttpStatus - convert gRPC -> Http
func (s *statusT) HttpStatus() int {
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
	s := statusT{errs: newErrors(), grpc: NewgRPCStatus(StatusOk, "")}
	return &s
}

func NewStatusOkMessage(a any) Status {
	s := statusT{errs: newErrors(), grpc: NewgRPCStatus(StatusOk, a)}
	return &s
}

func NewStatusCode(code int32, a any) Status {
	s := statusT{errs: NewErrorsAny(a), grpc: NewgRPCStatus(code, a)}
	return &s
}

func NewStatusGRPC(grpcStatus gRPCStatus) Status {
	s := statusT{errs: nil, grpc: grpcStatus}
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
	s := statusT{errs: NewErrorsList(err), grpc: nil}
	if len(s.errs.Errors()) > 0 {
		s.grpc = NewgRPCStatus(StatusInternal, "")
	} else {
		s.grpc = NewgRPCStatus(StatusOk, "")
	}
	return &s
}
