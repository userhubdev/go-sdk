// Code generated. DO NOT EDIT.

package adminv1

// The delegated connection.
type ConnectionDelegate struct {
	// The delegated connection identifier.
	Id string `json:"id"`
	// The client defined unique identifier of the delegated connection.
	UniqueId string `json:"uniqueId"`
	// The human-readable display name of the delegated connection.
	DisplayName string `json:"displayName"`
	// The current status of the delegated connection.
	State string `json:"state"`
	// The code that best describes the reason for the state.
	StateReason string `json:"stateReason"`
	// The delegated connection type.
	Type string `json:"type"`
}
