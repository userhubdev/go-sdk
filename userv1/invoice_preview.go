// Code generated. DO NOT EDIT.

package userv1

import (
	"time"
)

// A preview of an invoice.
type InvoicePreview struct {
	// The currency code for the invoice (e.g. `USD`).
	CurrencyCode string `json:"currencyCode"`
	// The contact information associated with the invoice.
	Account *InvoiceAccount `json:"account"`
	// The time the upcoming invoice will be finalized.
	//
	// This is an estimate and might not exactly match the finalized
	// invoice.
	//
	// This will be null if the preview would be applied
	// immediately.
	EffectiveTime time.Time `json:"effectiveTime"`
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
	// A token which can be passed to a create or update
	// operation to ensure the change matches the preview.
	//
	// This token generally expires within 10 minutes of
	// being generated.
	ChangeToken string `json:"changeToken"`
	// The line items for the invoice.
	Items []*InvoicePreviewItem `json:"items"`
}
