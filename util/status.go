package util

type status struct {
	errs Errors
	grpc gRPCStatus
}

// String - fmt.Stringer interface implementation
func (s *status) String() string {
	str := s.Message()
	if str == "" {
		str = "" //s.Error()
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
func (s *status) Errors() []error { return s.errs.Errors() }
func (s *status) Add(err error)   { s.errs.Add(err) }
func (s *status) Cat() string     { return s.errs.Cat() }

func NewStatusOk_GRPC() Status {
	s := status{errs: newErrors(), grpc: NewgRPCStatus(StatusOk, "")}
	return &s
}

func NewStatusInvalidArgument_GRPC(msg string) Status {
	s := status{errs: newErrors(), grpc: NewgRPCStatus(StatusInvalidArgument, msg)}
	return &s
}

func NewStatusNotFound_GRPC(msg string) Status {
	s := status{errs: newErrors(), grpc: NewgRPCStatus(StatusNotFound, msg)}
	return &s
}

func NewStatusDeadlineExceeded_GRPC(msg string) Status {
	s := status{errs: newErrors(), grpc: NewgRPCStatus(StatusDeadlineExceeded, msg)}
	return &s
}

func NewStatusAlreadyExists_GRPC(msg string) Status {
	s := status{errs: newErrors(), grpc: NewgRPCStatus(StatusAlreadyExists, msg)}
	return &s
}

func NewStatusErrors(err ...error) Status {
	s := status{errs: newErrors(), grpc: newgRPCStatus()}
	return &s
}
