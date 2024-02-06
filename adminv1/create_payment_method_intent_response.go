// Code generated. DO NOT EDIT.

package adminv1

// Response message for CreatePaymentMethodIntent.
type CreatePaymentMethodIntentResponse struct {
	// The setup token for the billing system (e.g. Stripe SetupIntent
	// Client Secret). This is generally used by a frontend
	// client to securely setup a payment method, the result of which
	// can be passed to CreatePaymentMethod to complete the setup
	// process.
	Token string `json:"token"`
}
