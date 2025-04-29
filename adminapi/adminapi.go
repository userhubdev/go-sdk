// Code generated. DO NOT EDIT.

package adminapi

import (
	"net/http"
	"strings"

	"github.com/userhubdev/go-sdk/internal"
	"github.com/userhubdev/go-sdk/internal/options"
	"github.com/userhubdev/go-sdk/option"
	"github.com/userhubdev/go-sdk/types"
)

type Error = types.UserHubError

// New returns a new Admin API client.
func New(adminKey string, opts ...option.ClientOption) (Client, error) {
	o := &options.ClientOptions{
		ApiVersion: internal.ApiVersion,
		BaseUrl:    internal.ApiBaseUrl,
		Headers:    http.Header{},
	}
	for _, opt := range opts {
		if opt != nil {
			opt.Apply(o)
		}
	}
	if err := o.Validate(); err != nil {
		return nil, err
	}

	o.Headers.Set(internal.ApiVersionHeader, o.ApiVersion)

	adminKey = strings.TrimSpace(adminKey)
	if adminKey == "" {
		return nil, internal.CallErrorf(nil, nil, "adminKey required")
	} else {
		o.Headers.Set(internal.AuthHeader, "Bearer "+adminKey)
	}

	transport := internal.NewHttpTransport(o.BaseUrl, o.Headers)

	return newClient(transport), nil
}

// Some is used to set an Optional value in update methods.
func Some[T any](v T) types.Optional[T] {
	return types.Optional[T]{
		Value:   v,
		Present: true,
	}
}
