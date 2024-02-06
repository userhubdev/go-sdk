// Code generated. DO NOT EDIT.

package adminv1

import (
	"time"
)

// Product describes a service a tenant provides.
type Product struct {
	// The system-assigned identifier of the product.
	Id string `json:"id"`
	// The client defined unique identifier of the product.
	//
	// It is restricted to letters, numbers, underscores, and hyphens,
	// with the first character a letter or a number, and a 255
	// character maximum.
	//
	// ID's starting with `pd_` are reserved.
	UniqueId string `json:"uniqueId"`
	// The customer facing human-readable display name of the product.
	//
	// The maximum length is 200 characters.
	DisplayName string `json:"displayName"`
	// The public description of the product.
	//
	// The maximum length is 1000 characters.
	Description string `json:"description"`
	// Whether the price has been committed.
	//
	// This is automatically set the first time the product is used
	// in a plan.
	Committed bool `json:"committed"`
	// The archived status of the product.
	//
	// It determines if the product can be activated by self-serve plans.
	Archived bool `json:"archived"`
	// The connected products.
	ProductConnections []*ProductConnection `json:"productConnections"`
	// The creation time of the product.
	CreateTime time.Time `json:"createTime"`
	// The last update time of the product.
	UpdateTime time.Time `json:"updateTime"`
}
