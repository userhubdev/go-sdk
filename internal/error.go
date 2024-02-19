package internal

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/userhubdev/go-sdk/types"
)

func CallErrorf(req *Request, res *http.Response, format string, a ...any) *types.UserHubError {
	msg, cause := createMsgCause(format, a)

	var call string
	if req != nil {
		call = req.call
	}

	var statusCode int
	if res != nil {
		statusCode = res.StatusCode
	}

	return types.NewErrorFromStatus(call, msg, nil, statusCode, cause)
}

func createMsgCause(format string, a []any) (string, error) {
	if len(a) == 0 {
		return format, nil
	}

	var cause error

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

	return fmt.Sprintf(format, a...), cause
}
