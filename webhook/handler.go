package webhook

import (
	"context"
	"encoding/json"

	"github.com/userhubdev/go-sdk/code"
	"github.com/userhubdev/go-sdk/types"
)

// Handler is the interface used for implementing webhook handlers.
type Handler interface {
	HandleAction(ctx context.Context, req *Request) (*Response, error)
}

// The HandlerFunc type is an adapter to allow the use of ordinary
// functions as action handlers.
type HandlerFunc func(ctx context.Context, name string, input []byte) (any, error)

func (h HandlerFunc) HandleAction(ctx context.Context, req *Request) (*Response, error) {
	name, err := req.GetAction()
	if err != nil {
		return nil, err
	}

	data, err := h(ctx, name, req.Body)
	if err != nil {
		return nil, err
	}
	return createResponse(data)
}

type challengeHandler struct{}

func (e challengeHandler) HandleAction(ctx context.Context, req *Request) (*Response, error) {
	return createResponse(req.Body)
}

type unimplementedHandler struct{}

func (e unimplementedHandler) HandleAction(ctx context.Context, req *Request) (*Response, error) {
	name, err := req.GetAction()
	if err != nil {
		return nil, err
	}

	return nil, types.NewError("Handler not implemented: %s", name).
		SetApiCode(code.Unimplemented)
}

type genericHandler[I any, O any] func(ctx context.Context, input *I) (*O, error)

func (e genericHandler[I, O]) HandleAction(ctx context.Context, req *Request) (*Response, error) {
	i := new(I)

	err := json.Unmarshal(req.Body, &i)
	if err != nil {
		return nil, err
	}

	data, err := e(ctx, i)
	if err != nil {
		return nil, err
	}

	return createResponse(data)
}

type genericInputOnlyHandler[I any] func(ctx context.Context, input *I) error

func (e genericInputOnlyHandler[I]) HandleAction(ctx context.Context, req *Request) (*Response, error) {
	i := new(I)

	err := json.Unmarshal(req.Body, &i)
	if err != nil {
		return nil, err
	}

	return nil, e(ctx, i)
}
