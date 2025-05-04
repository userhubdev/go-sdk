// Code generated. DO NOT EDIT.

package adminapi

import (
	"github.com/userhubdev/go-sdk/internal"
)

// The Admin API.
type Client interface {
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
	// The subscription methods.
	Subscriptions() Subscriptions
	// The user methods.
	Users() Users
}

type clientImpl struct {
	transport internal.Transport
}

func newClient(transport internal.Transport) Client {
	return &clientImpl{transport: transport}
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

func (c *clientImpl) Subscriptions() Subscriptions {
	return &subscriptionsImpl{transport: c.transport}
}

func (c *clientImpl) Users() Users {
	return &usersImpl{transport: c.transport}
}
