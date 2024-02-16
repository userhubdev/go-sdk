// Code generated. DO NOT EDIT.

package webhook

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/userhubdev/go-sdk/code"
	"github.com/userhubdev/go-sdk/eventsv1"
	"github.com/userhubdev/go-sdk/internal"
	"github.com/userhubdev/go-sdk/option"
	"github.com/userhubdev/go-sdk/types"
)

func TestWebhook_Handle(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	headers := func(input map[string]string) http.Header {
		output := http.Header{}

		for name, value := range input {
			output.Set(name, value)
		}

		return output
	}

	tests := []struct {
		Name         string
		Secret       string
		SetTimestamp bool
		SetSignature bool
		AddSignature bool
		Request      Request
		Response     Response
	}{

		{
			Name:    "Signing secret is required",
			Secret:  "",
			Request: Request{},
			Response: Response{
				StatusCode: 500,
				Body:       []byte(`{"message":"Signing secret is required","code":"UNKNOWN"}`),
			},
		},
		{
			Name:    "Headers are required",
			Secret:  "test",
			Request: Request{},
			Response: Response{
				StatusCode: 500,
				Body:       []byte(`{"message":"Headers are required","code":"UNKNOWN"}`),
			},
		},
		{
			Name:   "Body is required",
			Secret: "test",
			Request: Request{
				Headers: headers(map[string]string{
					"content-type": "application/json",
				}),
			},
			Response: Response{
				StatusCode: 500,
				Body:       []byte(`{"message":"Body is required","code":"UNKNOWN"}`),
			},
		},
		{
			Name:   "UserHub-Action header is missing",
			Secret: "test",
			Request: Request{
				Headers: headers(map[string]string{
					"content-type": "application/json",
				}),
				Body: []byte(`{}`),
			},
			Response: Response{
				StatusCode: 500,
				Body:       []byte(`{"message":"UserHub-Action header is missing","code":"UNKNOWN"}`),
			},
		},
		{
			Name:   "UserHub-Timestamp header is missing",
			Secret: "test",
			Request: Request{
				Headers: headers(map[string]string{
					"UserHub-Action": "events.handle",
				}),
				Body: []byte(`{}`),
			},
			Response: Response{
				StatusCode: 500,
				Body:       []byte(`{"message":"UserHub-Timestamp header is missing","code":"UNKNOWN"}`),
			},
		},
		{
			Name:         "UserHub-Signature header is missing",
			Secret:       "test",
			SetTimestamp: true,
			Request: Request{
				Headers: headers(map[string]string{
					"UserHub-Action": "events.handle",
				}),
				Body: []byte(`{}`),
			},
			Response: Response{
				StatusCode: 500,
				Body:       []byte(`{"message":"UserHub-Signature header is missing","code":"UNKNOWN"}`),
			},
		},
		{
			Name:         "Signatures normalized to nothing",
			Secret:       "test",
			SetTimestamp: true,
			Request: Request{
				Headers: headers(map[string]string{
					"UserHub-Action":    "events.handle",
					"UserHub-Signature": ",",
				}),
				Body: []byte(`{}`),
			},
			Response: Response{
				StatusCode: 500,
				Body:       []byte(`{"message":"UserHub-Signature header normalized to nothing","code":"UNKNOWN"}`),
			},
		},
		{
			Name:   "Timestamp is invalid",
			Secret: "test",
			Request: Request{
				Headers: headers(map[string]string{
					"UserHub-Action":    "events.handle",
					"UserHub-Timestamp": "timestamp",
					"UserHub-Signature": "fail",
				}),
				Body: []byte(`{}`),
			},
			Response: Response{
				StatusCode: 500,
				Body:       []byte(`{"message":"Timestamp is invalid: timestamp","code":"UNKNOWN"}`),
			},
		},
		{
			Name:   "Timestamp is too far in the past",
			Secret: "test",
			Request: Request{
				Headers: headers(map[string]string{
					"UserHub-Action":    "events.handle",
					"UserHub-Timestamp": "1",
					"UserHub-Signature": "fail",
				}),
				Body: []byte(`{}`),
			},
			Response: Response{
				StatusCode: 500,
				Body:       []byte(`{"message":"Timestamp is too far in the past: 1","code":"UNKNOWN"}`),
			},
		},
		{
			Name:   "Timestamp is too far in the past",
			Secret: "test",
			Request: Request{
				Headers: headers(map[string]string{
					"UserHub-Action":    "events.handle",
					"UserHub-Timestamp": "5000000000",
					"UserHub-Signature": "fail",
				}),
				Body: []byte(`{}`),
			},
			Response: Response{
				StatusCode: 500,
				Body:       []byte(`{"message":"Timestamp is too far in the future: 5000000000","code":"UNKNOWN"}`),
			},
		},
		{
			Name:         "Signatures do not match",
			Secret:       "test",
			SetTimestamp: true,
			Request: Request{
				Headers: headers(map[string]string{
					"UserHub-Action":    "events.handle",
					"UserHub-Signature": "fail",
				}),
				Body: []byte(`{}`),
			},
			Response: Response{
				StatusCode: 500,
				Body:       []byte(`{"message":"Signatures do not match","code":"UNKNOWN"}`),
			},
		},
		{
			Name:         "Challenge",
			Secret:       "test",
			SetTimestamp: true,
			AddSignature: true,
			Request: Request{
				Headers: headers(map[string]string{
					"UserHub-Action": "challenge",
				}),
				Body: []byte(`{"challenge": "some-value"}`),
			},
			Response: Response{
				StatusCode: 200,
				Body:       []byte(`{"challenge": "some-value"}`),
			},
		},
		{
			Name:         "Handle multiple signature headers",
			Secret:       "test",
			SetTimestamp: true,
			AddSignature: true,
			Request: Request{
				Headers: headers(map[string]string{
					"UserHub-Action":    "challenge",
					"UserHub-Signature": "fail",
				}),
				Body: []byte(`{"challenge": "some-value"}`),
			},
			Response: Response{
				StatusCode: 200,
				Body:       []byte(`{"challenge": "some-value"}`),
			},
		},
		{
			Name:         "Handle combined signature headers",
			Secret:       "test",
			SetTimestamp: true,
			SetSignature: true,
			Request: Request{
				Headers: headers(map[string]string{
					"UserHub-Action":    "challenge",
					"UserHub-Signature": "fail, {signature}",
				}),
				Body: []byte(`{"challenge": "some-value"}`),
			},
			Response: Response{
				StatusCode: 200,
				Body:       []byte(`{"challenge": "some-value"}`),
			},
		},
		{
			Name:         "Handler not implemented",
			Secret:       "test",
			SetTimestamp: true,
			AddSignature: true,
			Request: Request{
				Headers: headers(map[string]string{
					"UserHub-Action": "unknown",
				}),
				Body: []byte(`{}`),
			},
			Response: Response{
				StatusCode: 501,
				Body:       []byte(`{"message":"Handler not implemented: unknown","code":"UNIMPLEMENTED"}`),
			},
		},
		{
			Name:         "Event handler succeeds",
			Secret:       "test",
			SetTimestamp: true,
			AddSignature: true,
			Request: Request{
				Headers: headers(map[string]string{
					"UserHub-Action": "events.handle",
				}),
				Body: []byte(`{"type": "ok"}`),
			},
			Response: Response{
				StatusCode: 200,
				Body:       []byte(`{}`),
			},
		},
		{
			Name:         "Event handler errors",
			Secret:       "test",
			SetTimestamp: true,
			AddSignature: true,
			Request: Request{
				Headers: headers(map[string]string{
					"UserHub-Action": "events.handle",
				}),
				Body: []byte(`{"type": "fail"}`),
			},
			Response: Response{
				StatusCode: 400,
				Body:       []byte(`{"message":"Event failed: fail","code":"INVALID_ARGUMENT"}`),
			},
		},
	}

	noLogger := option.WithOnError(func(ctx context.Context, err error) {})

	for _, test := range tests {
		test := test

		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()

			wh := New(test.Secret, noLogger).
				OnEvent(func(ctx context.Context, input *eventsv1.Event) error {
					if input.Type != "ok" {
						return types.NewError("Event failed: %s", input.Type).
							SetApiCode(code.InvalidArgument)
					}

					return nil
				})

			req := test.Request
			req.Headers = req.Headers.Clone()

			if test.SetTimestamp {
				req.Headers.Set("UserHub-Timestamp", strconv.FormatInt(time.Now().Unix(), 10))
			}
			if test.SetSignature || test.AddSignature {
				mac := hmac.New(sha256.New, []byte(test.Secret))
				mac.Write([]byte(req.Headers.Get("UserHub-Timestamp")))
				mac.Write([]byte("."))
				mac.Write(req.Body)
				signature := hex.EncodeToString(mac.Sum(nil))

				if test.SetSignature {
					header := req.Headers.Get("UserHub-Signature")
					if strings.Contains(header, "{signature}") {
						header = strings.ReplaceAll(header, "{signature}", signature)
					} else {
						header = signature
					}
					req.Headers.Set("UserHub-Signature", header)
				} else {
					req.Headers.Add("UserHub-Signature", signature)
				}
			}

			res := wh.HandleAction(ctx, &req)
			require.Equal(t, test.Response.StatusCode, res.StatusCode, string(res.Body))

			require.Equal(t, "application/json", res.Headers.Get("content-type"))
			require.Equal(t, internal.UserAgent, res.Headers.Get("webhook-agent"))

			var expected map[string]any
			err := json.Unmarshal(test.Response.Body, &expected)
			require.NoError(t, err)

			var actual map[string]any
			err = json.Unmarshal(res.Body, &actual)
			require.NoError(t, err)

			require.Equal(t, expected, actual)
		})
	}
}
