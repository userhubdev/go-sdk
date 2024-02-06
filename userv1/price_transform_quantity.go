// Code generated. DO NOT EDIT.

package userv1

// A quantity transform for fixed prices.
type PriceTransformQuantity struct {
	// The amount to divide the quantity by.
	Divisor int32 `json:"divisor"`
	// Whether to round the result up or down.
	Round string `json:"round"`
}
