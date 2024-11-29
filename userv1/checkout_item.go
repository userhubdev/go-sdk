// Code generated. DO NOT EDIT.

package userv1

import (
	"github.com/userhubdev/go-sdk/commonv1"
)

// The checkout item.
type CheckoutItem struct {
	// The item identifier.
	Id string `json:"id"`
	// The type of the item.
	Type string `json:"type"`
	// The display name for the item.
	DisplayName string `json:"displayName"`
	// The quantity for the item.
	Price *Price `json:"price"`
	// The quantity for the item.
	Quantity int32 `json:"quantity"`
	// The quantity the plan is set to renew at.
	RenewQuantity int32 `json:"renewQuantity"`
	// The minimum quantity allowed.
	MinQuantity int32 `json:"minQuantity"`
	// The maximum quantity allowed.
	MaxQuantity int32 `json:"maxQuantity"`
	// The billing period for the item.
	Period *commonv1.Period `json:"period"`
	// The subtotal amount at checkout.
	SubtotalAmount string `json:"subtotalAmount"`
	// The item-level discount amount at checkout.
	DiscountAmount string `json:"discountAmount"`
	// The item-level normal recurring amount.
	RenewAmount string `json:"renewAmount"`
	// Whether this is a preview-only item.
	//
	// Preview-only items are generally prorations or other pending
	// charges or credits.
	Preview bool `json:"preview"`
	// The parent item.
	//
	// This allows you to group related items and is generally set for preview
	// items.
	ParentItemId string `json:"parentItemId"`
}
