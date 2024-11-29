// Code generated. DO NOT EDIT.

package userv1

import (
	"time"
)

// Response message for ExchangeToken.
type ExchangeSessionTokenResponse struct {
	// An authorization token which can be used to make authenticated
	// requests.
	//
	// This should be included in the `Authorization` header with a
	// `Bearer` prefix.
	AccessToken string `json:"accessToken"`
	// The expiration time for the session.
	ExpireTime time.Time `json:"expireTime"`
}
