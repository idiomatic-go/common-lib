package fncall

const (
	// gRPC status codes
	// https://grpc.github.io/grpc/core/md_doc_statuscodes.html
	StatusInvalidContent     = int32(-1) // Content is not available or is of the wrong type, usually found via unmarshalling
	StatusOk                 = int32(0)  // Not an error; returned on success.
	StatusCancelled          = int32(1)  // The operation was cancelled, typically by the caller.
	StatusUnknown            = int32(2)  // Unknown error. For example, this error may be returned when a Status value received from another address space belongs to an error space that is not known in this address space. Also errors raised by APIs that do not return enough error information may be converted to this error.
	StatusInvalidArgument    = int32(3)  // The client specified an invalid argument. Note that this differs from FAILED_PRECONDITION. INVALID_ARGUMENT indicates arguments that are problematic regardless of the state of the system (e.g., a malformed file name).
	StatusDeadlineExceeded   = int32(4)  // The deadline expired before the operation could complete. For operations that change the state of the system, this error may be returned even if the operation has completed successfully. For example, a successful response from a server could have been delayed long
	StatusNotFound           = int32(5)  // Some requested entity (e.g., file or directory) was not found. Note to server developers: if a request is denied for an entire class of users, such as gradual feature rollout or undocumented allowlist, NOT_FOUND may be used. If a request is denied for some users within a class of users, such as user-based access control, PERMISSION_DENIED must be used.
	StatusAlreadyExists      = int32(6)  // The entity that a client attempted to create (e.g., file or directory) already exists.
	StatusPermissionDenied   = int32(7)  // The caller does not have permission to execute the specified operation. PERMISSION_DENIED must not be used for rejections caused by exhausting some resource (use RESOURCE_EXHAUSTED instead for those errors). PERMISSION_DENIED must not be used if the caller can not be identified (use UNAUTHENTICATED instead for those errors). This error code does not imply the request is valid or the requested entity exists or satisfies other pre-conditions.
	StatusResourceExhausted  = int32(8)  // Some resource has been exhausted, perhaps a per-user quota, or perhaps the entire file system is out of space.
	StatusFailedPrecondition = int32(9)  // The operation was rejected because the system is not in a state required for the operation's execution. For example, the directory to be deleted is non-empty, an rmdir operation is applied to a non-directory, etc. Service implementors can use the following guidelines to decide between FAILED_PRECONDITION, ABORTED, and UNAVAILABLE: (a) Use UNAVAILABLE if the client can retry just the failing call. (b) Use ABORTED if the client should retry at a higher level (e.g., when a client-specified test-and-set fails, indicating the client should restart a read-modify-write sequence). (c) Use FAILED_PRECONDITION if the client should not retry until the system state has been explicitly fixed. E.g., if an "rmdir" fails because the directory is non-empty, FAILED_PRECONDITION should be returned since the client should not retry unless the files are deleted from the directory.
	StatusAborted            = int32(10) // The operation was aborted, typically due to a concurrency issue such as a sequencer check failure or transaction abort. See the guidelines above for deciding between FAILED_PRECONDITION, ABORTED, and UNAVAILABLE.
	StatusOutOfRange         = int32(11) // The operation was attempted past the valid range. E.g., seeking or reading past end-of-file. Unlike INVALID_ARGUMENT, this error indicates a problem that may be fixed if the system state changes. For example, a 32-bit file system will generate INVALID_ARGUMENT if asked to read at an offset that is not in the range [0,2^32-1], but it will generate OUT_OF_RANGE if asked to read from an offset past the current file size. There is a fair bit of overlap between FAILED_PRECONDITION and OUT_OF_RANGE. We recommend using OUT_OF_RANGE (the more specific error) when it applies so that callers who are iterating through a space can easily look for an OUT_OF_RANGE error to detect when they are done.
	StatusUnimplemented      = int32(12) // The operation is not implemented or is not supported/enabled in this service.
	StatusInternal           = int32(13) // Internal errors. This means that some invariants expected by the underlying system have been broken. This error code is reserved for serious errors.
	StatusUnavailable        = int32(14) // The service is currently unavailable. This is most likely a transient condition, which can be corrected by retrying with a backoff. Note that it is not always safe to retry non-idempotent operations.
	StatusDataLoss           = int32(15) // Unrecoverable data loss or corruption.
	StatusUnauthenticated    = int32(16) // The request does not have valid authentication credentials for the operation.
)

type gRPCStatus interface {
	Ok() bool
	InvalidArgument() bool
	Unauthenticated() bool
	PermissionDenied() bool
	NotFound() bool

	Internal() bool
	Unavailable() bool
	DeadlineExceeded() bool

	Cancelled() bool
	AlreadyExists() bool

	Code() int32
	IsMessage() bool
	Message() string
}

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

func (s *grpcStatus) Unauthenticated() bool {
	return s.code == StatusUnauthenticated
}

func (s *grpcStatus) PermissionDenied() bool {
	return s.code == StatusPermissionDenied
}

func (s *grpcStatus) NotFound() bool {
	return s.code == StatusNotFound
}

func (s *grpcStatus) Internal() bool {
	return s.code == StatusInternal
}

func (s *grpcStatus) Unavailable() bool {
	return s.code == StatusInternal
}

func (s *grpcStatus) DeadlineExceeded() bool {
	return s.code == StatusDeadlineExceeded
}

func (s *grpcStatus) AlreadyExists() bool {
	return s.code == StatusAlreadyExists
}

func (s *grpcStatus) Cancelled() bool {
	return s.code == StatusCancelled
}

func (s *grpcStatus) IsMessage() bool {
	return s.msg != ""
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