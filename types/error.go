package types

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/userhubdev/go-sdk/apiv1"
	"github.com/userhubdev/go-sdk/code"
)

// UserHubError is the UserHub API error type.
type UserHubError struct {
	apiCode    code.Code
	message    string
	reason     string
	param      string
	metadata   map[string]string
	call       string
	statusCode int
	cause      error
}

func NewError(format string, a ...any) *UserHubError {
	var message string
	var cause error

	if len(a) > 0 {
		for _, e := range a {
			c, ok := e.(error)
			if ok && c != nil {
				cause = c
				break
			}
		}

		if cause != nil {
			format = strings.ReplaceAll(format, "%w", "%v")
		}

		message = fmt.Sprintf(format, a...)
	} else {
		message = format
	}

	return &UserHubError{
		message: message,
		cause:   cause,
	}
}

func NewErrorFromStatus(
	call string,
	message string,
	status *apiv1.Status,
	statusCode int,
	cause error,
) *UserHubError {
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

		e.apiCode = code.Code(status.Code)
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
		p.WriteString(string(e.apiCode))
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

func (e *UserHubError) clone() *UserHubError {
	if e != nil {
		e2 := *e
		return &e2
	} else {
		return &UserHubError{}
	}
}

// ApiCode returns general error code (e.g. `INVALID_ARGUMENT`).
func (e *UserHubError) ApiCode() code.Code {
	if e == nil || !e.apiCode.Valid() {
		return "UNKNOWN"
	}
	return e.apiCode
}

// SetApiCode return a new error with the specified API code.
func (e *UserHubError) SetApiCode(v code.Code) *UserHubError {
	e2 := e.clone()
	e2.apiCode = v
	return e2
}

// Message returns the developer-facing error message.
func (e *UserHubError) Message() string {
	if e == nil || e.message == "" {
		return "<no message>"
	}
	return e.message
}

// SetMessage return a new error with the specified message.
func (e *UserHubError) SetMessage(v string) *UserHubError {
	e2 := e.clone()
	e2.message = v
	return e2
}

// Reason returns the reason code for the error (e.g. `USER_PENDING_DELETION`).
func (e *UserHubError) Reason() string {
	if e == nil {
		return ""
	}
	return e.reason
}

// SetReason return a new error with the specified reason.
func (e *UserHubError) SetReason(v string) *UserHubError {
	e2 := e.clone()
	e2.reason = v
	return e2
}

// Param returns the parameter path related to the error (e.g. `member.userId`).
func (e *UserHubError) Param() string {
	if e == nil {
		return ""
	}
	return e.param
}

// SetParam return a new error with the specified param.
func (e *UserHubError) SetParam(v string) *UserHubError {
	e2 := e.clone()
	e2.param = v
	return e2
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

// SetMetadata return a new error with the specified cause.
func (e *UserHubError) SetMetadata(v map[string]string) *UserHubError {
	e2 := e.clone()
	if len(v) > 0 {
		e2.metadata = map[string]string{}

		for key, value := range v {
			e2.metadata[key] = value
		}
	} else {
		e2.metadata = nil
	}
	return e2
}

// Call returns the method call name for the error (e.g. `admin.users.get`).
func (e *UserHubError) Call() string {
	if e == nil {
		return ""
	}
	return e.call
}

// SetCall return a new error with the specified call.
func (e *UserHubError) SetCall(v string) *UserHubError {
	e2 := e.clone()
	e2.call = v
	return e2
}

// StatusCode returns the HTTP status code (e.g. 500).
func (e *UserHubError) StatusCode() int {
	if e == nil {
		return 0
	}
	return e.statusCode
}

// SetStatusCode return a new error with the specified status code.
func (e *UserHubError) SetStatusCode(v int) *UserHubError {
	e2 := e.clone()
	e2.statusCode = v
	return e2
}

// Unwrap returns the underlying error.
func (e *UserHubError) Unwrap() error {
	if e != nil {
		return e.cause
	}
	return nil
}

// SetCause return a new error with the specified cause.
func (e *UserHubError) SetCause(v error) *UserHubError {
	e2 := e.clone()
	e2.cause = v
	return e2
}

// MarshalJSON returns the underlying error.
func (e *UserHubError) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"code":    e.ApiCode(),
		"message": e.Message(),
	})
}
