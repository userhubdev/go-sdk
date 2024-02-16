package option

import (
	"github.com/userhubdev/go-sdk/internal/options"
)

type ClientOption interface {
	Apply(*options.ClientOptions)
}

func WithBaseUrl(baseUrl string) ClientOption {
	return withBaseUrl{baseUrl: baseUrl}
}

type withBaseUrl struct{ baseUrl string }

func (w withBaseUrl) Apply(options *options.ClientOptions) {
	if w.baseUrl != "" {
		options.BaseUrl = w.baseUrl
	}
}
