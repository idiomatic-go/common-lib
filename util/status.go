package util

type status struct {
	errs Errors
	grpc gRPCStatus
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
func (s *status) NotFound() bool         { return s.grpc.NotFound() }
func (s *status) DeadlineExceeded() bool { return s.grpc.DeadlineExceeded() }
func (s *status) AlreadyExists() bool    { return s.grpc.AlreadyExists() }
func (s *status) Code() int32            { return s.grpc.Code() }
func (s *status) Message() string        { return s.grpc.Message() }

// Error - Errors interface implementation
func (s *status) Error() string   { return s.errs.Error() }
func (s *status) IsError() bool   { return s.errs.IsError() }
func (s *status) Errors() []error { return s.errs.Errors() }
func (s *status) Add(err error)   { s.errs.Add(err) }
func (s *status) Cat() string     { return s.errs.Cat() }

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

func NewStatusNotFound(a any) Status {
	return NewStatusCode(StatusNotFound, a)
}

func NewStatusDeadlineExceeded(a any) Status {
	return NewStatusCode(StatusDeadlineExceeded, a)
}

func NewStatusAlreadyExists(a any) Status {
	return NewStatusCode(StatusAlreadyExists, a)
}

func NewStatusError(err ...error) Status {
	s := status{errs: NewErrorsList(err), grpc: NewgRPCStatus(StatusNotProvided, "")}
	return &s
}
