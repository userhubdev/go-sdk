// Code generated. DO NOT EDIT.

package userv1

import (
	"time"

	"github.com/userhubdev/go-sdk/commonv1"
)

// A bill or statement.
type Invoice struct {
	// The system-assigned identifier of the invoice.
	Id string `json:"id"`
	// The status of the invoice.
	State string `json:"state"`
	// The time associated with the current state (e.g. paid time).
	StateTime time.Time `json:"stateTime"`
	// The invoice number.
	Number string `json:"number"`
	// The currency code for the invoice (e.g. `USD`).
	CurrencyCode string `json:"currencyCode"`
	// The user facing description for the invoice.
	Description string `json:"description"`
	// The contact information associated with the invoice.
	Account *InvoiceAccount `json:"account"`
	// The time the invoice was finalized.
	EffectiveTime time.Time `json:"effectiveTime"`
	// The billing period for the invoice.
	Period *commonv1.Period `json:"period"`
	// The subtotal amount for the invoice.
	//
	// This includes item-level discounts.
	SubtotalAmount string `json:"subtotalAmount"`
	// The invoice-level discount amount.
	//
	// This does not include item level discounts.
	DiscountAmount string `json:"discountAmount"`
	// The starting and ending account balance as
	// of the time the invoice was finalized.
	Balance *InvoiceBalance `json:"balance"`
	// The tax amount for the invoice.
	//
	// This is for rendering purposes only and is
	// not the reported tax amount.
	TaxAmount string `json:"taxAmount"`
	// The total amount for the invoice.
	TotalAmount string `json:"totalAmount"`
	// The total amount minus any credits automatically
	// associated with the invoice.
	DueAmount string `json:"dueAmount"`
	// The due amount minus the amount already paid.
	RemainingDueAmount string `json:"remainingDueAmount"`
	// The time the invoice must be paid by.
	DueTime time.Time `json:"dueTime"`
	// The amount already paid towards the invoice.
	PaidAmount string `json:"paidAmount"`
	// The status of the invoice payment.
	PaymentState string `json:"paymentState"`
	// The payment intent for the invoice.
	PaymentIntent *PaymentIntent `json:"paymentIntent"`
	// The line items for the invoice.
	Items []*InvoiceItem `json:"items"`
	// The creation time of the invoice.
	CreateTime time.Time `json:"createTime"`
	// The last update time of the invoice.
	UpdateTime time.Time `json:"updateTime"`
}
