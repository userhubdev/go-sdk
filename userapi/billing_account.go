// Code generated. DO NOT EDIT.

package userapi

import (
	"context"

	"github.com/userhubdev/go-sdk/commonv1"
	"github.com/userhubdev/go-sdk/internal"
	"github.com/userhubdev/go-sdk/types"
	"github.com/userhubdev/go-sdk/userv1"
)

type BillingAccount interface {
	// Get the default billing account.
	Get(ctx context.Context, input *BillingAccountGetInput) (*userv1.BillingAccount, error)
	// Update the default billing account.
	Update(ctx context.Context, input *BillingAccountUpdateInput) (*userv1.BillingAccount, error)
}

type billingAccountImpl struct {
	transport internal.Transport
}

// BillingAccountGetInput is the input param for the Get method.
type BillingAccountGetInput struct {
	// The identifier of the organization.
	//
	// If not specified the billing account for the user is returned.
	OrganizationId string
}

func (n *billingAccountImpl) Get(ctx context.Context, input *BillingAccountGetInput) (*userv1.BillingAccount, error) {
	req := internal.NewRequest(
		"user.billing_account.get",
		"GET",
		"/user/v1/billingAccount",
	)
	req.SetIdempotent(true)

	if input != nil {
		if !internal.IsEmpty(input.OrganizationId) {
			req.SetQuery("organizationId", input.OrganizationId)
		}
	}

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &userv1.BillingAccount{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// BillingAccountUpdateInput is the input param for the Update method.
type BillingAccountUpdateInput struct {
	// The human-readable display name of the billing account.
	//
	// The maximum length is 200 characters.
	//
	// This might be further restricted by the billing provider.
	DisplayName types.Optional[string]
	// The email address of the billing account.
	//
	// The maximum length is 320 characters.
	//
	// This might be further restricted by the billing provider.
	Email types.Optional[string]
	// The E164 phone number of the billing account (e.g. `+12125550123`).
	PhoneNumber types.Optional[string]
	// The address for the billing account.
	Address types.Optional[*commonv1.Address]

	// The identifier of the organization.
	//
	// If not specified the billing account for the user is used.
	OrganizationId string
}

func (n *billingAccountImpl) Update(ctx context.Context, input *BillingAccountUpdateInput) (*userv1.BillingAccount, error) {
	req := internal.NewRequest(
		"user.billing_account.update",
		"PATCH",
		"/user/v1/billingAccount",
	)
	req.SetIdempotent(true)

	body := map[string]any{}

	if input != nil {
		if !internal.IsEmpty(input.OrganizationId) {
			req.SetQuery("organizationId", input.OrganizationId)
		}
		if input.DisplayName.Present {
			body["displayName"] = input.DisplayName.Value
		}
		if input.Email.Present {
			body["email"] = input.Email.Value
		}
		if input.PhoneNumber.Present {
			body["phoneNumber"] = input.PhoneNumber.Value
		}
		if input.Address.Present {
			body["address"] = input.Address.Value
		}
	}

	req.SetBody(body)

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &userv1.BillingAccount{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}
