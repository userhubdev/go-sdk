// Code generated. DO NOT EDIT.

package userv1

// A card payment method (e.g. credit, debit, etc...).
type CardPaymentMethod struct {
	// The brand of the card (e.g. `VISA`).
	Brand string `json:"brand"`
	// The expiration year.
	ExpYear int32 `json:"expYear"`
	// The expiration month.
	ExpMonth int32 `json:"expMonth"`
	// The last for digits of the card.
	Last4 string `json:"last4"`
	// The funding method for the card (e.g. `DEBIT`)
	FundingType string `json:"fundingType"`
}
