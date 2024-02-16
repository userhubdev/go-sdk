// Code generated. DO NOT EDIT.

package userapi

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/userhubdev/go-sdk/internal"
)

func TestFlows_List(t *testing.T) {
	t.Parallel()

	tr := &internal.TestTransport{}
	tr.Body = `{
  "flows": [
    {
      "id": "string",
      "state": "START_PENDING",
      "stateReason": "DELETED",
      "type": "JOIN_ORGANIZATION",
      "expireTime": "2024-02-05T23:07:46.483Z",
      "createTime": "2024-02-05T23:07:46.483Z"
    }
  ],
  "nextPageToken": "string",
  "previousPageToken": "string"
}`

	n := &flowsImpl{transport: tr}

	res, err := n.List(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, `GET`, tr.Request.Method())
	require.Equal(t, `/user/v1/flows`, tr.Request.Path())

	res, err = n.List(context.Background(), &FlowListInput{})
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestFlows_CreateJoinOrganization(t *testing.T) {
	t.Parallel()

	tr := &internal.TestTransport{}
	tr.Body = `{
  "id": "string",
  "state": "START_PENDING",
  "stateReason": "DELETED",
  "type": "JOIN_ORGANIZATION",
  "organization": {
    "id": "string",
    "uniqueId": "test",
    "displayName": "Test",
    "email": "test@example.com",
    "emailVerified": true,
    "imageUrl": "https://example.com/test.png",
    "disabled": true
  },
  "user": {
    "id": "string",
    "uniqueId": "test",
    "displayName": "Test",
    "email": "test@example.com",
    "emailVerified": true,
    "imageUrl": "https://example.com/test.png",
    "disabled": true
  },
  "creator": {
    "id": "string",
    "uniqueId": "test",
    "displayName": "Test",
    "email": "test@example.com",
    "emailVerified": true,
    "imageUrl": "https://example.com/test.png",
    "disabled": true
  },
  "expireTime": "2024-02-05T23:07:46.483Z",
  "createTime": "2024-02-05T23:07:46.483Z",
  "joinOrganization": {
    "displayName": "Test",
    "email": "test@example.com"
  },
  "signup": {
    "email": "test@example.com",
    "displayName": "Test",
    "createOrganization": true
  }
}`

	n := &flowsImpl{transport: tr}

	res, err := n.CreateJoinOrganization(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, `POST`, tr.Request.Method())
	require.Equal(t, `/user/v1/flows:createJoinOrganization`, tr.Request.Path())

	res, err = n.CreateJoinOrganization(context.Background(), &FlowCreateJoinOrganizationInput{})
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestFlows_Get(t *testing.T) {
	t.Parallel()

	tr := &internal.TestTransport{}
	tr.Body = `{
  "id": "string",
  "state": "START_PENDING",
  "stateReason": "DELETED",
  "type": "JOIN_ORGANIZATION",
  "organization": {
    "id": "string",
    "uniqueId": "test",
    "displayName": "Test",
    "email": "test@example.com",
    "emailVerified": true,
    "imageUrl": "https://example.com/test.png",
    "disabled": true
  },
  "user": {
    "id": "string",
    "uniqueId": "test",
    "displayName": "Test",
    "email": "test@example.com",
    "emailVerified": true,
    "imageUrl": "https://example.com/test.png",
    "disabled": true
  },
  "creator": {
    "id": "string",
    "uniqueId": "test",
    "displayName": "Test",
    "email": "test@example.com",
    "emailVerified": true,
    "imageUrl": "https://example.com/test.png",
    "disabled": true
  },
  "expireTime": "2024-02-05T23:07:46.483Z",
  "createTime": "2024-02-05T23:07:46.483Z",
  "joinOrganization": {
    "displayName": "Test",
    "email": "test@example.com"
  },
  "signup": {
    "email": "test@example.com",
    "displayName": "Test",
    "createOrganization": true
  }
}`

	n := &flowsImpl{transport: tr}

	res, err := n.Get(context.Background(), "flowId", nil)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, `GET`, tr.Request.Method())
	require.Equal(t, `/user/v1/flows/flowId`, tr.Request.Path())

	res, err = n.Get(context.Background(), "flowId", &FlowGetInput{})
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestFlows_Consume(t *testing.T) {
	t.Parallel()

	tr := &internal.TestTransport{}
	tr.Body = `{
  "id": "string",
  "state": "START_PENDING",
  "stateReason": "DELETED",
  "type": "JOIN_ORGANIZATION",
  "organization": {
    "id": "string",
    "uniqueId": "test",
    "displayName": "Test",
    "email": "test@example.com",
    "emailVerified": true,
    "imageUrl": "https://example.com/test.png",
    "disabled": true
  },
  "user": {
    "id": "string",
    "uniqueId": "test",
    "displayName": "Test",
    "email": "test@example.com",
    "emailVerified": true,
    "imageUrl": "https://example.com/test.png",
    "disabled": true
  },
  "creator": {
    "id": "string",
    "uniqueId": "test",
    "displayName": "Test",
    "email": "test@example.com",
    "emailVerified": true,
    "imageUrl": "https://example.com/test.png",
    "disabled": true
  },
  "expireTime": "2024-02-05T23:07:46.483Z",
  "createTime": "2024-02-05T23:07:46.483Z",
  "joinOrganization": {
    "displayName": "Test",
    "email": "test@example.com"
  },
  "signup": {
    "email": "test@example.com",
    "displayName": "Test",
    "createOrganization": true
  }
}`

	n := &flowsImpl{transport: tr}

	res, err := n.Consume(context.Background(), "flowId", nil)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, `POST`, tr.Request.Method())
	require.Equal(t, `/user/v1/flows/flowId:consume`, tr.Request.Path())

	res, err = n.Consume(context.Background(), "flowId", &FlowConsumeInput{})
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestFlows_Cancel(t *testing.T) {
	t.Parallel()

	tr := &internal.TestTransport{}
	tr.Body = `{
  "id": "string",
  "state": "START_PENDING",
  "stateReason": "DELETED",
  "type": "JOIN_ORGANIZATION",
  "organization": {
    "id": "string",
    "uniqueId": "test",
    "displayName": "Test",
    "email": "test@example.com",
    "emailVerified": true,
    "imageUrl": "https://example.com/test.png",
    "disabled": true
  },
  "user": {
    "id": "string",
    "uniqueId": "test",
    "displayName": "Test",
    "email": "test@example.com",
    "emailVerified": true,
    "imageUrl": "https://example.com/test.png",
    "disabled": true
  },
  "creator": {
    "id": "string",
    "uniqueId": "test",
    "displayName": "Test",
    "email": "test@example.com",
    "emailVerified": true,
    "imageUrl": "https://example.com/test.png",
    "disabled": true
  },
  "expireTime": "2024-02-05T23:07:46.483Z",
  "createTime": "2024-02-05T23:07:46.483Z",
  "joinOrganization": {
    "displayName": "Test",
    "email": "test@example.com"
  },
  "signup": {
    "email": "test@example.com",
    "displayName": "Test",
    "createOrganization": true
  }
}`

	n := &flowsImpl{transport: tr}

	res, err := n.Cancel(context.Background(), "flowId", nil)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, `POST`, tr.Request.Method())
	require.Equal(t, `/user/v1/flows/flowId:cancel`, tr.Request.Path())

	res, err = n.Cancel(context.Background(), "flowId", &FlowCancelInput{})
	require.NoError(t, err)
	require.NotNil(t, res)
}
