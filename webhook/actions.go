// Code generated. DO NOT EDIT.

package webhook

import (
	"context"

	"github.com/userhubdev/go-sdk/connectionsv1"
	"github.com/userhubdev/go-sdk/eventsv1"
)

// OnChallenge registers a handler for the `challenge` action.
func (w *Webhook) OnChallenge(fn func(ctx context.Context, input *connectionsv1.Challenge) (*connectionsv1.Challenge, error)) *Webhook {
	var handler Handler
	if fn != nil {
		handler = genericHandler[connectionsv1.Challenge, connectionsv1.Challenge](fn)
	}
	return w.OnAction("challenge", handler)
}

// OnEvent registers a handler for the `events.handle` action.
func (w *Webhook) OnEvent(fn func(ctx context.Context, input *eventsv1.Event) error) *Webhook {
	var handler Handler
	if fn != nil {
		handler = genericInputOnlyHandler[eventsv1.Event](fn)
	}
	return w.OnAction("events.handle", handler)
}

// OnListUsers registers a handler for the `users.list` action.
func (w *Webhook) OnListUsers(fn func(ctx context.Context, input *connectionsv1.ListCustomUsersRequest) (*connectionsv1.ListCustomUsersResponse, error)) *Webhook {
	var handler Handler
	if fn != nil {
		handler = genericHandler[connectionsv1.ListCustomUsersRequest, connectionsv1.ListCustomUsersResponse](fn)
	}
	return w.OnAction("users.list", handler)
}

// OnCreateUser registers a handler for the `users.create` action.
func (w *Webhook) OnCreateUser(fn func(ctx context.Context, input *connectionsv1.CustomUser) (*connectionsv1.CustomUser, error)) *Webhook {
	var handler Handler
	if fn != nil {
		handler = genericHandler[connectionsv1.CustomUser, connectionsv1.CustomUser](fn)
	}
	return w.OnAction("users.create", handler)
}

// OnGetUser registers a handler for the `users.get` action.
func (w *Webhook) OnGetUser(fn func(ctx context.Context, input *connectionsv1.GetCustomUserRequest) (*connectionsv1.CustomUser, error)) *Webhook {
	var handler Handler
	if fn != nil {
		handler = genericHandler[connectionsv1.GetCustomUserRequest, connectionsv1.CustomUser](fn)
	}
	return w.OnAction("users.get", handler)
}

// OnUpdateUser registers a handler for the `users.update` action.
func (w *Webhook) OnUpdateUser(fn func(ctx context.Context, input *connectionsv1.CustomUser) (*connectionsv1.CustomUser, error)) *Webhook {
	var handler Handler
	if fn != nil {
		handler = genericHandler[connectionsv1.CustomUser, connectionsv1.CustomUser](fn)
	}
	return w.OnAction("users.update", handler)
}

// OnDeleteUser registers a handler for the `users.delete` action.
func (w *Webhook) OnDeleteUser(fn func(ctx context.Context, input *connectionsv1.DeleteCustomUserRequest) error) *Webhook {
	var handler Handler
	if fn != nil {
		handler = genericInputOnlyHandler[connectionsv1.DeleteCustomUserRequest](fn)
	}
	return w.OnAction("users.delete", handler)
}
