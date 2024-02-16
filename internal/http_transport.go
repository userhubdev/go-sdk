package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/userhubdev/go-sdk/apiv1"
	"github.com/userhubdev/go-sdk/code"
	"github.com/userhubdev/go-sdk/types"
)

var (
	defaultTransport http.RoundTripper = &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   ConnectTimeout,
			KeepAlive: 15 * time.Second,
		}).DialContext,
		MaxConnsPerHost:       MaxConns,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          MaxIdleConns,
		IdleConnTimeout:       IdleConnTimeout,
		TLSHandshakeTimeout:   TlsTimeout,
		ExpectContinueTimeout: 1 * time.Second,
	}

	defaultClient = &http.Client{
		Transport: defaultTransport,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Timeout: RequestTimeout,
	}

	removeWhitespace = regexp.MustCompile(`\s+`)
)

type HttpTransport struct {
	baseUrl string
	headers http.Header
}

func NewHttpTransport(
	baseUrl string,
	headers http.Header,
) *HttpTransport {
	if headers == nil {
		headers = http.Header{}
	} else {
		headers = headers.Clone()
	}

	headers.Set("user-agent", UserAgent)

	return &HttpTransport{
		baseUrl: baseUrl,
		headers: headers,
	}
}

func (t *HttpTransport) attempt(
	ctx context.Context,
	req *Request,
	reqUrl string,
	headers http.Header,
	body []byte,
) (*Response, error) {
	req.attempt++

	var reqReader io.Reader
	if len(body) > 0 {
		reqReader = bytes.NewReader(body)
	}

	httpReq, err := http.NewRequestWithContext(ctx, req.method, reqUrl, reqReader)
	if err != nil {
		return nil, CallErrorf(req, nil, "failed to create request: %w", err)
	}
	httpReq.Header = headers.Clone()

	httpRes, err := defaultClient.Do(httpReq)
	if err != nil {
		return nil, CallErrorf(req, httpRes, "failed to execute request: %w", err)
	}
	defer httpRes.Body.Close()

	resBody, err := io.ReadAll(&io.LimitedReader{R: httpRes.Body, N: MaxBodySizeBytes})
	if err != nil {
		return nil, CallErrorf(req, httpRes, "failed to read response body: %w", err)
	}

	if httpRes.StatusCode/100 != 2 {
		if strings.Contains(httpRes.Header.Get("content-type"), "json") {
			status := &apiv1.Status{}

			err = json.Unmarshal(resBody, &status)
			if err != nil {
				return nil, CallErrorf(req, httpRes, "failed to decode API%s", summarizeBody(resBody))
			}

			return nil, types.NewErrorFromStatus(req.call, "", status, httpRes.StatusCode, nil)
		} else if httpRes.StatusCode == 429 {
			return nil, CallErrorf(req, httpRes, "API call rate limited").
				SetApiCode(code.ResourceExhausted)
		} else {
			return nil, CallErrorf(req, httpRes, "API returned non-JSON error%s", summarizeBody(resBody))
		}
	}

	return &Response{res: httpRes, body: resBody}, nil
}

func (t *HttpTransport) Execute(ctx context.Context, req *Request) (*Response, error) {
	ctx, cancel := context.WithTimeout(ctx, req.timeout)
	defer cancel()

	var reqUrl string
	if len(req.query) > 0 {
		u, err := url.Parse(t.baseUrl + req.path)
		if err != nil {
			return nil, CallErrorf(req, nil, "failed to parse URL: %w", err)
		}

		q := u.Query()
		for key, value := range req.query {
			q.Set(key, value)
		}
		u.RawQuery = q.Encode()

		reqUrl = u.String()
	} else {
		reqUrl = t.baseUrl + req.path
	}

	headers := t.headers.Clone()
	if len(req.headers) > 0 {
		for key := range req.headers {
			headers.Set(key, req.headers.Get(key))
		}
	}

	var body []byte
	var err error
	if req.body != nil {
		body, err = json.Marshal(req.body)
		if err != nil {
			return nil, CallErrorf(req, nil, "failed to marshal body: %w", err)
		}
		headers.Set("content-type", "application/json")
	}

	for {
		res, err := t.attempt(ctx, req, reqUrl, headers, body)
		if err != nil {
			timeout := req.retry(ctx, err)
			if timeout > 0 {
				t := time.NewTimer(timeout)

				select {
				case <-t.C:
					continue
				case <-ctx.Done():
					t.Stop()
					if ctx.Err() != nil {
						return nil, fmt.Errorf("%w: %w", ctx.Err(), err)
					}
				}
			}

			return nil, err
		}

		return res, nil
	}
}
