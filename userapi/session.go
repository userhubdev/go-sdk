// Code generated. DO NOT EDIT.

package userapi

import (
	"context"

	"github.com/userhubdev/go-sdk/internal"
	"github.com/userhubdev/go-sdk/userv1"
)

type Session interface {
	// Get current session details.
	Get(ctx context.Context, input *SessionGetInput) (*userv1.Session, error)
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
