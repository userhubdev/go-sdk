package options

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/userhubdev/go-sdk/internal"
)

type ClientOptions struct {
	ApiVersion string
	BaseUrl    string
	Headers    http.Header
}

func (o *ClientOptions) Validate() error {
	if o == nil {
		return errors.New("userhub: client options required")
	}

	if o.ApiVersion == "" {
		return errors.New("userhub: API version required")
	} else if len(o.ApiVersion) != 10 {
		return errors.New("userhub: API version is invalid (e.g. 2022-11-15)")
	}

	if o.BaseUrl == "" {
		return errors.New("userhub: base URL required")
	} else if o.BaseUrl != internal.ApiBaseUrl {
		u, err := url.Parse(o.BaseUrl)
		if err != nil {
			return fmt.Errorf("userhub: base URL invalid: %w", err)
		}
		if u.Scheme != "http" && u.Scheme != "https" {
			return errors.New("userhub: base URL must have http(s) scheme")
		}
		if u.Hostname() == "" {
			return errors.New("userhub: base URL hostname required")
		}
	}

	return nil
}

type WebhooksOptions struct {
	OnError func(ctx context.Context, err error)
}
