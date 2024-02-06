// Code generated. DO NOT EDIT.

package adminv1

// The connection associated with the event.
type EventConnection struct {
	// The system-assigned identifier of the connection.
	Id string `json:"id"`
	// The human-readable display name of the connection.
	//
	// NOTE: this is the current display name and not
	// the one as of the time of the event.
	DisplayName string `json:"displayName"`
	// The connection type.
	Type string `json:"type"`
}
