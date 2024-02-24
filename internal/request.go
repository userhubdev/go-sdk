package internal

import (
	"context"
	"fmt"
	"io"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/userhubdev/go-sdk/types"
)

type Request struct {
	call   string
	method string
	path   string

	headers http.Header
	query   map[string]string
	body    any

	attempt    int
	idempotent bool
	timeout    time.Duration
}

func NewRequest(call, method, path string) *Request {
	return &Request{
		call:   call,
		method: method,
		path:   path,

		timeout: time.Minute,
	}
}

func (r *Request) Call() string   { return r.call }
func (r *Request) Method() string { return r.method }
func (r *Request) Path() string   { return r.path }

func (r *Request) SetBody(value any) {
	r.body = value
}

func (r *Request) SetHeader(name, value string) {
	if r.headers == nil {
		r.headers = http.Header{}
	}
	r.headers.Set(name, value)
}

func (r *Request) SetIdempotent(v bool) {
	r.idempotent = v
}

func (r *Request) SetQuery(name string, value any) {
	var param string

	switch v := value.(type) {
	case bool:
		if v {
			param = "true"
		} else {
			param = "false"
		}
	case *bool:
		if v != nil {
			if *v {
				param = "true"
			} else {
				param = "false"
			}
		}
	case int:
		param = strconv.FormatInt(int64(v), 10)
	case *int:
		if v != nil {
			param = strconv.FormatInt(int64(*v), 10)
		}
	case int32:
		param = strconv.FormatInt(int64(v), 10)
	case *int32:
		if v != nil {
			param = strconv.FormatInt(int64(*v), 10)
		}
	case int64:
		param = strconv.FormatInt(v, 10)
	case *int64:
		if v != nil {
			param = strconv.FormatInt(*v, 10)
		}
	case string:
		param = v
	case *string:
		if v != nil {
			param = *v
		}
	case time.Time:
		if !v.IsZero() {
			param = v.Format(time.RFC3339Nano)
		}
	case *time.Time:
		if v != nil && !v.IsZero() {
			param = v.Format(time.RFC3339Nano)
		}
	default:
		if v != nil {
			param = fmt.Sprintf("%v", v)
		}
	}

	if param != "" {
		if r.query == nil {
			r.query = map[string]string{}
		}

		r.query[name] = param
	}
}

func (r *Request) shouldRetry(err error) bool {
	if err == nil {
		return false
	}

	if ue, ok := err.(*types.UserHubError); ok && ue != nil {
		sc := ue.StatusCode()

		if sc == 429 {
			return true
		}

		if r.idempotent {
			if sc >= 500 && sc <= 599 {
				return true
			}
		}
	}

	if r.idempotent {
		if err == io.ErrUnexpectedEOF {
			return true
		}

		if osShouldRetry(err) {
			return true
		}

		if err, ok := err.(interface{ Temporary() bool }); ok {
			if err.Temporary() {
				return true
			}
		}
	}

	if err, ok := err.(interface{ Unwrap() error }); ok {
		return r.shouldRetry(err.Unwrap())
	}

	if errs, ok := err.(interface{ Unwrap() []error }); ok {
		for _, err := range errs.Unwrap() {
			if r.shouldRetry(err) {
				return true
			}
		}
	}

	return false
}

func (r *Request) retry(ctx context.Context, err error) time.Duration {
	if r.attempt > RetryMaxAttempts {
		return 0
	}

	if !r.shouldRetry(err) {
		return 0
	}

	timeout := time.Duration(math.Pow(2, float64(r.attempt-1))) * RetryMultiplier
	if timeout > RetryMaxSleep {
		timeout = RetryMaxSleep
	}

	deadline, ok := ctx.Deadline()
	if ok && time.Until(deadline)+RetryOverhead < timeout {
		return 0
	}

	return timeout
}
