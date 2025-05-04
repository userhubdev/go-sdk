// Code generated. DO NOT EDIT.

package userapi

import (
	"context"

	"github.com/userhubdev/go-sdk/internal"
	"github.com/userhubdev/go-sdk/userv1"
)

type Pricing interface {
	// Get pricing.
	Get(ctx context.Context, input *PricingGetInput) (*userv1.Pricing, error)
}

type pricingImpl struct {
	transport internal.Transport
}

// PricingGetInput is the input param for the Get method.
type PricingGetInput struct {
	// Whether to get pricing for users or organizations.
	//
	// This is not required if organization ID is specified
	// and will default to user if no options are specified.
	AccountType string
	// Show pricing for specified organization.
	//
	// If this and account type are not specified, it will default to
	// the requesting user for authenticated requests.
	OrganizationId string
	// Always get the current public pricing.
	//
	// For unauthenticated requests, this is always true.
	Public bool
}

func (n *pricingImpl) Get(ctx context.Context, input *PricingGetInput) (*userv1.Pricing, error) {
	req := internal.NewRequest(
		"user.pricing.get",
		"GET",
		"/user/v1/pricing",
	)
	req.SetIdempotent(true)

	if input != nil {
		if !internal.IsEmpty(input.AccountType) {
			req.SetQuery("accountType", input.AccountType)
		}
		if !internal.IsEmpty(input.OrganizationId) {
			req.SetQuery("organizationId", input.OrganizationId)
		}
		if !internal.IsEmpty(input.Public) {
			req.SetQuery("public", input.Public)
		}
	}

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &userv1.Pricing{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}
