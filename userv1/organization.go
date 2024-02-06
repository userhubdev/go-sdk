// Code generated. DO NOT EDIT.

package userv1

// A group account.
type Organization struct {
	// The system-assigned identifier of the organization.
	Id string `json:"id"`
	// The client defined unique identifier of the organization account.
	UniqueId string `json:"uniqueId"`
	// The human-readable display name of the organization.
	DisplayName string `json:"displayName"`
	// The email address of the organization.
	Email string `json:"email"`
	// Whether the organization's email address has been verified.
	EmailVerified bool `json:"emailVerified"`
	// The photo/avatar URL of the organization.
	ImageUrl string `json:"imageUrl"`
	// Whether the organization is disabled.
	Disabled bool `json:"disabled"`
}
