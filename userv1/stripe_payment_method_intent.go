// Code generated. DO NOT EDIT.

package userv1

// A Stripe Setup Intent.
type StripePaymentMethodIntent struct {
	// The Stripe account ID (e.g. `acct_1LcUvxQYGbxD2SPK`)
	AccountId string `json:"accountId"`
	// Whether the Stripe Setup Intent was created in live mode.
	Live bool `json:"live"`
	// The Stripe Setup Intent client secret.
	ClientSecret string `json:"clientSecret"`
	// The Stripe.js Payment Element options.
	PaymentElementOptions map[string]any `json:"paymentElementOptions"`
}
