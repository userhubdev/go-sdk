package webhook

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"reflect"

	"github.com/userhubdev/go-sdk/internal"
	"github.com/userhubdev/go-sdk/types"
)

func defaultOnError(ctx context.Context, err error) {
	if err != nil {
		log.Println("Userhub webhook:", err.Error())
	}
}

func isNil(value any) bool {
	if value == nil {
		return true
	}

	switch reflect.TypeOf(value).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(value).IsNil()
	default:
		return false
	}
}

func createResponse(value any) (*Response, error) {
	r := &Response{
		StatusCode: http.StatusOK,
		Headers:    http.Header{},
	}

	r.Headers.Set("Content-Type", "application/json")
	r.Headers.Set(internal.WebhookAgentHeader, internal.UserAgent)

	var err error

	if !isNil(value) {
		switch v := value.(type) {
		case []byte:
			r.Body = v
		case error:
			e := &types.UserHubError{}
			if !errors.As(v, &e) {
				e = errInternal
			}

			r.StatusCode = e.ApiCode().WebhookStatusCode()

			r.Body, err = json.Marshal(e)
		case *Response:
			return v, nil
		default:
			r.Body, err = json.Marshal(v)
		}
	}
	if err != nil {
		r.StatusCode = http.StatusInternalServerError
		r.Body = []byte(internal.WebhookServerErrorJson)
	} else if len(r.Body) == 0 || string(r.Body) == "null" {
		r.Body = []byte("{}")
	}

	return r, err
}
