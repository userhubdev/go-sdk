// Code generated. DO NOT EDIT.

package adminv1

// The actor associated with event.
type EventActor struct {
	// The system-assigned identifier of the actor.
	Id string `json:"id"`
	// The human-readable display name of the actor.
	//
	// NOTE: this is the current display name and not
	// the one as of the time of the event.
	DisplayName string `json:"displayName"`
	// The email address of the actor.
	//
	// NOTE: this is the current email and not
	// the one as of the time of the event.
	Email string `json:"email"`
	// Whether the actor is a tenant admin.
	Admin bool `json:"admin"`
}
