// Code generated. DO NOT EDIT.

package adminv1

import (
	"github.com/userhubdev/go-sdk/apiv1"
)

// The user or error.
type UserResult struct {
	// The user.
	User *User `json:"user"`
	// The result error.
	Error *apiv1.Status `json:"error"`
}
