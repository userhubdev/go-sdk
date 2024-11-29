// Code generated. DO NOT EDIT.

package userv1

// Plan group is a collection of related plans.
//
// A plan group will generally describe a tier of plans offered
// (e.g. Basic vs Pro) which might contain multiple billing options
// (e.g. monthly vs yearly, USD vs EUR).
type CheckoutPlanGroup struct {
	// The system-assigned identifier of the plan group.
	Id string `json:"id"`
	// The client defined unique identifier of the plan group (e.g. `pro`).
	UniqueId string `json:"uniqueId"`
	// Whether this is the current plan group for the subscription.
	Current bool `json:"current"`
	// Whether this is the selected plan group for the checkout.
	Selected bool `json:"selected"`
}
