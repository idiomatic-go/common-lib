package util

type grpcStatus struct {
	code int32
	msg  string
}

func creategRPCStatus(code int32, msg string) gRPCStatus {
	return &grpcStatus{code: code, msg: msg}
}

func (s *grpcStatus) Ok() bool {
	return s.code == StatusOk
}

func (s *grpcStatus) InvalidArgument() bool {
	return s.code == StatusInvalidArgument
}

func (s *grpcStatus) NotFound() bool {
	return s.code == StatusNotFound
}

func (s *grpcStatus) DeadlineExceeded() bool {
	return s.code == StatusDeadlineExceeded
}

func (s *grpcStatus) AlreadyExists() bool {
	return s.code == StatusAlreadyExists
}

func (s *grpcStatus) Internal() bool {
	return s.code == StatusInternal
}

func (s *grpcStatus) Message() string {
	return s.msg
}

func (s *grpcStatus) Code() int32 {
	return s.code
}

func NewgRPCStatus(code int32, msg any) gRPCStatus {
	if IsNil(msg) {
		return creategRPCStatus(code, "")
	}
	if msg, ok := msg.(string); ok {
		return creategRPCStatus(code, msg)
	}
	return creategRPCStatus(code, "")
}

func newgRPCStatus() gRPCStatus {
	return creategRPCStatus(0, "")
}
