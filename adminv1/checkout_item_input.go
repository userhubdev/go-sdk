// Code generated. DO NOT EDIT.

package adminv1

import "github.com/userhubdev/go-sdk/types"

// Checkout item input.
type CheckoutItemInput struct {
	// The identifier of the item.
	Id string `json:"id"`
	// The quantity for the item.
	Quantity types.Optional[int32] `json:"quantity"`
}
