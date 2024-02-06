// Code generated. DO NOT EDIT.

package adminv1

import (
	"time"
)

// The signing secret for the webhook.
type SigningSecret struct {
	// The signing secret value.
	Secret string `json:"secret"`
	// The time the signing secret is set to expire.
	ExpireTime time.Time `json:"expireTime"`
}
