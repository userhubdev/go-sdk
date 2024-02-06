// Code generated. DO NOT EDIT.

package userv1

// The join organization flow.
type JoinOrganizationFlow struct {
	// The display name of the invitee.
	DisplayName string `json:"displayName"`
	// The email address of the invitee.
	//
	// This is required if a user isn't provided
	// or the user's email address is empty.
	Email string `json:"email"`
}
