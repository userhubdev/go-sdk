// Code generated. DO NOT EDIT.

package userv1

import (
	"github.com/userhubdev/go-sdk/commonv1"
)

// The checkout item.
type CheckoutItem struct {
	// The item identifier.
	Id string `json:"id"`
	// The display name for the item.
	DisplayName string `json:"displayName"`
	// The input type of the item.
	InputType string `json:"inputType"`
	// The type of the item.
	Type string `json:"type"`
	// The unit for the item.
	Unit string `json:"unit"`
	// The price for the item.
	Price *Price `json:"price"`
	// The quantity for the item.
	Quantity int32 `json:"quantity"`
	// The minimum quantity allowed.
	//
	// This will only be set when quantity is settable.
	MinQuantity int32 `json:"minQuantity"`
	// The maximum quantity allowed.
	//
	// This will only be set when the quantity is settable and there is a
	// discrete (non-infinity) maximum.
	MaxQuantity int32 `json:"maxQuantity"`
	// The quantity at which the plan will renew.
	//
	// This will only be set when different from quantity and the
	// subscription is set to renew.
	RenewQuantity int32 `json:"renewQuantity"`
	// The minimum renew quantity allowed.
	//
	// This will only be set when renew quantity is settable.
	MinRenewQuantity int32 `json:"minRenewQuantity"`
	// The maximum renew quantity allowed.
	//
	// This will only be set when the new quantity is settable and there is a
	// discrete (non-infinity) maximum.
	MaxRenewQuantity int32 `json:"maxRenewQuantity"`
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
	// The item identifier for which you can group this item.
	//
	// This allows you to group credits and other preview items
	// with the related plan, seat, or add-on item.
	GroupItemId string `json:"groupItemId"`
}
