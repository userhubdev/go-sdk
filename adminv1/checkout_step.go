// Code generated. DO NOT EDIT.

package adminv1

// The checkout step.
type CheckoutStep struct {
	// The type of step.
	Type string `json:"type"`
	// The state of the step.
	State string `json:"state"`
	// The trial step details.
	Trial *CheckoutTrialStep `json:"trial"`
	// The cancel step details.
	Cancel *CheckoutCancelStep `json:"cancel"`
	// The complete payment step details.
	CompletePayment *CheckoutCompletePaymentStep `json:"completePayment"`
}
