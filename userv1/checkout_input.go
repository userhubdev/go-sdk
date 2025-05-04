// Code generated. DO NOT EDIT.

package userv1

// Checkout input parameters.
type CheckoutInput struct {
	// The identifier of the organization.
	//
	// This must be provided for organization checkouts.
	OrganizationId string `json:"organizationId"`
	// The type of the checkout.
	Type string `json:"type"`
	// The identifier of the plan.
	//
	// This allows you to specify the currently selected plan.
	PlanId string `json:"planId"`
}
