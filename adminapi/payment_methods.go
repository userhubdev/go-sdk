// Code generated. DO NOT EDIT.

package adminapi

import (
	"context"
	"fmt"
	"net/url"

	"github.com/userhubdev/go-sdk/adminv1"
	"github.com/userhubdev/go-sdk/apiv1"
	"github.com/userhubdev/go-sdk/commonv1"
	"github.com/userhubdev/go-sdk/internal"
	"github.com/userhubdev/go-sdk/types"
)

type PaymentMethods interface {
	// Create a payment method.
	Create(ctx context.Context, input *PaymentMethodCreateInput) (*adminv1.PaymentMethod, error)
	// Create a payment method intent.
	//
	// This can be used with a third-party billing provider API
	// to store a payment method.
	CreateIntent(ctx context.Context, input *PaymentMethodCreateIntentInput) (*adminv1.PaymentMethodIntent, error)
	// Get a payment method.
	Get(ctx context.Context, paymentMethodId string, input *PaymentMethodGetInput) (*adminv1.PaymentMethod, error)
	// Update a payment method.
	Update(ctx context.Context, paymentMethodId string, input *PaymentMethodUpdateInput) (*adminv1.PaymentMethod, error)
	// Set a default payment method for an account.
	SetDefault(ctx context.Context, paymentMethodId string, input *PaymentMethodSetDefaultInput) (*adminv1.PaymentMethod, error)
	// Delete a payment method.
	Delete(ctx context.Context, paymentMethodId string, input *PaymentMethodDeleteInput) (*apiv1.EmptyResponse, error)
}

type paymentMethodsImpl struct {
	transport internal.Transport
}

// PaymentMethodCreateInput is the input param for the Create method.
type PaymentMethodCreateInput struct {
	// The identifier of the organization.
	//
	// This is required if the user identifier not specified.
	OrganizationId string
	// The identifier of the user.
	//
	// This is required if the organization identifier not specified.
	UserId string
	// The identifier of the connection.
	ConnectionId string
	// The external identifier of the payment method to connect.
	ExternalId string
	// Whether to set the payment method as the default.
	//
	// This defaults to true.
	Default bool
}

func (n *paymentMethodsImpl) Create(ctx context.Context, input *PaymentMethodCreateInput) (*adminv1.PaymentMethod, error) {
	req := internal.NewRequest(
		"admin.payment_methods.create",
		"POST",
		"/admin/v1/paymentMethods",
	)

	body := map[string]any{}

	if input != nil {
		if !internal.IsEmpty(input.OrganizationId) {
			body["organizationId"] = input.OrganizationId
		}
		if !internal.IsEmpty(input.UserId) {
			body["userId"] = input.UserId
		}
		if !internal.IsEmpty(input.ConnectionId) {
			body["connectionId"] = input.ConnectionId
		}
		if !internal.IsEmpty(input.ExternalId) {
			body["externalId"] = input.ExternalId
		}
		if !internal.IsEmpty(input.Default) {
			body["default"] = input.Default
		}
	}

	req.SetBody(body)

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &adminv1.PaymentMethod{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// PaymentMethodCreateIntentInput is the input param for the CreateIntent method.
type PaymentMethodCreateIntentInput struct {
	// The identifier of the organization.
	//
	// This is required if the user identifier is not specified.
	OrganizationId string
	// The identifier of the user.
	//
	// This is required if the organization identifier is not not
	// specified.
	UserId string
	// The identifier of the connection.
	ConnectionId string
}

func (n *paymentMethodsImpl) CreateIntent(ctx context.Context, input *PaymentMethodCreateIntentInput) (*adminv1.PaymentMethodIntent, error) {
	req := internal.NewRequest(
		"admin.payment_methods.createIntent",
		"POST",
		"/admin/v1/paymentMethods:createIntent",
	)

	body := map[string]any{}

	if input != nil {
		if !internal.IsEmpty(input.OrganizationId) {
			body["organizationId"] = input.OrganizationId
		}
		if !internal.IsEmpty(input.UserId) {
			body["userId"] = input.UserId
		}
		if !internal.IsEmpty(input.ConnectionId) {
			body["connectionId"] = input.ConnectionId
		}
	}

	req.SetBody(body)

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &adminv1.PaymentMethodIntent{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// PaymentMethodGetInput is the input param for the Get method.
type PaymentMethodGetInput struct {
	// The identifier of the organization.
	//
	// Optionally restrict update to payment methods owned by
	// this organization.
	OrganizationId string
	// The identifier of the user.
	//
	// Optionally restrict update to payment methods owned by
	// this user.
	UserId string
}

func (n *paymentMethodsImpl) Get(ctx context.Context, paymentMethodId string, input *PaymentMethodGetInput) (*adminv1.PaymentMethod, error) {
	req := internal.NewRequest(
		"admin.payment_methods.get",
		"GET",
		fmt.Sprintf("/admin/v1/paymentMethods/%s",
			url.PathEscape(paymentMethodId),
		),
	)
	req.SetIdempotent(true)

	if input != nil {
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

	model := &adminv1.PaymentMethod{}

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

	// The identifier of the organization.
	//
	// Optionally restrict update to payment methods owned by
	// this organization.
	OrganizationId string
	// The identifier of the user.
	//
	// Optionally restrict update to payment methods owned by
	// this user.
	UserId string
}

func (n *paymentMethodsImpl) Update(ctx context.Context, paymentMethodId string, input *PaymentMethodUpdateInput) (*adminv1.PaymentMethod, error) {
	req := internal.NewRequest(
		"admin.payment_methods.update",
		"PATCH",
		fmt.Sprintf("/admin/v1/paymentMethods/%s",
			url.PathEscape(paymentMethodId),
		),
	)
	req.SetIdempotent(true)

	body := map[string]any{}

	if input != nil {
		if !internal.IsEmpty(input.OrganizationId) {
			req.SetQuery("organizationId", input.OrganizationId)
		}
		if !internal.IsEmpty(input.UserId) {
			req.SetQuery("userId", input.UserId)
		}
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

	model := &adminv1.PaymentMethod{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// PaymentMethodSetDefaultInput is the input param for the SetDefault method.
type PaymentMethodSetDefaultInput struct {
	// The identifier of the organization.
	//
	// Optionally restrict set default to payment methods owned by
	// this organization.
	OrganizationId string
	// The identifier of the user.
	//
	// Optionally restrict set default to payment methods owned by
	// this user.
	UserId string
}

func (n *paymentMethodsImpl) SetDefault(ctx context.Context, paymentMethodId string, input *PaymentMethodSetDefaultInput) (*adminv1.PaymentMethod, error) {
	req := internal.NewRequest(
		"admin.payment_methods.setDefault",
		"POST",
		fmt.Sprintf("/admin/v1/paymentMethods/%s:setDefault",
			url.PathEscape(paymentMethodId),
		),
	)

	body := map[string]any{}

	if input != nil {
		if !internal.IsEmpty(input.OrganizationId) {
			body["organizationId"] = input.OrganizationId
		}
		if !internal.IsEmpty(input.UserId) {
			body["userId"] = input.UserId
		}
	}

	req.SetBody(body)

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &adminv1.PaymentMethod{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// PaymentMethodDeleteInput is the input param for the Delete method.
type PaymentMethodDeleteInput struct {
	// The identifier of the organization.
	//
	// Optionally restrict delete to payment methods owned by
	// this organization.
	OrganizationId string
	// The identifier of the user.
	//
	// Optionally restrict delete to payment methods owned by
	// this user.
	UserId string
}

func (n *paymentMethodsImpl) Delete(ctx context.Context, paymentMethodId string, input *PaymentMethodDeleteInput) (*apiv1.EmptyResponse, error) {
	req := internal.NewRequest(
		"admin.payment_methods.delete",
		"DELETE",
		fmt.Sprintf("/admin/v1/paymentMethods/%s",
			url.PathEscape(paymentMethodId),
		),
	)

	if input != nil {
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

	model := &apiv1.EmptyResponse{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}
