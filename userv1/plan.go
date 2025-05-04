// Code generated. DO NOT EDIT.

package userv1

import (
	"github.com/userhubdev/go-sdk/commonv1"
)

// The plan.
type Plan struct {
	// The system-assigned identifier of the plan.
	Id string `json:"id"`
	// The status of the plan.
	State string `json:"state"`
	// The client defined unique identifier of the plan.
	UniqueId string `json:"uniqueId"`
	// The name of the plan.
	DisplayName string `json:"displayName"`
	// The description of the plan.
	Description string `json:"description"`
	// The tier for the plan.
	//
	// This is only available in checkout and pricing.
	Tier string `json:"tier"`
	// The currency code for the plan (e.g. `USD`).
	CurrencyCode string `json:"currencyCode"`
	// The billing interval for the plan.
	Interval *commonv1.Interval `json:"interval"`
	// The revision for the plan.
	Revision *PlanRevision `json:"revision"`
	// Whether this is the current plan for the subscription.
	//
	// This is only set in checkout.
	Current bool `json:"current"`
	// Whether this is the selected plan.
	//
	// This is only set in checkout.
	Selected bool `json:"selected"`
	// Whether this is the default term for the plan.
	Default bool `json:"default"`
	// The trial settings.
	//
	// For authenticated requests, this will only be set for accounts that
	// are eligible for a new trial.
	Trial *PlanTrial `json:"trial"`
	// Whether the change is considered an upgrade or
	// a downgrade.
	//
	// This is only set in checkout.
	ChangePath string `json:"changePath"`
	// The savings for the plan.
	//
	// The savings are in comparison to the plan in the revision with the
	// shortest billing interval (normally monthly).
	Savings *PlanSavings `json:"savings"`
	// The items associated with plan.
	Items []*PlanItem `json:"items"`
}
