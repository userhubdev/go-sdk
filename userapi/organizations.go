// Code generated. DO NOT EDIT.

package userapi

import (
	"context"
	"fmt"
	"net/url"

	"github.com/userhubdev/go-sdk/apiv1"
	"github.com/userhubdev/go-sdk/internal"
	"github.com/userhubdev/go-sdk/types"
	"github.com/userhubdev/go-sdk/userv1"
)

type Organizations interface {
	// Lists organizations.
	List(ctx context.Context, input *OrganizationListInput) (*userv1.ListOrganizationsResponse, error)
	// Creates a new organization.
	Create(ctx context.Context, input *OrganizationCreateInput) (*userv1.Organization, error)
	// Retrieves specified organization.
	Get(ctx context.Context, organizationId string, input *OrganizationGetInput) (*userv1.Organization, error)
	// Updates specified organization.
	Update(ctx context.Context, organizationId string, input *OrganizationUpdateInput) (*userv1.Organization, error)
	// Delete specified organization.
	Delete(ctx context.Context, organizationId string, input *OrganizationDeleteInput) (*userv1.Organization, error)
	// Leave organization.
	//
	// This allows a user to remove themselves from an organization
	// without have permission to manage the organization.
	Leave(ctx context.Context, organizationId string, input *OrganizationLeaveInput) (*apiv1.EmptyResponse, error)
}

type organizationsImpl struct {
	transport internal.Transport
}

// OrganizationListInput is the input param for the List method.
type OrganizationListInput struct {
	// The maximum number of organizations to return. The API may return fewer than
	// this value.
	//
	// If unspecified, at most 20 organizations will be returned.
	// The maximum value is 100; values above 100 will be coerced to 100.
	PageSize int32
	// A page token, received from a previous list organizations call.
	// Provide this to retrieve the subsequent page.
	//
	// When paginating, all other parameters provided to list organizations must match
	// the call that provided the page token.
	PageToken string
	// A comma-separated list of fields to order by, sorted in ascending order.
	// Use `desc` after a field name for descending.
	//
	// Supported fields:
	// - `displayName`
	// - `email`
	OrderBy string
}

func (n *organizationsImpl) List(ctx context.Context, input *OrganizationListInput) (*userv1.ListOrganizationsResponse, error) {
	req := internal.NewRequest(
		"user.organizations.list",
		"GET",
		"/user/v1/organizations",
	)
	req.SetIdempotent(true)

	if input != nil {
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

	model := &userv1.ListOrganizationsResponse{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// OrganizationCreateInput is the input param for the Create method.
type OrganizationCreateInput struct {
	// The client defined unique identifier of the organization account.
	//
	// It is restricted to letters, numbers, underscores, and hyphens,
	// with the first character a letter or a number, and a 255
	// character maximum.
	//
	// ID's starting with `org_` are reserved.
	UniqueId string
	// The human-readable display name of the organization.
	//
	// The maximum length is 200 characters.
	DisplayName string
	// The email address of the organization.
	//
	// The maximum length is 320 characters.
	Email string
	// The flow identifier associated with the creation of the organization.
	//
	// The flow type must be `SIGNUP` and associated with the user creating the organization.
	FlowId string
}

func (n *organizationsImpl) Create(ctx context.Context, input *OrganizationCreateInput) (*userv1.Organization, error) {
	req := internal.NewRequest(
		"user.organizations.create",
		"POST",
		"/user/v1/organizations",
	)

	body := map[string]any{}

	if input != nil {
		if !internal.IsEmpty(input.UniqueId) {
			body["uniqueId"] = input.UniqueId
		}
		if !internal.IsEmpty(input.DisplayName) {
			body["displayName"] = input.DisplayName
		}
		if !internal.IsEmpty(input.Email) {
			body["email"] = input.Email
		}
		if !internal.IsEmpty(input.FlowId) {
			body["flowId"] = input.FlowId
		}
	}

	req.SetBody(body)

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &userv1.Organization{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// OrganizationGetInput is the input param for the Get method.
type OrganizationGetInput struct {
}

func (n *organizationsImpl) Get(ctx context.Context, organizationId string, input *OrganizationGetInput) (*userv1.Organization, error) {
	req := internal.NewRequest(
		"user.organizations.get",
		"GET",
		fmt.Sprintf("/user/v1/organizations/%s",
			url.PathEscape(organizationId),
		),
	)
	req.SetIdempotent(true)

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &userv1.Organization{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// OrganizationUpdateInput is the input param for the Update method.
type OrganizationUpdateInput struct {
	// The client defined unique identifier of the organization account.
	//
	// It is restricted to letters, numbers, underscores, and hyphens,
	// with the first character a letter or a number, and a 255
	// character maximum.
	//
	// ID's starting with `org_` are reserved.
	UniqueId types.Optional[string]
	// The human-readable display name of the organization.
	//
	// The maximum length is 200 characters.
	DisplayName types.Optional[string]
	// The email address of the organization.
	//
	// The maximum length is 320 characters.
	Email types.Optional[string]
	// The flow identifier associated with the creation of the organization.
	//
	// The flow type must be `SIGNUP` and associated with the user creating the organization.
	FlowId types.Optional[string]
}

func (n *organizationsImpl) Update(ctx context.Context, organizationId string, input *OrganizationUpdateInput) (*userv1.Organization, error) {
	req := internal.NewRequest(
		"user.organizations.update",
		"PATCH",
		fmt.Sprintf("/user/v1/organizations/%s",
			url.PathEscape(organizationId),
		),
	)
	req.SetIdempotent(true)

	body := map[string]any{}

	if input != nil {
		if input.UniqueId.Present {
			body["uniqueId"] = input.UniqueId.Value
		}
		if input.DisplayName.Present {
			body["displayName"] = input.DisplayName.Value
		}
		if input.Email.Present {
			body["email"] = input.Email.Value
		}
		if input.FlowId.Present {
			body["flowId"] = input.FlowId.Value
		}
	}

	req.SetBody(body)

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &userv1.Organization{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// OrganizationDeleteInput is the input param for the Delete method.
type OrganizationDeleteInput struct {
}

func (n *organizationsImpl) Delete(ctx context.Context, organizationId string, input *OrganizationDeleteInput) (*userv1.Organization, error) {
	req := internal.NewRequest(
		"user.organizations.delete",
		"DELETE",
		fmt.Sprintf("/user/v1/organizations/%s",
			url.PathEscape(organizationId),
		),
	)

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &userv1.Organization{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// OrganizationLeaveInput is the input param for the Leave method.
type OrganizationLeaveInput struct {
}

func (n *organizationsImpl) Leave(ctx context.Context, organizationId string, input *OrganizationLeaveInput) (*apiv1.EmptyResponse, error) {
	req := internal.NewRequest(
		"user.organizations.leave",
		"DELETE",
		fmt.Sprintf("/user/v1/organizations/%s:leave",
			url.PathEscape(organizationId),
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
