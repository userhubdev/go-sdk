// Code generated. DO NOT EDIT.

package userv1

import (
	"time"
)

// The session details.
type Session struct {
	// The authenticated user.
	//
	// This will be null if the user is not authenticated.
	User *User `json:"user"`
	// The authenticated user's organization memberships.
	Memberships []*Membership `json:"memberships"`
	// The user's default active individual subscription.
	//
	// See memberships for organization subscription and
	// seat information.
	Subscription *AccountSubscription `json:"subscription"`
	// The expiration time for the current session.
	ExpireTime time.Time `json:"expireTime"`
	// The scopes available in the current session.
	Scopes []string `json:"scopes"`
}
