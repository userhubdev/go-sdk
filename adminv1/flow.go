// Code generated. DO NOT EDIT.

package adminv1

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
	//
	// This will not be set if the invitation was created by an admin.
	Creator *User `json:"creator"`
	// The start time of the flow.
	StartTime time.Time `json:"startTime"`
	// The time the flow will expire.
	ExpireTime time.Time `json:"expireTime"`
	// The expire duration of the flow.
	Ttl string `json:"ttl"`
	// The flow secret.
	//
	// This is only populated on create.
	Secret string `json:"secret"`
	// The join organization flow type specific data.
	JoinOrganization *JoinOrganizationFlow `json:"joinOrganization"`
	// The signup flow type specific data
	Signup *SignupFlow `json:"signup"`
	// The flow view.
	View string `json:"view"`
	// The creation time of the flow.
	CreateTime time.Time `json:"createTime"`
	// The last update time of the flow.
	UpdateTime time.Time `json:"updateTime"`
}
