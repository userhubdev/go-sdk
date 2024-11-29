// Code generated. DO NOT EDIT.

package userv1

// A quantity range within the tiered price.
type PriceTieredPriceTier struct {
	// The upper quantity for tier (inclusive).
	Upper int32 `json:"upper"`
	// The per quantity amount for the tier.
	UnitAmount string `json:"unitAmount"`
	// The flat amount for the tier.
	FlatAmount string `json:"flatAmount"`
}
