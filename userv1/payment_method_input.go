// Code generated. DO NOT EDIT.

package userv1

import (
	"github.com/userhubdev/go-sdk/commonv1"
	"github.com/userhubdev/go-sdk/types"
)

// Payment method input parameters.
type PaymentMethodInput struct {
	// The full name of the owner of the payment method (e.g. `Jane Doe`).
	FullName types.Optional[string] `json:"fullName"`
	// The address for the payment method.
	Address types.Optional[*commonv1.Address] `json:"address"`
	// The card expiration year (e.g. `2030`).
	ExpYear types.Optional[int32] `json:"expYear"`
	// The card expiration month (e.g. `12`).
	ExpMonth types.Optional[int32] `json:"expMonth"`
}
