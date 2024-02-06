// Code generated. DO NOT EDIT.

package userv1

// Organization input parameters.
type OrganizationInput struct {
	// The client defined unique identifier of the organization account.
	//
	// It is restricted to letters, numbers, underscores, and hyphens,
	// with the first character a letter or a number, and a 255
	// character maximum.
	//
	// ID's starting with `org_` are reserved.
	UniqueId *string `json:"uniqueId"`
	// The human-readable display name of the organization.
	//
	// The maximum length is 200 characters.
	DisplayName *string `json:"displayName"`
	// The email address of the organization.
	//
	// The maximum length is 320 characters.
	Email *string `json:"email"`
}
