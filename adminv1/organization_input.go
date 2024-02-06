// Code generated. DO NOT EDIT.

package adminv1

import (
	"time"

	"github.com/userhubdev/go-sdk/commonv1"
)

// Organization input parameters.
type OrganizationInput struct {
	// The system-assigned identifier of the organization.
	Id *string `json:"id"`
	// The client defined unique identifier of the organization account.
	//
	// It is restricted to letters, numbers, underscores, and hyphens,
	// with the first character a letter or a number, and a 255
	// character maximum.
	//
	// ID's starting with `org_` are reserved.
	UniqueId *string `json:"uniqueId"`
	// The human-readable display name of the organization.
	//
	// The maximum length is 200 characters.
	DisplayName *string `json:"displayName"`
	// The email address of the organization.
	//
	// The maximum length is 320 characters.
	Email *string `json:"email"`
	// Whether the organization's email address has been verified.
	EmailVerified *bool `json:"emailVerified"`
	// The E164 phone number for the organization (e.g. `+12125550123`).
	PhoneNumber *string `json:"phoneNumber"`
	// Whether the organization's phone number has been verified.
	PhoneNumberVerified *bool `json:"phoneNumberVerified"`
	// The photo/avatar URL of the organization.
	//
	// The maximum length is 2000 characters.
	ImageUrl *string `json:"imageUrl"`
	// The default ISO-4217 currency code for the organization (e.g. `USD`).
	CurrencyCode *string `json:"currencyCode"`
	// The IETF BCP-47 language code for the organization (e.g. `en`).
	LanguageCode *string `json:"languageCode"`
	// The country/region code for the organization (e.g. `US`).
	RegionCode *string `json:"regionCode"`
	// The IANA time zone for the organization (e.g. `America/New_York`).
	TimeZone *string `json:"timeZone"`
	// The billing address for the organization.
	Address *commonv1.Address `json:"address"`
	// The sign-up time for the organization.
	SignupTime *time.Time `json:"signupTime"`
	// Whether the organization is disabled.
	Disabled *bool `json:"disabled"`
}
