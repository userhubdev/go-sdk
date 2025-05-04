// Code generated. DO NOT EDIT.

package userapi

import (
	"context"
	"fmt"
	"net/url"

	"github.com/userhubdev/go-sdk/internal"
	"github.com/userhubdev/go-sdk/types"
	"github.com/userhubdev/go-sdk/userv1"
)

type Flows interface {
	// List flows.
	List(ctx context.Context, input *FlowListInput) (*userv1.ListFlowsResponse, error)
	// Create a new join organization flow.
	//
	// This invites a person to join an organization.
	CreateJoinOrganization(ctx context.Context, input *FlowCreateJoinOrganizationInput) (*userv1.Flow, error)
	// Update a join organization flow.
	UpdateJoinOrganization(ctx context.Context, flowId string, input *FlowUpdateJoinOrganizationInput) (*userv1.Flow, error)
	// Create a new signup flow.
	//
	// This invites a person to join the app.
	CreateSignup(ctx context.Context, input *FlowCreateSignupInput) (*userv1.Flow, error)
	// Get a flow.
	Get(ctx context.Context, flowId string, input *FlowGetInput) (*userv1.Flow, error)
	// Consume a flow.
	//
	// This accepts the flow (e.g. for a join organization flow it will
	// accept the invitation and add the member to the organization).
	Consume(ctx context.Context, flowId string, input *FlowConsumeInput) (*userv1.Flow, error)
	// Cancel a flow.
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
	// Whether to filter out flows not in the `START_PENDING` or `STARTED`
	// state.
	Active bool
	// Whether to only return flows created by the authenticated user.
	Creator bool
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
	// A comma-separated list of fields to order by.
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
		if !internal.IsEmpty(input.Active) {
			req.SetQuery("active", input.Active)
		}
		if !internal.IsEmpty(input.Creator) {
			req.SetQuery("creator", input.Creator)
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
	// The identifier of the role.
	RoleId string
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
		if !internal.IsEmpty(input.RoleId) {
			body["roleId"] = input.RoleId
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

// FlowUpdateJoinOrganizationInput is the input param for the UpdateJoinOrganization method.
type FlowUpdateJoinOrganizationInput struct {
	// The identifier of the role.
	RoleId types.Optional[string]
}

func (n *flowsImpl) UpdateJoinOrganization(ctx context.Context, flowId string, input *FlowUpdateJoinOrganizationInput) (*userv1.Flow, error) {
	req := internal.NewRequest(
		"user.flows.updateJoinOrganization",
		"PATCH",
		fmt.Sprintf("/user/v1/flows/%s:updateJoinOrganization",
			url.PathEscape(flowId),
		),
	)
	req.SetIdempotent(true)

	body := map[string]any{}

	if input != nil {
		if input.RoleId.Present {
			body["roleId"] = input.RoleId.Value
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

// FlowCreateSignupInput is the input param for the CreateSignup method.
type FlowCreateSignupInput struct {
	// The email address of the person to invite.
	Email string
	// The display name of the person to invite.
	DisplayName string
	// Whether to create an organization as part of the signup flow.
	CreateOrganization bool
}

func (n *flowsImpl) CreateSignup(ctx context.Context, input *FlowCreateSignupInput) (*userv1.Flow, error) {
	req := internal.NewRequest(
		"user.flows.createSignup",
		"POST",
		"/user/v1/flows:createSignup",
	)

	body := map[string]any{}

	if input != nil {
		if !internal.IsEmpty(input.Email) {
			body["email"] = input.Email
		}
		if !internal.IsEmpty(input.DisplayName) {
			body["displayName"] = input.DisplayName
		}
		if !internal.IsEmpty(input.CreateOrganization) {
			body["createOrganization"] = input.CreateOrganization
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
