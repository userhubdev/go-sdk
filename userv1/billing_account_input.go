// Code generated. DO NOT EDIT.

package userv1

import (
	"github.com/userhubdev/go-sdk/commonv1"
	"github.com/userhubdev/go-sdk/types"
)

// BillingAccountInput input parameters.
type BillingAccountInput struct {
	// The human-readable display name of the billing account.
	//
	// The maximum length is 200 characters.
	//
	// This might be further restricted by the billing provider.
	DisplayName types.Optional[string] `json:"displayName"`
	// The email address of the billing account.
	//
	// The maximum length is 320 characters.
	//
	// This might be further restricted by the billing provider.
	Email types.Optional[string] `json:"email"`
	// The E164 phone number of the billing account (e.g. `+12125550123`).
	PhoneNumber types.Optional[string] `json:"phoneNumber"`
	// The address for the billing account.
	Address types.Optional[*commonv1.Address] `json:"address"`
}
