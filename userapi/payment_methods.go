// Code generated. DO NOT EDIT.

package userapi

import (
	"context"
	"fmt"
	"net/url"

	"github.com/userhubdev/go-sdk/apiv1"
	"github.com/userhubdev/go-sdk/commonv1"
	"github.com/userhubdev/go-sdk/internal"
	"github.com/userhubdev/go-sdk/types"
	"github.com/userhubdev/go-sdk/userv1"
)

type PaymentMethods interface {
	// List payment methods for an account.
	List(ctx context.Context, input *PaymentMethodListInput) (*userv1.ListPaymentMethodsResponse, error)
	// Create a new payment method.
	Create(ctx context.Context, input *PaymentMethodCreateInput) (*userv1.PaymentMethod, error)
	// Create a new payment method intent.
	//
	// This can be used with a third-party billing provider to
	// store a payment method.
	CreateIntent(ctx context.Context, input *PaymentMethodCreateIntentInput) (*userv1.PaymentMethodIntent, error)
	// Get a payment method.
	Get(ctx context.Context, paymentMethodId string, input *PaymentMethodGetInput) (*userv1.PaymentMethod, error)
	// Update a payment method.
	Update(ctx context.Context, paymentMethodId string, input *PaymentMethodUpdateInput) (*userv1.PaymentMethod, error)
	// Set a default payment method for an account.
	SetDefault(ctx context.Context, paymentMethodId string, input *PaymentMethodSetDefaultInput) (*userv1.PaymentMethod, error)
	// Delete a payment method.
	Delete(ctx context.Context, paymentMethodId string, input *PaymentMethodDeleteInput) (*apiv1.EmptyResponse, error)
}

type paymentMethodsImpl struct {
	transport internal.Transport
}

// PaymentMethodListInput is the input param for the List method.
type PaymentMethodListInput struct {
	// Show results for specified organization.
	//
	// If this is not provided the user's individual subscription(s)
	// will be returned.
	OrganizationId string
	// The maximum number of payment methods to return. The API may return fewer than
	// this value.
	//
	// If unspecified, at most 20 payment methods will be returned.
	// The maximum value is 100; values above 100 will be coerced to 100.
	PageSize int32
	// A page token, received from a previous list payment methods call.
	// Provide this to retrieve the subsequent page.
	//
	// When paginating, all other parameters provided to list payment methods must match
	// the call that provided the page token.
	PageToken string
}

func (n *paymentMethodsImpl) List(ctx context.Context, input *PaymentMethodListInput) (*userv1.ListPaymentMethodsResponse, error) {
	req := internal.NewRequest(
		"user.payment_methods.list",
		"GET",
		"/user/v1/paymentMethods",
	)
	req.SetIdempotent(true)

	if input != nil {
		if !internal.IsEmpty(input.OrganizationId) {
			req.SetQuery("organizationId", input.OrganizationId)
		}
		if !internal.IsEmpty(input.PageSize) {
			req.SetQuery("pageSize", input.PageSize)
		}
		if !internal.IsEmpty(input.PageToken) {
			req.SetQuery("pageToken", input.PageToken)
		}
	}

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &userv1.ListPaymentMethodsResponse{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// PaymentMethodCreateInput is the input param for the Create method.
type PaymentMethodCreateInput struct {
	// The identifier of the organization.
	OrganizationId string
	// The external identifier of the payment method to connect.
	ExternalId string
}

func (n *paymentMethodsImpl) Create(ctx context.Context, input *PaymentMethodCreateInput) (*userv1.PaymentMethod, error) {
	req := internal.NewRequest(
		"user.payment_methods.create",
		"POST",
		"/user/v1/paymentMethods",
	)

	body := map[string]any{}

	if input != nil {
		if !internal.IsEmpty(input.OrganizationId) {
			body["organizationId"] = input.OrganizationId
		}
		if !internal.IsEmpty(input.ExternalId) {
			body["externalId"] = input.ExternalId
		}
	}

	req.SetBody(body)

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &userv1.PaymentMethod{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// PaymentMethodCreateIntentInput is the input param for the CreateIntent method.
type PaymentMethodCreateIntentInput struct {
	// The identifier of the organization.
	OrganizationId string
}

func (n *paymentMethodsImpl) CreateIntent(ctx context.Context, input *PaymentMethodCreateIntentInput) (*userv1.PaymentMethodIntent, error) {
	req := internal.NewRequest(
		"user.payment_methods.createIntent",
		"POST",
		"/user/v1/paymentMethods:createIntent",
	)

	body := map[string]any{}

	if input != nil {
		if !internal.IsEmpty(input.OrganizationId) {
			body["organizationId"] = input.OrganizationId
		}
	}

	req.SetBody(body)

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &userv1.PaymentMethodIntent{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// PaymentMethodGetInput is the input param for the Get method.
type PaymentMethodGetInput struct {
}

func (n *paymentMethodsImpl) Get(ctx context.Context, paymentMethodId string, input *PaymentMethodGetInput) (*userv1.PaymentMethod, error) {
	req := internal.NewRequest(
		"user.payment_methods.get",
		"GET",
		fmt.Sprintf("/user/v1/paymentMethods/%s",
			url.PathEscape(paymentMethodId),
		),
	)
	req.SetIdempotent(true)

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &userv1.PaymentMethod{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// PaymentMethodUpdateInput is the input param for the Update method.
type PaymentMethodUpdateInput struct {
	// The full name of the owner of the payment method (e.g. `Jane Doe`).
	FullName types.Optional[string]
	// The address for the payment method.
	Address types.Optional[*commonv1.Address]
	// The card expiration year (e.g. `2030`).
	ExpYear types.Optional[int32]
	// The card expiration month (e.g. `12`).
	ExpMonth types.Optional[int32]
}

func (n *paymentMethodsImpl) Update(ctx context.Context, paymentMethodId string, input *PaymentMethodUpdateInput) (*userv1.PaymentMethod, error) {
	req := internal.NewRequest(
		"user.payment_methods.update",
		"PATCH",
		fmt.Sprintf("/user/v1/paymentMethods/%s",
			url.PathEscape(paymentMethodId),
		),
	)
	req.SetIdempotent(true)

	body := map[string]any{}

	if input != nil {
		if input.FullName.Present {
			body["fullName"] = input.FullName.Value
		}
		if input.Address.Present {
			body["address"] = input.Address.Value
		}
		if input.ExpYear.Present {
			body["expYear"] = input.ExpYear.Value
		}
		if input.ExpMonth.Present {
			body["expMonth"] = input.ExpMonth.Value
		}
	}

	req.SetBody(body)

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &userv1.PaymentMethod{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// PaymentMethodSetDefaultInput is the input param for the SetDefault method.
type PaymentMethodSetDefaultInput struct {
}

func (n *paymentMethodsImpl) SetDefault(ctx context.Context, paymentMethodId string, input *PaymentMethodSetDefaultInput) (*userv1.PaymentMethod, error) {
	req := internal.NewRequest(
		"user.payment_methods.setDefault",
		"POST",
		fmt.Sprintf("/user/v1/paymentMethods/%s:setDefault",
			url.PathEscape(paymentMethodId),
		),
	)

	body := map[string]any{}

	req.SetBody(body)

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &userv1.PaymentMethod{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// PaymentMethodDeleteInput is the input param for the Delete method.
type PaymentMethodDeleteInput struct {
}

func (n *paymentMethodsImpl) Delete(ctx context.Context, paymentMethodId string, input *PaymentMethodDeleteInput) (*apiv1.EmptyResponse, error) {
	req := internal.NewRequest(
		"user.payment_methods.delete",
		"DELETE",
		fmt.Sprintf("/user/v1/paymentMethods/%s",
			url.PathEscape(paymentMethodId),
		),
	)

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &apiv1.EmptyResponse{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}
