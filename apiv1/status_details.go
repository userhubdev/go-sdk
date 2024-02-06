// Code generated. DO NOT EDIT.

package apiv1

// The API error with the fields that already exist
// in Status removed.
type StatusDetails struct {
	// A reason code for the error (e.g. `PENDING_DELETION`).
	Reason string `json:"reason"`
	// The parameter path related to the error (e.g. `member.userId`).
	Param string `json:"param"`
	// Additional metadata related to the error.
	Metadata map[string]string `json:"metadata"`
}
