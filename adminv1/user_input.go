// Code generated. DO NOT EDIT.

package adminv1

import (
	"time"

	"github.com/userhubdev/go-sdk/commonv1"
)

// User input parameters.
type UserInput struct {
	// The system-assigned identifier of the user.
	Id string `json:"id"`
	// The client defined unique identifier of the user account.
	//
	// It is restricted to letters, numbers, underscores, and hyphens,
	// with the first character a letter or a number, and a 255
	// character maximum.
	//
	// ID's starting with `usr_` are reserved.
	UniqueId *string `json:"uniqueId"`
	// The human-readable display name of the user.
	//
	// The maximum length is 200 characters.
	DisplayName *string `json:"displayName"`
	// The email address of the user.
	//
	// The maximum length is 320 characters.
	Email *string `json:"email"`
	// Whether the user's email address has been verified.
	EmailVerified *bool `json:"emailVerified"`
	// The E164 phone number for the user (e.g. `+12125550123`).
	PhoneNumber *string `json:"phoneNumber"`
	// Whether the user's phone number has been verified.
	PhoneNumberVerified *bool `json:"phoneNumberVerified"`
	// The photo/avatar URL of the user.
	//
	// The maximum length is 2000 characters.
	ImageUrl *string `json:"imageUrl"`
	// The default ISO-4217 currency code for the user (e.g. `USD`).
	CurrencyCode *string `json:"currencyCode"`
	// The IETF BCP-47 language code for the user (e.g. `en`).
	LanguageCode *string `json:"languageCode"`
	// The country/region code for the user (e.g. `US`).
	RegionCode *string `json:"regionCode"`
	// The IANA time zone for the user (e.g. `America/New_York`).
	TimeZone *string `json:"timeZone"`
	// The billing address for the user.
	Address *commonv1.Address `json:"address"`
	// The sign-up time for the user.
	SignupTime *time.Time `json:"signupTime"`
	// Whether the user is disabled.
	Disabled *bool `json:"disabled"`
}
