// Code generated. DO NOT EDIT.

package adminv1

import (
	"time"

	"github.com/userhubdev/go-sdk/commonv1"
	"github.com/userhubdev/go-sdk/types"
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
	UniqueId types.Optional[string] `json:"uniqueId"`
	// The human-readable display name of the user.
	//
	// The maximum length is 200 characters.
	DisplayName types.Optional[string] `json:"displayName"`
	// The email address of the user.
	//
	// The maximum length is 320 characters.
	Email types.Optional[string] `json:"email"`
	// Whether the user's email address has been verified.
	EmailVerified types.Optional[bool] `json:"emailVerified"`
	// The E164 phone number for the user (e.g. `+12125550123`).
	PhoneNumber types.Optional[string] `json:"phoneNumber"`
	// Whether the user's phone number has been verified.
	PhoneNumberVerified types.Optional[bool] `json:"phoneNumberVerified"`
	// The photo/avatar URL of the user.
	//
	// The maximum length is 2000 characters.
	ImageUrl types.Optional[string] `json:"imageUrl"`
	// The default ISO-4217 currency code for the user (e.g. `USD`).
	CurrencyCode types.Optional[string] `json:"currencyCode"`
	// The IETF BCP-47 language code for the user (e.g. `en`).
	LanguageCode types.Optional[string] `json:"languageCode"`
	// The country/region code for the user (e.g. `US`).
	RegionCode types.Optional[string] `json:"regionCode"`
	// The IANA time zone for the user (e.g. `America/New_York`).
	TimeZone types.Optional[string] `json:"timeZone"`
	// The default address for the user.
	Address types.Optional[*commonv1.Address] `json:"address"`
	// The sign-up time for the user.
	SignupTime types.Optional[time.Time] `json:"signupTime"`
	// Whether the user is disabled.
	Disabled types.Optional[bool] `json:"disabled"`
}
