// Code generated. DO NOT EDIT.

package adminv1

import (
	"time"

	"github.com/userhubdev/go-sdk/commonv1"
)

// Price for a product.
type Price struct {
	// The system-assigned identifier of the price.
	Id string `json:"id"`
	// The connection.
	Connection *Connection `json:"connection"`
	// The external identifier of the connected price.
	ExternalId string `json:"externalId"`
	// The status of the connected price.
	State string `json:"state"`
	// The code that best describes the reason for the state.
	StateReason string `json:"stateReason"`
	// The currency for the price.
	CurrencyCode string `json:"currencyCode"`
	// The billing mode for the price.
	BillingMode string `json:"billingMode"`
	// The billing interval for the price.
	Interval *commonv1.Interval `json:"interval"`
	// The admin-facing display name of the price.
	DisplayName string `json:"displayName"`
	// The product associated with the price.
	Product *Product `json:"product"`
	// The price is dependent on the quantity.
	Empty *PriceEmptyPrice `json:"empty"`
	// The price is fixed per quantity.
	Fixed *PriceFixedPrice `json:"fixed"`
	// The price is dependent on the quantity.
	Tiered *PriceTieredPrice `json:"tiered"`
	// The archived status of the price.
	//
	// It determines if the price can be used.
	Archived bool `json:"archived"`
	// The price view.
	View string `json:"view"`
	// The creation time of the price.
	CreateTime time.Time `json:"createTime"`
	// The last update time of the price.
	UpdateTime time.Time `json:"updateTime"`
}
