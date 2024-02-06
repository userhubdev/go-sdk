// Code generated. DO NOT EDIT.

package adminapi

import (
	"context"
	"fmt"
	"net/url"

	"github.com/userhubdev/go-sdk/adminv1"
	"github.com/userhubdev/go-sdk/internal"
)

type Subscriptions interface {
	// Lists subscriptions.
	List(ctx context.Context, input *SubscriptionListInput) (*adminv1.ListSubscriptionsResponse, error)
	// Retrieves specified subscription.
	Get(ctx context.Context, subscriptionId string, input *SubscriptionGetInput) (*adminv1.Subscription, error)
}

type subscriptionsImpl struct {
	transport internal.Transport
}

// SubscriptionListInput is the input param for the List method.
type SubscriptionListInput struct {
	// Filter results by organization identifier.
	//
	// This is required if user identifier is not specified.
	OrganizationId string
	// Filter results by user identifier.
	//
	// This is required if organization identifier is not specified.
	UserId string
	// The maximum number of subscriptions to return. The API may return fewer than
	// this value.
	//
	// If unspecified, at most 20 subscriptions will be returned.
	// The maximum value is 100; values above 100 will be coerced to 100.
	PageSize int32
	// A page token, received from a previous list subscriptions call.
	// Provide this to retrieve the subsequent page.
	//
	// When paginating, all other parameters provided to list subscriptions must match
	// the call that provided the page token.
	PageToken string
	// A comma-separated list of fields to order by, sorted in ascending order.
	// Use `desc` after a field name for descending.
	//
	// Supported fields:
	// - `createTime`
	OrderBy string
	// The Subscription view to return in the results.
	//
	// This defaults to the `BASIC` view.
	View string
}

func (n *subscriptionsImpl) List(ctx context.Context, input *SubscriptionListInput) (*adminv1.ListSubscriptionsResponse, error) {
	req := internal.NewRequest(
		"admin.subscriptions.list",
		"GET",
		"/admin/v1/subscriptions",
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

	model := &adminv1.ListSubscriptionsResponse{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// SubscriptionGetInput is the input param for the Get method.
type SubscriptionGetInput struct {
	// Restrict by organization identifier.
	OrganizationId string
	// Restrict by user identifier.
	UserId string
}

func (n *subscriptionsImpl) Get(ctx context.Context, subscriptionId string, input *SubscriptionGetInput) (*adminv1.Subscription, error) {
	req := internal.NewRequest(
		"admin.subscriptions.get",
		"GET",
		fmt.Sprintf("/admin/v1/subscriptions/%s",
			url.PathEscape(subscriptionId),
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

	model := &adminv1.Subscription{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}
