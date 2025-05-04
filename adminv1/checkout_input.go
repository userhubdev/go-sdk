// Code generated. DO NOT EDIT.

package adminv1

// Checkout input parameters.
type CheckoutInput struct {
	// The identifier of the organization.
	//
	// This is required if the user identifier is not specified.
	OrganizationId string `json:"organizationId"`
	// The identifier of the user.
	//
	// This is required if the organization identifier is not specified.
	UserId string `json:"userId"`
	// The type of the checkout.
	Type string `json:"type"`
	// The identifier of the plan.
	//
	// This allows you to specify the currently selected plan.
	PlanId string `json:"planId"`
	// The identifier of the subscriptions.
	//
	// This allows you to specify a non-default subscription.
	SubscriptionId string `json:"subscriptionId"`
	// The identifier of the connection.
	//
	// This allows you to specify a non-default billing connection.
	ConnectionId string `json:"connectionId"`
}
