package types

import (
	"strconv"
	"strings"

	"github.com/userhubdev/go-sdk/apiv1"
)

// UserHubError is the UserHub API error type.
type UserHubError struct {
	apiCode    string
	message    string
	reason     string
	param      string
	metadata   map[string]string
	call       string
	statusCode int
	cause      error
}

func NewUserHubError(
	call string,
	message string,
	status *apiv1.Status,
	statusCode int,
	cause error,
) error {
	e := &UserHubError{
		message:    message,
		call:       call,
		cause:      cause,
		statusCode: statusCode,
	}

	if status != nil {
		if status.Message != "" {
			e.message = status.Message
		}

		e.apiCode = status.Code
		e.reason = status.Reason
		e.param = status.Param
		e.metadata = status.Metadata
	}

	return e
}

// Error returns the error description.
func (e *UserHubError) Error() string {
	var b strings.Builder

	b.WriteString("userhub: ")

	if e == nil {
		b.WriteString("<nil error>")
		return b.String()
	}

	b.WriteString(e.Message())

	var p strings.Builder

	if e.call != "" {
		p.WriteString(", call: ")
		p.WriteString(e.call)
	}

	hasApiCode := e.apiCode != "" && e.apiCode != "UNKNOWN"

	if hasApiCode {
		p.WriteString(", apiCode: ")
		p.WriteString(e.apiCode)
	}

	if e.reason != "" {
		p.WriteString(", reason: ")
		p.WriteString(e.reason)
	}

	if e.param != "" {
		p.WriteString(", param: ")
		p.WriteString(e.param)
	}

	if !hasApiCode && e.statusCode > 0 {
		p.WriteString(", statusCode: ")
		p.WriteString(strconv.Itoa(e.statusCode))
	}

	parts := p.String()
	if len(parts) > 0 {
		b.WriteString(" (")
		b.WriteString(parts[2:])
		b.WriteString(")")
	}

	return b.String()
}

// ApiCode returns general error code (e.g. `INVALID_ARGUMENT`).
func (e *UserHubError) ApiCode() string {
	if e == nil || e.apiCode == "" {
		return "UNKNOWN"
	}
	return e.apiCode
}

// Message returns the developer-facing error message.
func (e *UserHubError) Message() string {
	if e == nil || e.message == "" {
		return "<no message>"
	}
	return e.message
}

// Reason returns the reason code for the error (e.g. `USER_PENDING_DELETION`).
func (e *UserHubError) Reason() string {
	if e == nil {
		return ""
	}
	return e.reason
}

// Param returns the parameter path related to the error (e.g. `member.userId`).
func (e *UserHubError) Param() string {
	if e == nil {
		return ""
	}
	return e.param
}

// Metadata returns the additional error related data.
func (e *UserHubError) Metadata() map[string]string {
	m := map[string]string{}

	if e != nil && e.metadata != nil {
		for key, value := range e.metadata {
			m[key] = value
		}
	}

	return m
}

// Call returns the method call name for the error (e.g. `admin.users.get`).
func (e *UserHubError) Call() string {
	if e == nil {
		return ""
	}
	return e.call
}

// StatusCode returns the HTTP status code (e.g. 500).
func (e *UserHubError) StatusCode() int {
	if e == nil {
		return 0
	}
	return e.statusCode
}

// Unwrap returns the underlying error.
func (e *UserHubError) Unwrap() error {
	if e != nil {
		return e.cause
	}
	return nil
}
