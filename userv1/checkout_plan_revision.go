// Code generated. DO NOT EDIT.

package userv1

// Revision is a collection of plans with the same revision.
//
// This will group plans with different billing cycles.
type CheckoutPlanRevision struct {
	// The system-assigned identifier of the plan group revision.
	Id string `json:"id"`
	// Whether this is the current revision for the subscription.
	Current bool `json:"current"`
	// Whether this is the selected revision for the checkout.
	Selected bool `json:"selected"`
	// Whether this is the latest revision for the plan group.
	//
	// This will only be set for a current or selected plan group.
	Latest bool `json:"latest"`
}
