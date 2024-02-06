// Code generated. DO NOT EDIT.

package connectionsv1

// The user object for the Custom Users connection.
type CustomUser struct {
	// The external identifier for the user.
	Id string `json:"id"`
	// The human-readable display name of the user.
	//
	// The maximum length is 200 characters.
	DisplayName string `json:"displayName"`
	// The email address of the user.
	//
	// The maximum length is 320 characters.
	Email string `json:"email"`
	// Whether the user's email address has been verified.
	EmailVerified bool `json:"emailVerified"`
	// The E164 phone number for the user (e.g. `+12125550123`).
	PhoneNumber string `json:"phoneNumber"`
	// Whether the user's phone number has been verified.
	PhoneNumberVerified bool `json:"phoneNumberVerified"`
	// The photo/avatar URL of the user.
	//
	// The maximum length is 2000 characters.
	ImageUrl string `json:"imageUrl"`
	// Whether the user is disabled.
	Disabled bool `json:"disabled"`
}
