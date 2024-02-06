// Code generated. DO NOT EDIT.

package adminv1

import (
	"github.com/userhubdev/go-sdk/commonv1"
)

// The actual plan within the plan group. This defines the associated
// connection and billing interval.
type PlanGroupRevisionPlan struct {
	// The client defined unique identifier for the plan (e.g. `monthly`).
	UniqueId string `json:"uniqueId"`
	// The details of the associated connection.
	Connection *Connection `json:"connection"`
	// The billing interval for the plan.
	Interval *commonv1.Interval `json:"interval"`
	// The customer facing human-readable display name for the plan.
	DisplayName string `json:"displayName"`
	// The admin facing description of the plan.
	//
	// The maximum length is 1000 characters.
	Description string `json:"description"`
	// The prices associated with the plan.
	//
	// There should be a price for each product/currency
	// combination.
	Prices []*Price `json:"prices"`
	// The visibility of the plan.
	Visibility string `json:"visibility"`
}
