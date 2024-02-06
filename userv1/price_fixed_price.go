// Code generated. DO NOT EDIT.

package userv1

// A pricing strategy that defines a constant price per
// quantity.
type PriceFixedPrice struct {
	// The decimal amount for the defined currency (e.g. `9.95`).
	Amount string `json:"amount"`
	// Whether to transform the quantity before multiplying amount.
	TransformQuantity *PriceTransformQuantity `json:"transformQuantity"`
}
