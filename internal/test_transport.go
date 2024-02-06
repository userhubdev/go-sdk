package internal

import (
	"context"
	"net/http"
)

type TestTransport struct {
	Request *Request
	Body    string
	Error   error
}

func (t *TestTransport) Execute(ctx context.Context, req *Request) (*Response, error) {
	t.Request = req

	if t.Error != nil {
		return nil, t.Error
	}

	return &Response{
		req:  req,
		res:  &http.Response{StatusCode: 200},
		body: []byte(t.Body),
	}, nil
}
