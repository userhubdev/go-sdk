// Code generated. DO NOT EDIT.

package adminv1

import (
	"time"

	"github.com/userhubdev/go-sdk/commonv1"
)

// A link between an organization/user and an external account.
type AccountConnection struct {
	// The tenant connection.
	Connection *Connection `json:"connection"`
	// The external identifier of the connected account.
	ExternalId string `json:"externalId"`
	// The external admin URL for the connected account.
	AdminUrl string `json:"adminUrl"`
	// The status of the connected account.
	State string `json:"state"`
	// The code that best describes the reason for the state.
	StateReason string `json:"stateReason"`
	// The human-readable display name of the external account.
	DisplayName string `json:"displayName"`
	// The email address of the external account.
	Email string `json:"email"`
	// Whether the external account's email address has been verified.
	EmailVerified bool `json:"emailVerified"`
	// The E164 phone number for the external account (e.g. `+12125550123`).
	PhoneNumber string `json:"phoneNumber"`
	// Whether the external account's phone number has been verified.
	PhoneNumberVerified bool `json:"phoneNumberVerified"`
	// The billing address for the external account.
	Address *commonv1.Address `json:"address"`
	// The currency code for the account.
	CurrencyCode string `json:"currencyCode"`
	// The balance amount for the account.
	//
	// A negative value indicates an amount which will be subtracted from the next
	// invoice (credit).
	//
	// A positive value indicates an amount which will be added to the next
	// invoice (debt).
	BalanceAmount string `json:"balanceAmount"`
	// The payment methods for connections that support it.
	PaymentMethods []*PaymentMethod `json:"paymentMethods"`
	// The creation time of the account connection.
	CreateTime time.Time `json:"createTime"`
	// The last update time of the account connection.
	UpdateTime time.Time `json:"updateTime"`
}
