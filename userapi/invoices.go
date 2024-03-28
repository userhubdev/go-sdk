// Code generated. DO NOT EDIT.

package userapi

import (
	"context"
	"fmt"
	"net/url"

	"github.com/userhubdev/go-sdk/internal"
	"github.com/userhubdev/go-sdk/userv1"
)

type Invoices interface {
	// Lists invoices.
	List(ctx context.Context, input *InvoiceListInput) (*userv1.ListInvoicesResponse, error)
	// Retrieves specified invoice.
	Get(ctx context.Context, invoiceId string, input *InvoiceGetInput) (*userv1.Invoice, error)
}

type invoicesImpl struct {
	transport internal.Transport
}

// InvoiceListInput is the input param for the List method.
type InvoiceListInput struct {
	// Show results for specified organization.
	//
	// If this is not provided the user's individual subscription(s)
	// will be returned.
	OrganizationId string
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
	//
	// Supports:
	// - `createTime asc`
	// - `createTime desc`
	OrderBy string
}

func (n *invoicesImpl) List(ctx context.Context, input *InvoiceListInput) (*userv1.ListInvoicesResponse, error) {
	req := internal.NewRequest(
		"user.invoices.list",
		"GET",
		"/user/v1/invoices",
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
		if !internal.IsEmpty(input.OrderBy) {
			req.SetQuery("orderBy", input.OrderBy)
		}
	}

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &userv1.ListInvoicesResponse{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// InvoiceGetInput is the input param for the Get method.
type InvoiceGetInput struct {
}

func (n *invoicesImpl) Get(ctx context.Context, invoiceId string, input *InvoiceGetInput) (*userv1.Invoice, error) {
	req := internal.NewRequest(
		"user.invoices.get",
		"GET",
		fmt.Sprintf("/user/v1/invoices/%s",
			url.PathEscape(invoiceId),
		),
	)
	req.SetIdempotent(true)

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &userv1.Invoice{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}
