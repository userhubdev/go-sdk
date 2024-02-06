// Code generated. DO NOT EDIT.

package apiv1

// The full API error.
type Status struct {
	// The general error code (e.g. `INVALID_ARGUMENT`).
	Code string `json:"code"`
	// A developer-facing error message.
	Message string `json:"message"`
	// A reason code for the error (e.g. `USER_PENDING_DELETION`).
	Reason string `json:"reason"`
	// The parameter path related to the error (e.g. `member.userId`).
	Param string `json:"param"`
	// Additional metadata related to the error.
	Metadata map[string]string `json:"metadata"`
}
