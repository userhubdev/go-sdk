// Code generated. DO NOT EDIT.

package adminapi

import (
	"context"
	"fmt"
	"net/url"

	"github.com/userhubdev/go-sdk/adminv1"
	"github.com/userhubdev/go-sdk/internal"
)

type Invoices interface {
	// List invoices.
	List(ctx context.Context, input *InvoiceListInput) (*adminv1.ListInvoicesResponse, error)
	// Get an invoice.
	Get(ctx context.Context, invoiceId string, input *InvoiceGetInput) (*adminv1.Invoice, error)
	// Pay an invoice.
	Pay(ctx context.Context, invoiceId string, input *InvoicePayInput) (*adminv1.Invoice, error)
}

type invoicesImpl struct {
	transport internal.Transport
}

// InvoiceListInput is the input param for the List method.
type InvoiceListInput struct {
	// Filter results by organization identifier.
	//
	// This is required if the user identifier is not specified.
	OrganizationId string
	// Filter results by user identifier.
	//
	// This is required if the organization identifier is not specified.
	UserId string
	// The maximum number of invoices to return. The API may return fewer than
	// this value.
	//
	// If unspecified, at most 20 invoices will be returned.
	// The maximum value is 100; values above 100 will be coerced to 100.
	PageSize int32
	// A page token, received from a previous list invoices call.
	// Provide this to retrieve the subsequent page.
	//
	// When paginating, all other parameters provided to list invoices must match
	// the call that provided the page token.
	PageToken string
	// A comma-separated list of fields to order by.
	OrderBy string
	// The Invoice view to return in the results.
	//
	// This defaults to the `BASIC` view.
	View string
}

func (n *invoicesImpl) List(ctx context.Context, input *InvoiceListInput) (*adminv1.ListInvoicesResponse, error) {
	req := internal.NewRequest(
		"admin.invoices.list",
		"GET",
		"/admin/v1/invoices",
	)
	req.SetIdempotent(true)

	if input != nil {
		if !internal.IsEmpty(input.OrganizationId) {
			req.SetQuery("organizationId", input.OrganizationId)
		}
		if !internal.IsEmpty(input.UserId) {
			req.SetQuery("userId", input.UserId)
		}
		if !internal.IsEmpty(input.PageSize) {
			req.SetQuery("pageSize", input.PageSize)
		}
		if !internal.IsEmpty(input.PageToken) {
			req.SetQuery("pageToken", input.PageToken)
		}
		if !internal.IsEmpty(input.OrderBy) {
			req.SetQuery("orderBy", input.OrderBy)
		}
		if !internal.IsEmpty(input.View) {
			req.SetQuery("view", input.View)
		}
	}

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &adminv1.ListInvoicesResponse{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// InvoiceGetInput is the input param for the Get method.
type InvoiceGetInput struct {
	// Restrict by organization identifier.
	OrganizationId string
	// Restrict by user identifier.
	UserId string
}

func (n *invoicesImpl) Get(ctx context.Context, invoiceId string, input *InvoiceGetInput) (*adminv1.Invoice, error) {
	req := internal.NewRequest(
		"admin.invoices.get",
		"GET",
		fmt.Sprintf("/admin/v1/invoices/%s",
			url.PathEscape(invoiceId),
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

	model := &adminv1.Invoice{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// InvoicePayInput is the input param for the Pay method.
type InvoicePayInput struct {
	// Restrict by organization identifier.
	OrganizationId string
	// Restrict by user identifier.
	UserId string
	// The identifier of the payment method.
	//
	// The default payment method will be used if not specified.
	PaymentMethodId string
}

func (n *invoicesImpl) Pay(ctx context.Context, invoiceId string, input *InvoicePayInput) (*adminv1.Invoice, error) {
	req := internal.NewRequest(
		"admin.invoices.pay",
		"POST",
		fmt.Sprintf("/admin/v1/invoices/%s:pay",
			url.PathEscape(invoiceId),
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
		if !internal.IsEmpty(input.PaymentMethodId) {
			body["paymentMethodId"] = input.PaymentMethodId
		}
	}

	req.SetBody(body)

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &adminv1.Invoice{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}
