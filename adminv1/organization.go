// Code generated. DO NOT EDIT.

package adminv1

import (
	"time"

	"github.com/userhubdev/go-sdk/commonv1"
)

// A group account.
type Organization struct {
	// The system-assigned identifier of the organization.
	Id string `json:"id"`
	// The current state of the organization.
	State string `json:"state"`
	// The code that best describes the reason for the state.
	StateReason string `json:"stateReason"`
	// The client defined unique identifier of the organization account.
	UniqueId string `json:"uniqueId"`
	// The human-readable display name of the organization.
	DisplayName string `json:"displayName"`
	// The email address of the organization.
	Email string `json:"email"`
	// Whether the organization's email address has been verified.
	EmailVerified bool `json:"emailVerified"`
	// The E164 phone number for the organization (e.g. `+12125550123`).
	PhoneNumber string `json:"phoneNumber"`
	// Whether the organization's phone number has been verified.
	PhoneNumberVerified bool `json:"phoneNumberVerified"`
	// The photo/avatar URL of the organization.
	ImageUrl string `json:"imageUrl"`
	// The default ISO-4217 currency code for the organization (e.g. `USD`).
	CurrencyCode string `json:"currencyCode"`
	// The IETF BCP-47 language code for the organization (e.g. `en`).
	LanguageCode string `json:"languageCode"`
	// The country/region code for the organization (e.g. `US`).
	RegionCode string `json:"regionCode"`
	// The IANA time zone for the organization (e.g. `America/New_York`).
	TimeZone string `json:"timeZone"`
	// The address for the organization.
	Address *commonv1.Address `json:"address"`
	// The connected accounts.
	AccountConnections []*AccountConnection `json:"accountConnections"`
	// The organization's default active subscription.
	Subscription *AccountSubscription `json:"subscription"`
	// The sign-up time for the organization.
	SignupTime time.Time `json:"signupTime"`
	// Whether the organization is disabled.
	Disabled bool `json:"disabled"`
	// The creation time of the organization.
	CreateTime time.Time `json:"createTime"`
	// The last update time of the organization.
	UpdateTime time.Time `json:"updateTime"`
}
