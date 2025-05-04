// Code generated. DO NOT EDIT.

package userv1

import (
	"time"

	"github.com/userhubdev/go-sdk/apiv1"
	"github.com/userhubdev/go-sdk/commonv1"
)

// A link to an external payment method (e.g. credit card, bank account).
type PaymentMethod struct {
	// The system-assigned identifier of the payment method.
	Id string `json:"id"`
	// The payment method type.
	Type string `json:"type"`
	// A human-readable description of the payment method.
	//
	// This can be used to show a description of the payment method
	// when the type is UNKNOWN or not explicitly handled.
	DisplayName string `json:"displayName"`
	// The full name of the owner of the payment method.
	FullName string `json:"fullName"`
	// The address for the payment method.
	Address *commonv1.Address `json:"address"`
	// Whether the payment method is the default for the account.
	Default bool `json:"default"`
	// The last payment error.
	//
	// This will be unset if the payment method is updated
	// or if a payment succeeds.
	LastPaymentError *apiv1.Status `json:"lastPaymentError"`
	// Card payment method (e.g. Visa credit card).
	Card *CardPaymentMethod `json:"card"`
	// The creation time of the payment method connection.
	CreateTime time.Time `json:"createTime"`
	// The last update time of the payment method connection.
	UpdateTime time.Time `json:"updateTime"`
}
