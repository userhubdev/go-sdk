// Code generated. DO NOT EDIT.

package adminv1

import (
	"github.com/userhubdev/go-sdk/commonv1"
	"github.com/userhubdev/go-sdk/types"
)

// AccountConnection input parameters.
type AccountConnectionInput struct {
	// The system-assigned identifier for the connection of the external account.
	ConnectionId string `json:"connectionId"`
	// The human-readable display name of the external account.
	//
	// The maximum length is 200 characters.
	//
	// This might be further restricted by the external provider.
	DisplayName types.Optional[string] `json:"displayName"`
	// The email address of the external account.
	//
	// The maximum length is 320 characters.
	//
	// This might be further restricted by the external provider.
	Email types.Optional[string] `json:"email"`
	// Whether the external account's email address has been verified.
	EmailVerified types.Optional[bool] `json:"emailVerified"`
	// The E164 phone number for the external account (e.g. `+12125550123`).
	PhoneNumber types.Optional[string] `json:"phoneNumber"`
	// Whether the external account's phone number has been verified.
	PhoneNumberVerified types.Optional[bool] `json:"phoneNumberVerified"`
	// The default ISO-4217 currency code for the external account (e.g. `USD`).
	CurrencyCode types.Optional[string] `json:"currencyCode"`
	// The billing address for the external account.
	Address types.Optional[*commonv1.Address] `json:"address"`
	// Whether the external account is disabled.
	Disabled types.Optional[bool] `json:"disabled"`
}
