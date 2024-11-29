// Code generated. DO NOT EDIT.

package userapi

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/userhubdev/go-sdk/internal"
)

func TestSession_Get(t *testing.T) {
	t.Parallel()

	tr := &internal.TestTransport{}
	tr.Body = `{
  "user": {
    "id": "string",
    "uniqueId": "test",
    "displayName": "Test",
    "email": "test@example.com",
    "emailVerified": true,
    "imageUrl": "https://example.com/test.png",
    "disabled": true
  },
  "memberships": [
    {}
  ],
  "subscription": {
    "id": "string",
    "state": "ACTIVE",
    "anchorTime": "2024-02-05T23:07:46.483Z",
    "plan": {
      "id": "string",
      "displayName": "Test"
    },
    "seat": {}
  },
  "expireTime": "2024-02-05T23:07:46.483Z",
  "scopes": [
    "string"
  ]
}`

	n := &sessionImpl{transport: tr}

	res, err := n.Get(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, `GET`, tr.Request.Method())
	require.Equal(t, `/user/v1/session`, tr.Request.Path())

	res, err = n.Get(context.Background(), &SessionGetInput{})
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestSession_ExchangeToken(t *testing.T) {
	t.Parallel()

	tr := &internal.TestTransport{}
	tr.Body = `{
  "accessToken": "string",
  "expireTime": "2024-02-05T23:07:46.483Z"
}`

	n := &sessionImpl{transport: tr}

	res, err := n.ExchangeToken(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, `POST`, tr.Request.Method())
	require.Equal(t, `/user/v1/session:exchangeToken`, tr.Request.Path())

	res, err = n.ExchangeToken(context.Background(), &SessionExchangeTokenInput{})
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestSession_CreatePortal(t *testing.T) {
	t.Parallel()

	tr := &internal.TestTransport{}
	tr.Body = `{
  "redirectUrl": "https://example.com"
}`

	n := &sessionImpl{transport: tr}

	res, err := n.CreatePortal(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, `POST`, tr.Request.Method())
	require.Equal(t, `/user/v1/session:createPortal`, tr.Request.Path())

	res, err = n.CreatePortal(context.Background(), &SessionCreatePortalInput{})
	require.NoError(t, err)
	require.NotNil(t, res)
}
