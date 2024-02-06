// Code generated. DO NOT EDIT.

package adminv1

// The webhook specific connection data.
type WebhookConnection struct {
	// The URL of the events webhook.
	Url string `json:"url"`
	// The headers sent with requests to the connection URL.
	Headers map[string]string `json:"headers"`
	// The webhook secrets
	SigningSecrets []*SigningSecret `json:"signingSecrets"`
}
