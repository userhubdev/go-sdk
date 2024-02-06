// Code generated. DO NOT EDIT.

package adminv1

// The user associated with event.
type EventUser struct {
	// The system-assigned identifier of the user.
	Id string `json:"id"`
	// The human-readable display name of the user.
	//
	// NOTE: this is the current display name and not
	// the one as of the time of the event.
	DisplayName string `json:"displayName"`
	// The email address of the user.
	//
	// NOTE: this is the current email and not
	// the one as of the time of the event.
	Email string `json:"email"`
}
