// Code generated. DO NOT EDIT.

package userv1

// Individual account.
type User struct {
	// The system-assigned identifier of the user.
	Id string `json:"id"`
	// The client defined unique identifier of the user account.
	UniqueId string `json:"uniqueId"`
	// The human-readable display name of the user.
	DisplayName string `json:"displayName"`
	// The email address of the user.
	Email string `json:"email"`
	// Whether the user's email address has been verified.
	EmailVerified bool `json:"emailVerified"`
	// The photo/avatar URL of the user.
	ImageUrl string `json:"imageUrl"`
	// Whether the user is disabled.
	Disabled bool `json:"disabled"`
}
