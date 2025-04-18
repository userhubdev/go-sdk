// Code generated. DO NOT EDIT.

package adminv1

import (
	"time"

	"github.com/userhubdev/go-sdk/commonv1"
)

// Individual account.
type User struct {
	// The system-assigned identifier of the user.
	Id string `json:"id"`
	// The current state of the user.
	State string `json:"state"`
	// The code that best describes the reason for the state.
	StateReason string `json:"stateReason"`
	// The client defined unique identifier of the user account.
	UniqueId string `json:"uniqueId"`
	// The human-readable display name of the user.
	DisplayName string `json:"displayName"`
	// The email address of the user.
	Email string `json:"email"`
	// Whether the user's email address has been verified.
	EmailVerified bool `json:"emailVerified"`
	// The E164 phone number for the user (e.g. `+12125550123`).
	PhoneNumber string `json:"phoneNumber"`
	// Whether the user's phone number has been verified.
	PhoneNumberVerified bool `json:"phoneNumberVerified"`
	// The photo/avatar URL of the user.
	ImageUrl string `json:"imageUrl"`
	// The default ISO-4217 currency code for the user (e.g. `USD`).
	CurrencyCode string `json:"currencyCode"`
	// The IETF BCP-47 language code for the user (e.g. `en`).
	LanguageCode string `json:"languageCode"`
	// The country/region code for the user (e.g. `US`).
	RegionCode string `json:"regionCode"`
	// The IANA time zone for the user (e.g. `America/New_York`).
	TimeZone string `json:"timeZone"`
	// The default address for the user.
	Address *commonv1.Address `json:"address"`
	// The connected accounts.
	AccountConnections []*AccountConnection `json:"accountConnections"`
	// The user's default active individual subscription.
	//
	// See memberships for organization subscription and
	// seat information.
	Subscription *AccountSubscription `json:"subscription"`
	// The user's organization memberships.
	Memberships []*Membership `json:"memberships"`
	// The sign-up time for the user.
	SignupTime time.Time `json:"signupTime"`
	// Whether the user is disabled.
	Disabled bool `json:"disabled"`
	// The user view.
	View string `json:"view"`
	// The creation time of the user.
	CreateTime time.Time `json:"createTime"`
	// The last update time of the user.
	UpdateTime time.Time `json:"updateTime"`
}
