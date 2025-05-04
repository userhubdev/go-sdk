// Code generated. DO NOT EDIT.

package userv1

// The plan information.
type AccountSubscriptionPlan struct {
	// The identifier of the plan.
	Id string `json:"id"`
	// The client defined unique identifier of the plan.
	UniqueId string `json:"uniqueId"`
	// The human-readable display name of the plan.
	DisplayName string `json:"displayName"`
	// The plan product.
	Product *Product `json:"product"`
}
