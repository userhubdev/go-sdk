// Code generated. DO NOT EDIT.

package userv1

import (
	"github.com/userhubdev/go-sdk/commonv1"
)

// The billing account for an organization or user.
type BillingAccount struct {
	// The status of the billing account.
	State string `json:"state"`
	// The human-readable display name of the billing account.
	DisplayName string `json:"displayName"`
	// The email address of the billing account.
	Email string `json:"email"`
	// The E164 phone number for the billing account (e.g. `+12125550123`).
	PhoneNumber string `json:"phoneNumber"`
	// The billing address for the billing account.
	Address *commonv1.Address `json:"address"`
	// The ISO-4217 currency code for the billing account (e.g. `USD`).
	CurrencyCode string `json:"currencyCode"`
	// The balance amount for the account.
	//
	// A negative value indicates an amount which will be subtracted from the next
	// invoice (credit).
	//
	// A positive value indicates an amount which will be added to the next
	// invoice (debt).
	BalanceAmount string `json:"balanceAmount"`
	// The available checkouts.
	Checkouts []*BillingAccountCheckout `json:"checkouts"`
	// The default and latest 10 payment methods for the account.
	PaymentMethods []*PaymentMethod `json:"paymentMethods"`
	// The subscription for the account.
	Subscription *Subscription `json:"subscription"`
}
