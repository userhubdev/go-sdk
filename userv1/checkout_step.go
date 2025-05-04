// Code generated. DO NOT EDIT.

package userv1

// The checkout step.
type CheckoutStep struct {
	// The type of step.
	Type string `json:"type"`
	// The state of the step.
	State string `json:"state"`
	// The complete payment step details.
	CompletePayment *CheckoutCompletePaymentStep `json:"completePayment"`
}
