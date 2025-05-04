// Code generated. DO NOT EDIT.

package adminv1

import (
	"time"
)

// The representation of an activated plan.
type Subscription struct {
	// The system-assigned identifier of the subscription.
	Id string `json:"id"`
	// The status of the subscription.
	State string `json:"state"`
	// The code that best describes the reason for the state.
	StateReason string `json:"stateReason"`
	// The billing connection.
	Connection *Connection `json:"connection"`
	// The external identifier of the subscription.
	ExternalId string `json:"externalId"`
	// The plan.
	Plan *Plan `json:"plan"`
	// The currency code for the subscription (e.g. `USD`).
	CurrencyCode string `json:"currencyCode"`
	// The subscription items.
	Items []*SubscriptionItem `json:"items"`
	// The seat information.
	Seats []*SubscriptionSeatInfo `json:"seats"`
	// The payment method.
	PaymentMethod *PaymentMethod `json:"paymentMethod"`
	// Whether the subscription is scheduled to be canceled
	// at the end of the current billing period.
	RenewCanceled bool `json:"renewCanceled"`
	// Whether the subscription is scheduled to be canceled
	// at the end of the current billing period.
	//
	// Deprecated: Use `RenewCanceled` instead.
	CancelPeriodEnd bool `json:"cancelPeriodEnd"`
	// The anchor time for the billing cycle.
	AnchorTime time.Time `json:"anchorTime"`
	// The time the subscription started.
	StartTime time.Time `json:"startTime"`
	// The time the subscription ends/ended.
	EndTime time.Time `json:"endTime"`
	// The trial information for the subscription.
	Trial *SubscriptionTrial `json:"trial"`
	// The current billing period for the subscription.
	CurrentPeriod *SubscriptionCurrentPeriod `json:"currentPeriod"`
	// The organization owner of the subscription.
	Organization *Organization `json:"organization"`
	// The user owner of the subscription.
	User *User `json:"user"`
	// Whether the subscription is the default for the account.
	Default bool `json:"default"`
	// The subscription view.
	View string `json:"view"`
	// The creation time of the subscription.
	CreateTime time.Time `json:"createTime"`
	// The last update time of the subscription.
	UpdateTime time.Time `json:"updateTime"`
}
