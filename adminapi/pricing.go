// Code generated. DO NOT EDIT.

package adminapi

import (
	"context"

	"github.com/userhubdev/go-sdk/adminv1"
	"github.com/userhubdev/go-sdk/internal"
)

type Pricing interface {
	// Get pricing.
	Get(ctx context.Context, input *PricingGetInput) (*adminv1.Pricing, error)
}

type pricingImpl struct {
	transport internal.Transport
}

// PricingGetInput is the input param for the Get method.
type PricingGetInput struct {
	// Whether to get pricing for users or organizations.
	//
	// This is not required if either user ID or organization ID is specified
	// and will default to user if no options are specified.
	AccountType string
	// Show pricing for specified organization.
	OrganizationId string
	// Show pricing for the specified user.
	UserId string
}

func (n *pricingImpl) Get(ctx context.Context, input *PricingGetInput) (*adminv1.Pricing, error) {
	req := internal.NewRequest(
		"admin.pricing.get",
		"GET",
		"/admin/v1/pricing",
	)
	req.SetIdempotent(true)

	if input != nil {
		if !internal.IsEmpty(input.AccountType) {
			req.SetQuery("accountType", input.AccountType)
		}
		if !internal.IsEmpty(input.OrganizationId) {
			req.SetQuery("organizationId", input.OrganizationId)
		}
		if !internal.IsEmpty(input.UserId) {
			req.SetQuery("userId", input.UserId)
		}
	}

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &adminv1.Pricing{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}
