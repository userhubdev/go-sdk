// Code generated. DO NOT EDIT.

package userv1

// The billing account for an organization or user.
type BillingAccount struct {
	// The status of the billing account.
	State string `json:"state"`
	// The balance amount for the account.
	//
	// A negative value indicates an amount which will be subtracted from the next
	// invoice (credit).
	//
	// A positive value indicates an amount which will be added to the next
	// invoice (debt).
	BalanceAmount string `json:"balanceAmount"`
	// The ISO-4217 currency code for the account (e.g. `USD`).
	CurrencyCode string `json:"currencyCode"`
	// The default and latest 10 payment methods for the account.
	PaymentMethods []*PaymentMethod `json:"paymentMethods"`
	// The subscription for the account.
	Subscription *Subscription `json:"subscription"`
}
