// Code generated. DO NOT EDIT.

package adminv1

import (
	"time"

	"github.com/userhubdev/go-sdk/commonv1"
)

// A link to an external payment method (e.g. credit card, bank account).
type PaymentMethod struct {
	// The system-assigned identifier of the payment method.
	Id string `json:"id"`
	// The external identifier of the connected payment method.
	ExternalId string `json:"externalId"`
	// The status of the connected payment method.
	State string `json:"state"`
	// The code that best describes the reason for the state.
	StateReason string `json:"stateReason"`
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
	// Whether the payment method is the default for the connected account.
	Default bool `json:"default"`
	// The last time the payment method was pulled from the connection.
	PullTime time.Time `json:"pullTime"`
	// The creation time of the payment method connection.
	CreateTime time.Time `json:"createTime"`
	// The last update time of the payment method connection.
	UpdateTime time.Time `json:"updateTime"`
	// Card payment method (e.g. Visa credit card).
	Card *CardPaymentMethod `json:"card"`
}
