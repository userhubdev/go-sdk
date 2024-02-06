// Code generated. DO NOT EDIT.

package eventsv1

import (
	"github.com/userhubdev/go-sdk/adminv1"
)

// The members changed event.
type MembersChanged struct {
	// The organization.
	Organization *adminv1.Organization `json:"organization"`
	// The member.
	Member *adminv1.Member `json:"member"`
}
