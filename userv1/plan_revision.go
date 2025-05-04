// Code generated. DO NOT EDIT.

package userv1

// The revision information for the plan.
type PlanRevision struct {
	// The system-assigned identifier of the plan revision.
	Id string `json:"id"`
	// Whether this is the current revision for the subscription.
	//
	// This is only set in checkout.
	Current bool `json:"current"`
	// Whether this is the selected revision.
	//
	// This is only set in checkout.
	Selected bool `json:"selected"`
	// Whether this is the latest revision for the plan.
	//
	// This is only set for a current or selected plan in checkout.
	Latest bool `json:"latest"`
}
