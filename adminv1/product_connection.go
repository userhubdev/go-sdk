// Code generated. DO NOT EDIT.

package adminv1

import (
	"time"
)

// A link between a UserHub product and an external product.
type ProductConnection struct {
	// The basic view of the connection.
	Connection *Connection `json:"connection"`
	// The external identifier of the connected product.
	ExternalId string `json:"externalId"`
	// The status of the connected product.
	State string `json:"state"`
	// The code that best describes the reason for the state.
	StateReason string `json:"stateReason"`
	// The creation time of the product connection.
	CreateTime time.Time `json:"createTime"`
	// The last update time of the product connection.
	UpdateTime time.Time `json:"updateTime"`
}
