// Code generated. DO NOT EDIT.

package userapi

import (
	"net/http"
	"strings"

	"github.com/userhubdev/go-sdk/internal"
	"github.com/userhubdev/go-sdk/internal/options"
	"github.com/userhubdev/go-sdk/option"
	"github.com/userhubdev/go-sdk/types"
)

type Error = types.UserHubError

// New returns a new User API client.
func New(userKey, accessToken string, opts ...option.ClientOption) (Client, error) {
	o := &options.ClientOptions{
		BaseUrl: internal.ApiBaseUrl,
		Headers: http.Header{},
	}
	for _, opt := range opts {
		if opt != nil {
			opt.Apply(o)
		}
	}
	if err := o.Validate(); err != nil {
		return nil, err
	}

	userKey = strings.TrimSpace(userKey)
	if userKey == "" {
		return nil, internal.CallErrorf(nil, nil, "userKey required")
	} else if !strings.HasPrefix(userKey, internal.UserKeyPrefix) {
		return nil, internal.CallErrorf(nil, nil, "userKey must start with `%s`", internal.UserKeyPrefix)
	} else {
		o.Headers.Set(internal.ApiKeyHeader, userKey)
	}

	accessToken = strings.TrimSpace(accessToken)
	if accessToken != "" {
		o.Headers.Set(internal.AuthHeader, "Bearer "+accessToken)
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
