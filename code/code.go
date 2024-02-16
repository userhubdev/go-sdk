package code

import "net/http"

type Code string

const (
	OK                 Code = "OK"
	Canceled           Code = "CANCELLED"
	Unknown            Code = "UNKNOWN"
	InvalidArgument    Code = "INVALID_ARGUMENT"
	DeadlineExceeded   Code = "DEADLINE_EXCEEDED"
	NotFound           Code = "NOT_FOUND"
	AlreadyExists      Code = "ALREADY_EXISTS"
	PermissionDenied   Code = "PERMISSION_DENIED"
	Unauthenticated    Code = "UNAUTHENTICATED"
	ResourceExhausted  Code = "RESOURCE_EXHAUSTED"
	FailedPrecondition Code = "FAILED_PRECONDITION"
	Aborted            Code = "ABORTED"
	OutOfRange         Code = "OUT_OF_RANGE"
	Unimplemented      Code = "UNIMPLEMENTED"
	Internal           Code = "INTERNAL"
	Unavailable        Code = "UNAVAILABLE"
	DataLoss           Code = "DATA_LOSS"
)

func (c Code) String() string {
	return string(c)
}

func (c Code) Valid() bool {
	switch c {
	case OK,
		Canceled,
		Unknown,
		InvalidArgument,
		DeadlineExceeded,
		NotFound,
		AlreadyExists,
		PermissionDenied,
		Unauthenticated,
		ResourceExhausted,
		FailedPrecondition,
		Aborted,
		OutOfRange,
		Unimplemented,
		Internal,
		Unavailable,
		DataLoss:
		return true
	default:
		return false
	}
}

func (c Code) WebhookStatusCode() int {
	switch c {
	case OK:
		return http.StatusOK
	case AlreadyExists, FailedPrecondition, InvalidArgument:
		return http.StatusBadRequest
	case NotFound:
		return http.StatusNotFound
	case ResourceExhausted:
		return http.StatusTooManyRequests
	case Unimplemented:
		return http.StatusNotImplemented
	default:
		return http.StatusInternalServerError
	}
}
