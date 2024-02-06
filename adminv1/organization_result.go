// Code generated. DO NOT EDIT.

package adminv1

import (
	"github.com/userhubdev/go-sdk/apiv1"
)

// Result wrapper for an organization.
type OrganizationResult struct {
	// The organization.
	Organization *Organization `json:"organization"`
	// The organization error.
	Error *apiv1.Status `json:"error"`
}
