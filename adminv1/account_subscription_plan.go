// Code generated. DO NOT EDIT.

package adminv1

// The plan information.
type AccountSubscriptionPlan struct {
	// The identifier of the plan.
	Id string `json:"id"`
	// The human-readable display name of the plan.
	DisplayName string `json:"displayName"`
	// The plan product.
	Product *AccountSubscriptionProduct `json:"product"`
}
