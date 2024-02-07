// Code generated. DO NOT EDIT.

package userapi

import (
	"context"
	"fmt"
	"net/url"

	"github.com/userhubdev/go-sdk/internal"
	"github.com/userhubdev/go-sdk/userv1"
)

type Flows interface {
	// Lists flows.
	List(ctx context.Context, input *FlowListInput) (*userv1.ListFlowsResponse, error)
	// Creates a join organization flow.
	//
	// This invites a person to join an organization.
	CreateJoinOrganization(ctx context.Context, input *FlowCreateJoinOrganizationInput) (*userv1.Flow, error)
	// Retrieves specified flow.
	Get(ctx context.Context, flowId string, input *FlowGetInput) (*userv1.Flow, error)
	// Consume the flow.
	//
	// This accepts the flow (e.g. for a join organization flow it will
	// accept the invitation and add the member to the organization).
	Consume(ctx context.Context, flowId string, input *FlowConsumeInput) (*userv1.Flow, error)
	// Cancels specified flow.
	//
	// This cancels the flow and hides it from the user.
	Cancel(ctx context.Context, flowId string, input *FlowCancelInput) (*userv1.Flow, error)
}

type flowsImpl struct {
	transport internal.Transport
}

// FlowListInput is the input param for the List method.
type FlowListInput struct {
	// The identifier of the organization.
	//
	// When not set the user's flows are returned.
	//
	// Otherwise if the user is an admin of the provided organization then
	// the flows associated with that organization are returned.
	OrganizationId string
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
}

func (n *flowsImpl) List(ctx context.Context, input *FlowListInput) (*userv1.ListFlowsResponse, error) {
	req := internal.NewRequest(
		"user.flows.list",
		"GET",
		"/user/v1/flows",
	)
	req.SetIdempotent(true)

	if input != nil {
		if !internal.IsEmpty(input.OrganizationId) {
			req.SetQuery("organizationId", input.OrganizationId)
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
	}

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &userv1.ListFlowsResponse{}

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
}

func (n *flowsImpl) CreateJoinOrganization(ctx context.Context, input *FlowCreateJoinOrganizationInput) (*userv1.Flow, error) {
	req := internal.NewRequest(
		"user.flows.createJoinOrganization",
		"POST",
		"/user/v1/flows:createJoinOrganization",
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
	}

	req.SetBody(body)

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &userv1.Flow{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// FlowGetInput is the input param for the Get method.
type FlowGetInput struct {
}

func (n *flowsImpl) Get(ctx context.Context, flowId string, input *FlowGetInput) (*userv1.Flow, error) {
	req := internal.NewRequest(
		"user.flows.get",
		"GET",
		fmt.Sprintf("/user/v1/flows/%s",
			url.PathEscape(flowId),
		),
	)
	req.SetIdempotent(true)

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &userv1.Flow{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// FlowConsumeInput is the input param for the Consume method.
type FlowConsumeInput struct {
}

func (n *flowsImpl) Consume(ctx context.Context, flowId string, input *FlowConsumeInput) (*userv1.Flow, error) {
	req := internal.NewRequest(
		"user.flows.consume",
		"POST",
		fmt.Sprintf("/user/v1/flows/%s:consume",
			url.PathEscape(flowId),
		),
	)

	body := map[string]any{}

	req.SetBody(body)

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &userv1.Flow{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// FlowCancelInput is the input param for the Cancel method.
type FlowCancelInput struct {
}

func (n *flowsImpl) Cancel(ctx context.Context, flowId string, input *FlowCancelInput) (*userv1.Flow, error) {
	req := internal.NewRequest(
		"user.flows.cancel",
		"POST",
		fmt.Sprintf("/user/v1/flows/%s:cancel",
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

	model := &userv1.Flow{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}
