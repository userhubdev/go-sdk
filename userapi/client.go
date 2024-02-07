// Code generated. DO NOT EDIT.

package userapi

import (
	"github.com/userhubdev/go-sdk/internal"
)

// The User API.
type Client interface {
	// The flow methods.
	Flows() Flows
	// The invoice methods.
	Invoices() Invoices
	// The organization methods.
	Organizations() Organizations
	// The Portal session.
	Session() Session
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

func (c *clientImpl) Session() Session {
	return &sessionImpl{transport: c.transport}
}
