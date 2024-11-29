// Code generated. DO NOT EDIT.

package userapi

import (
	"context"

	"github.com/userhubdev/go-sdk/internal"
	"github.com/userhubdev/go-sdk/userv1"
)

type Session interface {
	// Get the current session details.
	Get(ctx context.Context, input *SessionGetInput) (*userv1.Session, error)
	// Exchange an ID token from your IdP for an access token.
	ExchangeToken(ctx context.Context, input *SessionExchangeTokenInput) (*userv1.ExchangeSessionTokenResponse, error)
	// Create Portal session.
	CreatePortal(ctx context.Context, input *SessionCreatePortalInput) (*userv1.CreatePortalSessionResponse, error)
}

type sessionImpl struct {
	transport internal.Transport
}

// SessionGetInput is the input param for the Get method.
type SessionGetInput struct {
}

func (n *sessionImpl) Get(ctx context.Context, input *SessionGetInput) (*userv1.Session, error) {
	req := internal.NewRequest(
		"user.session.get",
		"GET",
		"/user/v1/session",
	)
	req.SetIdempotent(true)

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &userv1.Session{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// SessionExchangeTokenInput is the input param for the ExchangeToken method.
type SessionExchangeTokenInput struct {
	// The IdP ID token which is used to authenticated the user.
	Token string
}

func (n *sessionImpl) ExchangeToken(ctx context.Context, input *SessionExchangeTokenInput) (*userv1.ExchangeSessionTokenResponse, error) {
	req := internal.NewRequest(
		"user.session.exchangeToken",
		"POST",
		"/user/v1/session:exchangeToken",
	)

	body := map[string]any{}

	if input != nil {
		if !internal.IsEmpty(input.Token) {
			body["token"] = input.Token
		}
	}

	req.SetBody(body)

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &userv1.ExchangeSessionTokenResponse{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// SessionCreatePortalInput is the input param for the CreatePortal method.
type SessionCreatePortalInput struct {
	// The identifier of the organization.
	//
	// When specified the `{accountId}` in the `portalUrl` will be
	// replaced with the organization ID, otherwise the user ID
	// will be used.
	OrganizationId string
	// The Portal URL, this is the target URL on the portal site.
	//
	// If not defined the root URL for the portal will be used.
	//
	// This does not need to be the full URL, you have the option
	// of passing in a path instead (e.g. `/`).
	//
	// You also have the option of including the `{accountId}`
	// string in the path/URL which will be replaced with either the
	// UserHub user ID (if `organizationId` is not specified)
	// or the UserHub organization ID (if specified).
	//
	// Examples:
	// * `/{accountId}` - the billing dashboard
	// * `/{accountId}/checkout` - start a checkout
	// * `/{accountId}/checkout/<some-plan-id>` - start a checkout with a specified plan
	// * `/{accountId}/cancel` - cancel current plan
	// * `/{accountId}/members` - manage organization members
	// * `/{accountId}/invite` - invite a user to an organization
	PortalUrl string
	// The URL the user should be sent to when they want to return to
	// the app (e.g. cancel checkout).
	//
	// If not defined the app URL will be used.
	ReturnUrl string
	// The URL the user should be sent after they successfully complete
	// an action (e.g. checkout).
	//
	// If not defined the return URL will be used.
	SuccessUrl string
}

func (n *sessionImpl) CreatePortal(ctx context.Context, input *SessionCreatePortalInput) (*userv1.CreatePortalSessionResponse, error) {
	req := internal.NewRequest(
		"user.session.createPortal",
		"POST",
		"/user/v1/session:createPortal",
	)

	body := map[string]any{}

	if input != nil {
		if !internal.IsEmpty(input.OrganizationId) {
			body["organizationId"] = input.OrganizationId
		}
		if !internal.IsEmpty(input.PortalUrl) {
			body["portalUrl"] = input.PortalUrl
		}
		if !internal.IsEmpty(input.ReturnUrl) {
			body["returnUrl"] = input.ReturnUrl
		}
		if !internal.IsEmpty(input.SuccessUrl) {
			body["successUrl"] = input.SuccessUrl
		}
	}

	req.SetBody(body)

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &userv1.CreatePortalSessionResponse{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}
