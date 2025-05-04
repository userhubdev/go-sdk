// Code generated. DO NOT EDIT.

package userapi

import (
	"github.com/userhubdev/go-sdk/internal"
)

// The User API.
type Client interface {
	// The billing account methods.
	BillingAccount() BillingAccount
	// The checkout methods.
	Checkouts() Checkouts
	// The flow methods.
	Flows() Flows
	// The invoice methods.
	Invoices() Invoices
	// The organization methods.
	Organizations() Organizations
	// The payment method methods.
	PaymentMethods() PaymentMethods
	// The pricing methods.
	Pricing() Pricing
	// The role methods.
	Roles() Roles
	// The Portal session.
	Session() Session
}

type clientImpl struct {
	transport internal.Transport
}

func newClient(transport internal.Transport) Client {
	return &clientImpl{transport: transport}
}

func (c *clientImpl) BillingAccount() BillingAccount {
	return &billingAccountImpl{transport: c.transport}
}

func (c *clientImpl) Checkouts() Checkouts {
	return &checkoutsImpl{transport: c.transport}
}

func (c *clientImpl) Flows() Flows {
	return &flowsImpl{transport: c.transport}
}

func (c *clientImpl) Invoices() Invoices {
	return &invoicesImpl{transport: c.transport}
}

func (c *clientImpl) Organizations() Organizations {
	return &organizationsImpl{transport: c.transport}
}

func (c *clientImpl) PaymentMethods() PaymentMethods {
	return &paymentMethodsImpl{transport: c.transport}
}

func (c *clientImpl) Pricing() Pricing {
	return &pricingImpl{transport: c.transport}
}

func (c *clientImpl) Roles() Roles {
	return &rolesImpl{transport: c.transport}
}

func (c *clientImpl) Session() Session {
	return &sessionImpl{transport: c.transport}
}
