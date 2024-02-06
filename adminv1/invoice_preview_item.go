// Code generated. DO NOT EDIT.

package adminv1

import (
	"github.com/userhubdev/go-sdk/commonv1"
)

// The line items for the preview.
type InvoicePreviewItem struct {
	// The details of the associated product.
	Product *Product `json:"product"`
	// The details of the associated price.
	Price *Price `json:"price"`
	// The quantity of the item product/price.
	Quantity int32 `json:"quantity"`
	// The total amount for the line item excluding
	// taxes and discounts.
	SubtotalAmount string `json:"subtotalAmount"`
	// The item-level discount amount.
	DiscountAmount string `json:"discountAmount"`
	// The user facing description for the line item.
	Description string `json:"description"`
	// Whether the item was a proration.
	Proration bool `json:"proration"`
	// The billing period for the item.
	Period *commonv1.Period `json:"period"`
}
