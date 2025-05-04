// Code generated. DO NOT EDIT.

package adminv1

import (
	"time"
)

// A member of an organization.
type Member struct {
	// Whether the membership is active.
	State string `json:"state"`
	// The user.
	User *User `json:"user"`
	// The user's role within the organization.
	Role *Role `json:"role"`
	// The seat associated with the member.
	//
	// This will be absent if there is no active
	// subscription for the organization or the user
	// has not been assigned a seat.
	Seat *AccountSubscriptionSeat `json:"seat"`
	// The member view.
	View string `json:"view"`
	// The creation time of the membership.
	CreateTime time.Time `json:"createTime"`
	// The last update time of the membership.
	UpdateTime time.Time `json:"updateTime"`
}
