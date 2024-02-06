// Code generated. DO NOT EDIT.

package adminv1

import (
	"github.com/userhubdev/go-sdk/apiv1"
)

// Result wrapper for a trigger.
type TriggerResult struct {
	// The trigger.
	Trigger *Trigger `json:"trigger"`
	// The trigger error.
	Error *apiv1.Status `json:"error"`
}
