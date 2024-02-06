// Code generated. DO NOT EDIT.

package userv1

import (
	"time"
)

// The user's or organization's subscription.
type Subscription struct {
	// The system-assigned identifier of the subscription.
	Id string `json:"id"`
	// The status of the subscription.
	State string `json:"state"`
	// The currency code for the subscription (e.g. `USD`).
	CurrencyCode string `json:"currencyCode"`
	// The subscription items.
	Plan *Plan `json:"plan"`
	// The payment method.
	PaymentMethod *PaymentMethod `json:"paymentMethod"`
	// The subscription is scheduled to be canceled at the end of the
	// current billing period.
	CancelPeriodEnd bool `json:"cancelPeriodEnd"`
	// The time the subscription started.
	StartTime time.Time `json:"startTime"`
	// The time the subscription ends/ended.
	EndTime time.Time `json:"endTime"`
	// The trial information for the subscription.
	Trial *SubscriptionTrial `json:"trial"`
	// The current billing period for the subscription.
	CurrentPeriod *SubscriptionCurrentPeriod `json:"currentPeriod"`
	// The subscription items.
	Items []*SubscriptionItem `json:"items"`
	// The information about the various seats available in
	// the subscription.
	Seats []*SubscriptionSeatInfo `json:"seats"`
	// The creation time of the subscription.
	CreateTime time.Time `json:"createTime"`
	// The last update time of the subscription.
	UpdateTime time.Time `json:"updateTime"`
}
