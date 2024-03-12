// Code generated. DO NOT EDIT.

package adminv1

import (
	"github.com/userhubdev/go-sdk/commonv1"
)

// Payment method input parameters.
type PaymentMethodInput struct {
	// The full name of the owner of the payment method (e.g. `Jane Doe`).
	FullName *string `json:"fullName"`
	// The address for the payment method.
	Address *commonv1.Address `json:"address"`
	// The card expiration year (e.g. `2030`).
	ExpYear *int32 `json:"expYear"`
	// The card expiration month (e.g. `12`).
	ExpMonth *int32 `json:"expMonth"`
}
