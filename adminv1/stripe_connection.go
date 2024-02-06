// Code generated. DO NOT EDIT.

package adminv1

// The stripe billing specific connection data.
type StripeConnection struct {
	// The Stripe account ID (e.g. `acct_1LcUvxQYGbxD2SPK`)
	AccountId string `json:"accountId"`
	// The live vs test status for the Stripe account.
	Live bool `json:"live"`
}
