// Code generated. DO NOT EDIT.

package adminapi

import (
	"github.com/userhubdev/go-sdk/internal"
)

// The Admin API.
type Client interface {
	// The flow methods.
	Flows() Flows
	// The invoice methods.
	Invoices() Invoices
	// The organization methods.
	Organizations() Organizations
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

func (c *clientImpl) Flows() Flows {
	return &flowsImpl{transport: c.transport}
}

func (c *clientImpl) Invoices() Invoices {
	return &invoicesImpl{transport: c.transport}
}

func (c *clientImpl) Organizations() Organizations {
	return &organizationsImpl{transport: c.transport}
}

func (c *clientImpl) Subscriptions() Subscriptions {
	return &subscriptionsImpl{transport: c.transport}
}

func (c *clientImpl) Users() Users {
	return &usersImpl{transport: c.transport}
}
