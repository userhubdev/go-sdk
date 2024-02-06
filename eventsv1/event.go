// Code generated. DO NOT EDIT.

package eventsv1

import (
	"time"
)

// The event.
type Event struct {
	// The event identifier
	Id string `json:"id"`
	// The event time.
	Time time.Time `json:"time"`
	// The event type.
	Type string `json:"type"`
	// The `flows.changed` data.
	FlowsChanged *FlowsChanged `json:"flowsChanged"`
	// The `members.changed` data.
	MembersChanged *MembersChanged `json:"membersChanged"`
	// The `organizations.changed` data.
	OrganizationsChanged *OrganizationsChanged `json:"organizationsChanged"`
	// The `subscriptions.changed` data.
	SubscriptionsChanged *SubscriptionsChanged `json:"subscriptionsChanged"`
	// The `users.changed` data.
	UsersChanged *UsersChanged `json:"usersChanged"`
}
