// Code generated. DO NOT EDIT.

package adminv1

import (
	"github.com/userhubdev/go-sdk/commonv1"
)

// The plan.
type Plan struct {
	// The system-assigned identifier of the plan.
	Id string `json:"id"`
	// The name of the plan.
	DisplayName string `json:"displayName"`
	// The description of the plan.
	Description string `json:"description"`
	// The currency code for the plan (e.g. `USD`).
	CurrencyCode string `json:"currencyCode"`
	// The billing interval for the plan.
	BillingInterval *commonv1.Interval `json:"billingInterval"`
	// The tags associated with the revision.
	Tags []string `json:"tags"`
	// The items associated with plan.
	Items []*PlanItem `json:"items"`
}
