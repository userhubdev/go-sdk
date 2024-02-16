// Code generated. DO NOT EDIT.

package adminv1

import (
	"time"
)

// A prorated change that occurred mid-billing cycle.
type InvoiceChange struct {
	// The time the change occurred.
	Time time.Time `json:"time"`
	// The user-facing description for the change.
	Description string `json:"description"`
	// The total amount for the change excluding
	// taxes and discounts.
	SubtotalAmount string `json:"subtotalAmount"`
	// The change discount amount.
	DiscountAmount string `json:"discountAmount"`
	// The starting quantity of the change.
	StartQuantity int32 `json:"startQuantity"`
	// The ending quantity of the change.
	EndQuantity int32 `json:"endQuantity"`
	// The starting (credited) item identifiers.
	StartItemIds []string `json:"startItemIds"`
	// The ending (charged) item identifiers.
	EndItemIds []string `json:"endItemIds"`
}
