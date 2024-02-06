// Code generated. DO NOT EDIT.

package userv1

// A pricing strategy that dynamically sets the price for a given
// quantity range.
type PriceTieredPrice struct {
	// The strategy for evaluating the tiers.
	Mode string `json:"mode"`
	// The tiers for the price.
	Tiers []*TieredPriceTier `json:"tiers"`
}
