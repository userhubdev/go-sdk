// Code generated. DO NOT EDIT.

package userv1

import (
	"github.com/userhubdev/go-sdk/commonv1"
)

// The plan.
type CheckoutPlan struct {
	// The system-assigned identifier of the plan.
	Id string `json:"id"`
	// The name of the plan.
	DisplayName string `json:"displayName"`
	// The description of the plan.
	Description string `json:"description"`
	// The billing interval for the plan.
	BillingInterval *commonv1.Interval `json:"billingInterval"`
	// The plan group for the plan.
	Group *CheckoutPlanGroup `json:"group"`
	// The revision for the plan.
	Revision *CheckoutPlanRevision `json:"revision"`
	// Whether this is the current plan for the subscription.
	Current bool `json:"current"`
	// Whether this is the selected plan for the checkout.
	Selected bool `json:"selected"`
	// Whether this is the default plan for the plan group.
	Default bool `json:"default"`
	// The trial settings.
	//
	// For authenticated requests, this will not be set when the account
	// isn't eligible for a trial.
	Trial *CheckoutPlanTrial `json:"trial"`
	// Whether the change is considered an upgrade or
	// a downgrade.
	ChangePath string `json:"changePath"`
	// The flat price for the plan.
	//
	// This might not exists for plans that are billed by seat.
	Price *Price `json:"price"`
	// The primary seat price for the plan.
	SeatPrice *Price `json:"seatPrice"`
	// The savings for the plan.
	//
	// The savings are in comparison to the plan in the revision with the
	// shortest billing interval.
	Savings *CheckoutPlanSavings `json:"savings"`
	// Whether this plan is disabled for checkout.
	//
	// This will only be set when the current/selected plan is disabled, all
	// other plans will be available for checkout.
	Disabled bool `json:"disabled"`
}
