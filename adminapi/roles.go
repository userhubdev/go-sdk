// Code generated. DO NOT EDIT.

package adminapi

import (
	"context"

	"github.com/userhubdev/go-sdk/adminv1"
	"github.com/userhubdev/go-sdk/internal"
)

type Roles interface {
	// List roles.
	List(ctx context.Context, input *RoleListInput) (*adminv1.ListRolesResponse, error)
}

type rolesImpl struct {
	transport internal.Transport
}

// RoleListInput is the input param for the List method.
type RoleListInput struct {
	// The maximum number of roles to return. The API may return fewer than
	// this value.
	//
	// If unspecified, at most 20 roles will be returned.
	// The maximum value is 100; values above 100 will be coerced to 100.
	PageSize int32
	// A page token, received from a previous list roles call.
	// Provide this to retrieve the subsequent page.
	//
	// When paginating, all other parameters provided to list roles must match
	// the call that provided the page token.
	PageToken string
	// A comma-separated list of fields to order by.
	OrderBy string
}

func (n *rolesImpl) List(ctx context.Context, input *RoleListInput) (*adminv1.ListRolesResponse, error) {
	req := internal.NewRequest(
		"admin.roles.list",
		"GET",
		"/admin/v1/roles",
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

	model := &adminv1.ListRolesResponse{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}
