// Code generated. DO NOT EDIT.

package adminv1

// A Stripe payment intent.
type StripePaymentIntent struct {
	// The Stripe account ID (e.g. `acct_1LcUvxQYGbxD2SPK`)
	AccountId string `json:"accountId"`
	// Whether the Stripe payment intent was created in live mode.
	Live bool `json:"live"`
	// The Stripe payment intent client secret.
	ClientSecret string `json:"clientSecret"`
}
