// Code generated. DO NOT EDIT.

package adminv1

import (
	"time"
)

// A link between a account and an external account.
type AccountConnection struct {
	// The tenant connection.
	Connection *Connection `json:"connection"`
	// The external identifier of the connected account.
	ExternalId string `json:"externalId"`
	// The external admin URL for the connected account.
	AdminUrl string `json:"adminUrl"`
	// The status of the connected account.
	State string `json:"state"`
	// The code that best describes the reason for the state.
	StateReason string `json:"stateReason"`
	// The balance amount for the account.
	//
	// A negative value indicates an amount which will be subtracted from the next
	// invoice (credit).
	//
	// A positive value indicates an amount which will be added to the next
	// invoice (debt).
	BalanceAmount string `json:"balanceAmount"`
	// The currency code for the account.
	CurrencyCode string `json:"currencyCode"`
	// The payment methods for connections that support it.
	PaymentMethods []*PaymentMethod `json:"paymentMethods"`
	// The last time the account was pulled from the connection.
	PullTime time.Time `json:"pullTime"`
	// The last time the account was pushed to the connection.
	PushTime time.Time `json:"pushTime"`
	// The creation time of the account connection.
	CreateTime time.Time `json:"createTime"`
	// The last update time of the account connection.
	UpdateTime time.Time `json:"updateTime"`
}
