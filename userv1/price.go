// Code generated. DO NOT EDIT.

package userv1

import (
	"github.com/userhubdev/go-sdk/commonv1"
)

// Price for a product.
type Price struct {
	// The system-assigned identifier of the price.
	Id string `json:"id"`
	// The currency for the price.
	CurrencyCode string `json:"currencyCode"`
	// The billing mode for the price.
	BillingMode string `json:"billingMode"`
	// The billing interval for the price.
	Interval *commonv1.Interval `json:"interval"`
	// The price is fixed per quantity.
	Fixed *PriceFixedPrice `json:"fixed"`
	// The price is dependent on the quantity.
	Tiered *PriceTieredPrice `json:"tiered"`
}
