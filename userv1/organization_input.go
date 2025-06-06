// Code generated. DO NOT EDIT.

package userv1

import "github.com/userhubdev/go-sdk/types"

// Organization input parameters.
type OrganizationInput struct {
	// The client defined unique identifier of the organization account.
	//
	// It is restricted to letters, numbers, underscores, and hyphens,
	// with the first character a letter or a number, and a 255
	// character maximum.
	//
	// ID's starting with `org_` are reserved.
	UniqueId types.Optional[string] `json:"uniqueId"`
	// The human-readable display name of the organization.
	//
	// The maximum length is 200 characters.
	DisplayName types.Optional[string] `json:"displayName"`
	// The email address of the organization.
	//
	// The maximum length is 320 characters.
	Email types.Optional[string] `json:"email"`
	// The flow identifier associated with the creation of the organization.
	//
	// The flow type must be `SIGNUP` and associated with the user creating the organization.
	FlowId types.Optional[string] `json:"flowId"`
}
