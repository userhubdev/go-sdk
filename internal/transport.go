package internal

import "context"

type Transport interface {
	Execute(ctx context.Context, req *Request) (*Response, error)
}
