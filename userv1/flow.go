// Code generated. DO NOT EDIT.

package userv1

import (
	"time"
)

// An invitation, task, or user request (e.g. join organization).
type Flow struct {
	// The system-assigned identifier of the flow.
	Id string `json:"id"`
	// The current state of the flow.
	State string `json:"state"`
	// The code that best describes the reason for the state.
	StateReason string `json:"stateReason"`
	// The flow type.
	Type string `json:"type"`
	// The target organization for the flow.
	Organization *Organization `json:"organization"`
	// The target user for the flow.
	User *User `json:"user"`
	// The user who created the flow.
	Creator *User `json:"creator"`
	// The time the flow will expires.
	ExpireTime time.Time `json:"expireTime"`
	// The creation time of the flow.
	CreateTime time.Time `json:"createTime"`
	// The join organization flow type specific data.
	JoinOrganization *JoinOrganizationFlow `json:"joinOrganization"`
	// The signup flow type specific data.
	Signup *SignupFlow `json:"signup"`
}
