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

func WithApiVersion(apiVersion string) ClientOption {
	return withApiVersion{apiVersion: apiVersion}
}

type withApiVersion struct{ apiVersion string }

func (w withApiVersion) Apply(options *options.ClientOptions) {
	if w.apiVersion != "" {
		options.ApiVersion = w.apiVersion
	}
}
