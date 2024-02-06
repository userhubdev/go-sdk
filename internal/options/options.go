package options

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/userhubdev/go-sdk/internal"
)

type ClientOptions struct {
	BaseUrl string
	Headers http.Header
}

func (o *ClientOptions) Validate() error {
	if o == nil {
		return errors.New("userhub: client options required")
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
