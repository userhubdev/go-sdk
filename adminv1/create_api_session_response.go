// Code generated. DO NOT EDIT.

package adminv1

import (
	"time"
)

// Response message for CreateApiSession.
type CreateApiSessionResponse struct {
	// An authorization token which can be used to access the User API.
	//
	// This should be included in the `Authorization` header with a
	// `Bearer` prefix.
	AccessToken string `json:"accessToken"`
	// The expiration time for the session.
	ExpireTime time.Time `json:"expireTime"`
}
