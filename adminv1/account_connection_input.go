// Code generated. DO NOT EDIT.

package adminv1

import (
	"github.com/userhubdev/go-sdk/commonv1"
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
	DisplayName *string `json:"displayName"`
	// The email address of the external account.
	//
	// The maximum length is 320 characters.
	//
	// This might be further restricted by the external provider.
	Email *string `json:"email"`
	// Whether the external account's email address has been verified.
	EmailVerified *bool `json:"emailVerified"`
	// The E164 phone number for the external account (e.g. `+12125550123`).
	PhoneNumber *string `json:"phoneNumber"`
	// Whether the external account's phone number has been verified.
	PhoneNumberVerified *bool `json:"phoneNumberVerified"`
	// The default ISO-4217 currency code for the external account (e.g. `USD`).
	CurrencyCode *string `json:"currencyCode"`
	// The billing address for the external account.
	Address *commonv1.Address `json:"address"`
	// Whether the external account is disabled.
	Disabled *bool `json:"disabled"`
}
