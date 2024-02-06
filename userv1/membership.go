// Code generated. DO NOT EDIT.

package userv1

// A user's membership in an organization.
//
// This is the user view, see Member for the organizational
// view.
type Membership struct {
	// The organization.
	Organization *Organization `json:"organization"`
	// The user's role within the organization.
	Role *Role `json:"role"`
	// The subscription associated with the organization.
	Subscription *AccountSubscription `json:"subscription"`
}
