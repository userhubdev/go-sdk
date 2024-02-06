// Code generated. DO NOT EDIT.

package userv1

import (
	"github.com/userhubdev/go-sdk/commonv1"
)

// THe contact information for the invoice.
type InvoiceAccount struct {
	// The company or individual's full name.
	FullName string `json:"fullName"`
	// The billing email address.
	Email string `json:"email"`
	// The billing phone number.
	PhoneNumber string `json:"phoneNumber"`
	// The billing address.
	Address *commonv1.Address `json:"address"`
}
