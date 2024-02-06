// Code generated. DO NOT EDIT.

package adminv1

import (
	"time"
)

// A user's membership in an organization.
//
// This is the user view, see Member for the organizational
// view.
type Membership struct {
	// The organization.
	Organization *Organization `json:"organization"`
	// The user's role within the organization.
	Role *Role `json:"role"`
	// The seat associated with the membership.
	//
	// This will be absent if there is no active default
	// subscription for the organization or the user
	// has not been assigned a seat.
	Seat *AccountSubscriptionSeat `json:"seat"`
	// The creation time of the membership.
	CreateTime time.Time `json:"createTime"`
	// The last update time of the membership.
	UpdateTime time.Time `json:"updateTime"`
}
