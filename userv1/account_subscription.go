// Code generated. DO NOT EDIT.

package userv1

import (
	"time"
)

// The account view of the subscription.
type AccountSubscription struct {
	// The system-assigned identifier of the subscription.
	Id string `json:"id"`
	// The state of the subscription.
	State string `json:"state"`
	// The anchor time of the billing cycle.
	AnchorTime time.Time `json:"anchorTime"`
	// The subscription plan.
	Plan *AccountSubscriptionPlan `json:"plan"`
	// The user's seat.
	//
	// This will only be set for organization subscriptions where
	// the user has been assigned a seat.
	Seat *AccountSubscriptionSeat `json:"seat"`
}
