package webhook

import (
	"net/http"
	"strings"

	"github.com/userhubdev/go-sdk/internal"
	"github.com/userhubdev/go-sdk/types"
)

// Request is a wrapper for webhook requests.
type Request struct {
	Headers http.Header
	Body    []byte
}

// GetAction returns the webhook action.
func (r *Request) GetAction() (string, error) {
	var name string

	if r != nil {
		name = r.Headers.Get(internal.WebhookActionHeader)
	}

	if name == "" {
		return "", types.NewError("%s header is missing", internal.WebhookActionHeader)
	}

	return name, nil
}

// GetTimestamp returns the webhook timestamp.
func (r *Request) GetTimestamp() (string, error) {
	timestamp := r.Headers.Get(internal.WebhookTimestampHeader)
	if timestamp == "" {
		return "", types.NewError("%s header is missing", internal.WebhookTimestampHeader)
	}

	return timestamp, nil
}

// GetSignatures returns the webhook signatures.
func (r *Request) GetSignatures() ([]string, error) {
	var signatures []string

	headers := r.Headers.Values(internal.WebhookSignatureHeader)

	if len(headers) == 0 {
		return nil, types.NewError("%s header is missing", internal.WebhookSignatureHeader)
	}

	for _, header := range headers {
		header = strings.TrimSpace(header)
		if header == "" {
			continue
		}

		if !strings.Contains(header, ",") {
			signatures = append(signatures, header)
			continue
		}

		headerParts := strings.Split(header, ",")

		for _, signature := range headerParts {
			signature = strings.TrimSpace(signature)
			if signature != "" {
				signatures = append(signatures, signature)
			}
		}
	}

	if len(signatures) == 0 {
		return nil, types.NewError("%s header normalized to nothing", internal.WebhookSignatureHeader)
	}

	return signatures, nil
}

// Response is a wrapper for webhook responses.
type Response struct {
	StatusCode int
	Headers    http.Header
	Body       []byte
}
