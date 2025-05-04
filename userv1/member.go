// Code generated. DO NOT EDIT.

package userv1

// A member of an organization.
type Member struct {
	// The user.
	User *User `json:"user"`
	// The user's role within the organization.
	Role *Role `json:"role"`
	// The seat assigned to the member.
	//
	// This will be absent if there is no active
	// subscription for the organization or the user
	// has not been assigned a seat.
	Seat *AccountSubscriptionSeat `json:"seat"`
}
