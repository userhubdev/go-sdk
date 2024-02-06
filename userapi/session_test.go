// Code generated. DO NOT EDIT.

package userapi

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/userhubdev/go-sdk/internal"
)

func TestSession_Get(t *testing.T) {
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
