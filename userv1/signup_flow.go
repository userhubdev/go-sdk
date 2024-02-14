// Code generated. DO NOT EDIT.

package userv1

// The signup flow.
type SignupFlow struct {
	// The email address of the invitee.
	Email string `json:"email"`
	// The display name of the invitee.
	DisplayName string `json:"displayName"`
	// Whether to create an organization as part of the signup flow.
	CreateOrganization bool `json:"createOrganization"`
}
