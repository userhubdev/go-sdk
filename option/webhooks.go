package option

import (
	"context"

	"github.com/userhubdev/go-sdk/internal/options"
)

type WebhooksOption interface {
	Apply(webhooksOptions *options.WebhooksOptions)
}

func WithOnError(fn func(ctx context.Context, err error)) WebhooksOption {
	return withOnError{fn: fn}
}

type withOnError struct {
	fn func(ctx context.Context, err error)
}

func (w withOnError) Apply(options *options.WebhooksOptions) {
	options.OnError = w.fn
}
