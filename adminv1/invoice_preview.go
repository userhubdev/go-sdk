// Code generated. DO NOT EDIT.

package adminv1

import (
	"time"
)

// A preview of an invoice.
type InvoicePreview struct {
	// The currency code for the preview (e.g. `USD`).
	CurrencyCode string `json:"currencyCode"`
	// The bill to contact information.
	Account *InvoiceAccount `json:"account"`
	// The time the upcoming invoice will be finalized.
	//
	// This is an estimate and might not exactly match the finalized
	// invoice.
	//
	// This will be null if the preview would be applied
	// immediately.
	EffectiveTime time.Time `json:"effectiveTime"`
	// The subtotal amount for the preview.
	//
	// This includes item-level discounts.
	SubtotalAmount string `json:"subtotalAmount"`
	// The preview-level discount amount.
	//
	// This does not include item level discounts.
	DiscountAmount string `json:"discountAmount"`
	// The starting and ending account balance as
	// of the time the preview.
	Balance *InvoiceBalance `json:"balance"`
	// The tax amount for the preview.
	TaxAmount string `json:"taxAmount"`
	// The total amount for the preview.
	TotalAmount string `json:"totalAmount"`
	// The total amount minus any credits automatically
	// associated with the preview.
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
