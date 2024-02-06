// Code generated. DO NOT EDIT.

package adminapi

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/userhubdev/go-sdk/adminv1"
	"github.com/userhubdev/go-sdk/internal"
)

type Flows interface {
	// Lists flows.
	List(ctx context.Context, input *FlowListInput) (*adminv1.ListFlowsResponse, error)
	// Create a join organization flow.
	//
	// This invites a person to join an organization.
	CreateJoinOrganization(ctx context.Context, input *FlowCreateJoinOrganizationInput) (*adminv1.Flow, error)
	// Retrieves specified flow.
	Get(ctx context.Context, flowId string, input *FlowGetInput) (*adminv1.Flow, error)
	// Cancels specified flow.
	Cancel(ctx context.Context, flowId string, input *FlowCancelInput) (*adminv1.Flow, error)
}

type flowsImpl struct {
	transport internal.Transport
}

// FlowListInput is the input param for the List method.
type FlowListInput struct {
	// Filter results by organization identifier.
	OrganizationId string
	// Filter results by user identifier.
	UserId string
	// Filter the results by the specified flow type.
	Type string
	// The maximum number of flows to return. The API may return fewer than
	// this value.
	//
	// If unspecified, at most 20 flows will be returned.
	// The maximum value is 100; values above 100 will be coerced to 100.
	PageSize int32
	// A page token, received from a previous list flows call.
	// Provide this to retrieve the subsequent page.
	//
	// When paginating, all other parameters provided to list flows must match
	// the call that provided the page token.
	PageToken string
	// A comma-separated list of fields to order by, sorted in ascending order.
	// Use `desc` after a field name for descending.
	//
	// Supported fields:
	// - `type`
	// - `createTime`
	OrderBy string
	// The Flow view to return in the results.
	//
	// This defaults to the `BASIC` view.
	View string
}

func (n *flowsImpl) List(ctx context.Context, input *FlowListInput) (*adminv1.ListFlowsResponse, error) {
	req := internal.NewRequest(
		"admin.flows.list",
		"GET",
		"/admin/v1/flows",
	)
	req.SetIdempotent(true)

	if input != nil {
		if !internal.IsEmpty(input.OrganizationId) {
			req.SetQuery("organizationId", input.OrganizationId)
		}
		if !internal.IsEmpty(input.UserId) {
			req.SetQuery("userId", input.UserId)
		}
		if !internal.IsEmpty(input.Type) {
			req.SetQuery("type", input.Type)
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

	model := &adminv1.ListFlowsResponse{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// FlowCreateJoinOrganizationInput is the input param for the CreateJoinOrganization method.
type FlowCreateJoinOrganizationInput struct {
	// The identifier of the organization.
	OrganizationId string
	// The identifier of the user.
	//
	// This is required if email is not specified.
	UserId string
	// The email address of the person to invite.
	//
	// This is required if user is not specified or the user
	// does not have an email address.
	Email string
	// The display name of the person to invite.
	DisplayName string
	// The identifier of the user sending the invite.
	CreatorUserId string
	// The time the flow will expire.
	//
	// This field is not allowed if `ttl` is specified.
	ExpireTime time.Time
	// The amount of time a flow will be available (e.g. `86400s`).
	//
	// This must be a string with the number of seconds followed by a
	// trailing `s`.
	//
	// This field is not allowed if `expireTime` is specified.
	Ttl string
}

func (n *flowsImpl) CreateJoinOrganization(ctx context.Context, input *FlowCreateJoinOrganizationInput) (*adminv1.Flow, error) {
	req := internal.NewRequest(
		"admin.flows.createJoinOrganization",
		"POST",
		"/admin/v1/flows:createJoinOrganization",
	)

	body := map[string]any{}

	if input != nil {
		if !internal.IsEmpty(input.OrganizationId) {
			body["organizationId"] = input.OrganizationId
		}
		if !internal.IsEmpty(input.UserId) {
			body["userId"] = input.UserId
		}
		if !internal.IsEmpty(input.Email) {
			body["email"] = input.Email
		}
		if !internal.IsEmpty(input.DisplayName) {
			body["displayName"] = input.DisplayName
		}
		if !internal.IsEmpty(input.CreatorUserId) {
			body["creatorUserId"] = input.CreatorUserId
		}
		if !internal.IsEmpty(input.ExpireTime) {
			body["expireTime"] = input.ExpireTime
		}
		if !internal.IsEmpty(input.Ttl) {
			body["ttl"] = input.Ttl
		}
	}

	req.SetBody(body)

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &adminv1.Flow{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// FlowGetInput is the input param for the Get method.
type FlowGetInput struct {
}

func (n *flowsImpl) Get(ctx context.Context, flowId string, input *FlowGetInput) (*adminv1.Flow, error) {
	req := internal.NewRequest(
		"admin.flows.get",
		"GET",
		fmt.Sprintf("/admin/v1/flows/%s",
			url.PathEscape(flowId),
		),
	)
	req.SetIdempotent(true)

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &adminv1.Flow{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// FlowCancelInput is the input param for the Cancel method.
type FlowCancelInput struct {
}

func (n *flowsImpl) Cancel(ctx context.Context, flowId string, input *FlowCancelInput) (*adminv1.Flow, error) {
	req := internal.NewRequest(
		"admin.flows.cancel",
		"POST",
		fmt.Sprintf("/admin/v1/flows/%s:cancel",
			url.PathEscape(flowId),
		),
	)
	req.SetIdempotent(true)

	body := map[string]any{}

	req.SetBody(body)

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &adminv1.Flow{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}