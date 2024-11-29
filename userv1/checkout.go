// Code generated. DO NOT EDIT.

package userv1

import (
	"github.com/userhubdev/go-sdk/apiv1"
	"github.com/userhubdev/go-sdk/commonv1"
)

// The checkout.
type Checkout struct {
	// The system-assigned identifier of the checkout.
	//
	// This might not match the requested checkout identifier and you
	// should update all references to it if it changes.
	//
	// Finalized identifiers will start with `co_`, all other identifiers
	// are temporary.
	Id string `json:"id"`
	// The type of checkout.
	Type string `json:"type"`
	// The state of the checkout.
	State string `json:"state"`
	// The checkout error.
	Error *apiv1.Status `json:"error"`
	// The currently selected currency code.
	CurrencyCode string `json:"currencyCode"`
	// The plans available for checkout.
	Plans []*CheckoutPlan `json:"plans"`
	// The checkout payment method.
	PaymentMethod *PaymentMethod `json:"paymentMethod"`
	// The billing details company or individual name.
	FullName string `json:"fullName"`
	// The billing details address.
	Address *commonv1.Address `json:"address"`
	// The steps required to complete the checkout.
	Steps []*CheckoutStep `json:"steps"`
	// The checkout items.
	Items []*CheckoutItem `json:"items"`
	// The discounts applied to the checkout.
	Discounts []*CheckoutDiscount `json:"discounts"`
	// The subtotal amount for the checkout.
	//
	// This includes item-level discounts.
	SubtotalAmount string `json:"subtotalAmount"`
	// The top-level discount amount.
	//
	// This does not include item level discounts.
	DiscountAmount string `json:"discountAmount"`
	// The tax amount for the checkout.
	//
	// This is for rendering purposes only and is
	// not the reported tax amount.
	TaxAmount string `json:"taxAmount"`
	// The total amount for the checkout.
	TotalAmount string `json:"totalAmount"`
	// The amount applied to the checkout from the balance.
	//
	// A negative amount means a debit from the account balance.
	// A positive amount means a credit to the account balance.
	BalanceAppliedAmount string `json:"balanceAppliedAmount"`
	// The total amount minus any credits automatically
	// associated with the invoice.
	DueAmount string `json:"dueAmount"`
	// The total normal recurring amount.
	RenewAmount string `json:"renewAmount"`
}
