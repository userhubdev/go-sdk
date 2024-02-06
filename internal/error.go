package internal

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/userhubdev/go-sdk/types"
)

func Errorf(req *Request, res *http.Response, format string, a ...any) error {
	var cause error
	var call string

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

	if req != nil {
		call = req.call
	}

	var statusCode int

	if res != nil {
		statusCode = res.StatusCode
	}

	return types.NewUserHubError(call, fmt.Sprintf(format, a...), nil, statusCode, cause)
}
