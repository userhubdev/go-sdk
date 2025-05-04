// Code generated. DO NOT EDIT.

package userapi

import (
	"context"
	"fmt"
	"net/url"

	"github.com/userhubdev/go-sdk/commonv1"
	"github.com/userhubdev/go-sdk/internal"
	"github.com/userhubdev/go-sdk/types"
	"github.com/userhubdev/go-sdk/userv1"
)

type Checkouts interface {
	// Create a new checkout.
	Create(ctx context.Context, input *CheckoutCreateInput) (*userv1.Checkout, error)
	// Get a checkout.
	Get(ctx context.Context, checkoutId string, input *CheckoutGetInput) (*userv1.Checkout, error)
	// Set plan for a checkout.
	SetPlan(ctx context.Context, checkoutId string, input *CheckoutSetPlanInput) (*userv1.Checkout, error)
	// Set terms for a checkout.
	//
	// This is generally used to select a billing cycle for
	// the plan.
	SetTerms(ctx context.Context, checkoutId string, input *CheckoutSetTermsInput) (*userv1.Checkout, error)
	// Set item quantities for a checkout.
	SetItems(ctx context.Context, checkoutId string, input *CheckoutSetItemsInput) (*userv1.Checkout, error)
	// Set payment method for a checkout.
	SetPaymentMethod(ctx context.Context, checkoutId string, input *CheckoutSetPaymentMethodInput) (*userv1.Checkout, error)
	// Set billing details for a checkout.
	SetBillingDetails(ctx context.Context, checkoutId string, input *CheckoutSetBillingDetailsInput) (*userv1.Checkout, error)
	// Add discount to a checkout.
	AddDiscount(ctx context.Context, checkoutId string, input *CheckoutAddDiscountInput) (*userv1.Checkout, error)
	// Remove discount from a checkout.
	RemoveDiscount(ctx context.Context, checkoutId string, input *CheckoutRemoveDiscountInput) (*userv1.Checkout, error)
	// Complete payment for a checkout.
	CompletePayment(ctx context.Context, checkoutId string, input *CheckoutCompletePaymentInput) (*userv1.Checkout, error)
	// Submit a checkout for processing.
	Submit(ctx context.Context, checkoutId string, input *CheckoutSubmitInput) (*userv1.Checkout, error)
	// Cancel a checkout.
	Cancel(ctx context.Context, checkoutId string, input *CheckoutCancelInput) (*userv1.Checkout, error)
}

type checkoutsImpl struct {
	transport internal.Transport
}

// CheckoutCreateInput is the input param for the Create method.
type CheckoutCreateInput struct {
	// The identifier of the organization.
	//
	// This must be provided for organization checkouts.
	OrganizationId string
	// The type of the checkout.
	Type string
	// The identifier of the plan.
	//
	// This allows you to specify the currently selected plan.
	PlanId string

	// Attempt to submit checkout if ready and due amount is zero.
	Submit bool
}

func (n *checkoutsImpl) Create(ctx context.Context, input *CheckoutCreateInput) (*userv1.Checkout, error) {
	req := internal.NewRequest(
		"user.checkouts.create",
		"POST",
		"/user/v1/checkouts",
	)

	body := map[string]any{}

	if input != nil {
		if !internal.IsEmpty(input.Submit) {
			req.SetQuery("submit", input.Submit)
		}
		if !internal.IsEmpty(input.OrganizationId) {
			body["organizationId"] = input.OrganizationId
		}
		if !internal.IsEmpty(input.Type) {
			body["type"] = input.Type
		}
		if !internal.IsEmpty(input.PlanId) {
			body["planId"] = input.PlanId
		}
	}

	req.SetBody(body)

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &userv1.Checkout{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// CheckoutGetInput is the input param for the Get method.
type CheckoutGetInput struct {
}

func (n *checkoutsImpl) Get(ctx context.Context, checkoutId string, input *CheckoutGetInput) (*userv1.Checkout, error) {
	req := internal.NewRequest(
		"user.checkouts.get",
		"GET",
		fmt.Sprintf("/user/v1/checkouts/%s",
			url.PathEscape(checkoutId),
		),
	)
	req.SetIdempotent(true)

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &userv1.Checkout{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// CheckoutSetPlanInput is the input param for the SetPlan method.
type CheckoutSetPlanInput struct {
	// The identifier of the plan.
	//
	// This is required if completed isn't set to true.
	PlanId string
	// Mark the step completed if it is optional.
	Completed types.Optional[bool]
}

func (n *checkoutsImpl) SetPlan(ctx context.Context, checkoutId string, input *CheckoutSetPlanInput) (*userv1.Checkout, error) {
	req := internal.NewRequest(
		"user.checkouts.setPlan",
		"POST",
		fmt.Sprintf("/user/v1/checkouts/%s:setPlan",
			url.PathEscape(checkoutId),
		),
	)

	body := map[string]any{}

	if input != nil {
		if !internal.IsEmpty(input.PlanId) {
			body["planId"] = input.PlanId
		}
		if input.Completed.Present {
			body["completed"] = input.Completed.Value
		}
	}

	req.SetBody(body)

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &userv1.Checkout{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// CheckoutSetTermsInput is the input param for the SetTerms method.
type CheckoutSetTermsInput struct {
	// The identifier of the plan.
	//
	// This is required if completed isn't set to true.
	PlanId string
	// Mark the step completed if it is optional.
	Completed types.Optional[bool]
}

func (n *checkoutsImpl) SetTerms(ctx context.Context, checkoutId string, input *CheckoutSetTermsInput) (*userv1.Checkout, error) {
	req := internal.NewRequest(
		"user.checkouts.setTerms",
		"POST",
		fmt.Sprintf("/user/v1/checkouts/%s:setTerms",
			url.PathEscape(checkoutId),
		),
	)

	body := map[string]any{}

	if input != nil {
		if !internal.IsEmpty(input.PlanId) {
			body["planId"] = input.PlanId
		}
		if input.Completed.Present {
			body["completed"] = input.Completed.Value
		}
	}

	req.SetBody(body)

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &userv1.Checkout{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// CheckoutSetItemsInput is the input param for the SetItems method.
type CheckoutSetItemsInput struct {
	// The items to update.
	Items []*userv1.CheckoutItemInput
	// Mark the step completed if it is optional.
	Completed types.Optional[bool]
}

func (n *checkoutsImpl) SetItems(ctx context.Context, checkoutId string, input *CheckoutSetItemsInput) (*userv1.Checkout, error) {
	req := internal.NewRequest(
		"user.checkouts.setItems",
		"POST",
		fmt.Sprintf("/user/v1/checkouts/%s:setItems",
			url.PathEscape(checkoutId),
		),
	)

	body := map[string]any{}

	if input != nil {
		if len(input.Items) > 0 {
			body["items"] = input.Items
		}
		if input.Completed.Present {
			body["completed"] = input.Completed.Value
		}
	}

	req.SetBody(body)

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &userv1.Checkout{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// CheckoutSetPaymentMethodInput is the input param for the SetPaymentMethod method.
type CheckoutSetPaymentMethodInput struct {
	// The identifier of the payment method.
	//
	// This is required if external ID isn't specified or completed
	// isn't set to true.
	PaymentMethodId string
	// The external identifier of the payment method to connect.
	//
	// This is required if payment method ID isn't specified or
	// completed isn't set to true.
	ExternalId string
	// Mark the step completed if it is optional.
	Completed types.Optional[bool]
}

func (n *checkoutsImpl) SetPaymentMethod(ctx context.Context, checkoutId string, input *CheckoutSetPaymentMethodInput) (*userv1.Checkout, error) {
	req := internal.NewRequest(
		"user.checkouts.setPaymentMethod",
		"POST",
		fmt.Sprintf("/user/v1/checkouts/%s:setPaymentMethod",
			url.PathEscape(checkoutId),
		),
	)

	body := map[string]any{}

	if input != nil {
		if !internal.IsEmpty(input.PaymentMethodId) {
			body["paymentMethodId"] = input.PaymentMethodId
		}
		if !internal.IsEmpty(input.ExternalId) {
			body["externalId"] = input.ExternalId
		}
		if input.Completed.Present {
			body["completed"] = input.Completed.Value
		}
	}

	req.SetBody(body)

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &userv1.Checkout{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// CheckoutSetBillingDetailsInput is the input param for the SetBillingDetails method.
type CheckoutSetBillingDetailsInput struct {
	// The company or individual's full name.
	//
	// The maximum length is 200 characters.
	FullName string
	// The billing details address.
	Address *commonv1.Address
	// Mark the step completed if it is optional.
	Completed types.Optional[bool]
}

func (n *checkoutsImpl) SetBillingDetails(ctx context.Context, checkoutId string, input *CheckoutSetBillingDetailsInput) (*userv1.Checkout, error) {
	req := internal.NewRequest(
		"user.checkouts.setBillingDetails",
		"POST",
		fmt.Sprintf("/user/v1/checkouts/%s:setBillingDetails",
			url.PathEscape(checkoutId),
		),
	)

	body := map[string]any{}

	if input != nil {
		if !internal.IsEmpty(input.FullName) {
			body["fullName"] = input.FullName
		}
		if !internal.IsEmpty(input.Address) {
			body["address"] = input.Address
		}
		if input.Completed.Present {
			body["completed"] = input.Completed.Value
		}
	}

	req.SetBody(body)

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &userv1.Checkout{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// CheckoutAddDiscountInput is the input param for the AddDiscount method.
type CheckoutAddDiscountInput struct {
	// The discount code.
	Code string
	// Mark the step completed if it is optional.
	Completed types.Optional[bool]
}

func (n *checkoutsImpl) AddDiscount(ctx context.Context, checkoutId string, input *CheckoutAddDiscountInput) (*userv1.Checkout, error) {
	req := internal.NewRequest(
		"user.checkouts.addDiscount",
		"POST",
		fmt.Sprintf("/user/v1/checkouts/%s:addDiscount",
			url.PathEscape(checkoutId),
		),
	)

	body := map[string]any{}

	if input != nil {
		if !internal.IsEmpty(input.Code) {
			body["code"] = input.Code
		}
		if input.Completed.Present {
			body["completed"] = input.Completed.Value
		}
	}

	req.SetBody(body)

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &userv1.Checkout{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// CheckoutRemoveDiscountInput is the input param for the RemoveDiscount method.
type CheckoutRemoveDiscountInput struct {
	// The identifier of the checkout discount.
	CheckoutDiscountId string
}

func (n *checkoutsImpl) RemoveDiscount(ctx context.Context, checkoutId string, input *CheckoutRemoveDiscountInput) (*userv1.Checkout, error) {
	req := internal.NewRequest(
		"user.checkouts.removeDiscount",
		"POST",
		fmt.Sprintf("/user/v1/checkouts/%s:removeDiscount",
			url.PathEscape(checkoutId),
		),
	)

	body := map[string]any{}

	if input != nil {
		if !internal.IsEmpty(input.CheckoutDiscountId) {
			body["checkoutDiscountId"] = input.CheckoutDiscountId
		}
	}

	req.SetBody(body)

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &userv1.Checkout{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// CheckoutCompletePaymentInput is the input param for the CompletePayment method.
type CheckoutCompletePaymentInput struct {
}

func (n *checkoutsImpl) CompletePayment(ctx context.Context, checkoutId string, input *CheckoutCompletePaymentInput) (*userv1.Checkout, error) {
	req := internal.NewRequest(
		"user.checkouts.completePayment",
		"POST",
		fmt.Sprintf("/user/v1/checkouts/%s:completePayment",
			url.PathEscape(checkoutId),
		),
	)

	body := map[string]any{}

	req.SetBody(body)

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &userv1.Checkout{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// CheckoutSubmitInput is the input param for the Submit method.
type CheckoutSubmitInput struct {
}

func (n *checkoutsImpl) Submit(ctx context.Context, checkoutId string, input *CheckoutSubmitInput) (*userv1.Checkout, error) {
	req := internal.NewRequest(
		"user.checkouts.submit",
		"POST",
		fmt.Sprintf("/user/v1/checkouts/%s:submit",
			url.PathEscape(checkoutId),
		),
	)

	body := map[string]any{}

	req.SetBody(body)

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &userv1.Checkout{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// CheckoutCancelInput is the input param for the Cancel method.
type CheckoutCancelInput struct {
}

func (n *checkoutsImpl) Cancel(ctx context.Context, checkoutId string, input *CheckoutCancelInput) (*userv1.Checkout, error) {
	req := internal.NewRequest(
		"user.checkouts.cancel",
		"POST",
		fmt.Sprintf("/user/v1/checkouts/%s:cancel",
			url.PathEscape(checkoutId),
		),
	)

	body := map[string]any{}

	req.SetBody(body)

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &userv1.Checkout{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}
