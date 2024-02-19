// Code generated. DO NOT EDIT.

package userapi

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/userhubdev/go-sdk/internal"
)

func TestOrganizations_List(t *testing.T) {
	t.Parallel()

	tr := &internal.TestTransport{}
	tr.Body = `{
  "organizations": [
    {
      "id": "string",
      "uniqueId": "test",
      "displayName": "Test",
      "email": "test@example.com",
      "emailVerified": true,
      "imageUrl": "https://example.com/test.png",
      "disabled": true
    }
  ],
  "nextPageToken": "string",
  "previousPageToken": "string"
}`

	n := &organizationsImpl{transport: tr}

	res, err := n.List(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, `GET`, tr.Request.Method())
	require.Equal(t, `/user/v1/organizations`, tr.Request.Path())

	res, err = n.List(context.Background(), &OrganizationListInput{})
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestOrganizations_Create(t *testing.T) {
	t.Parallel()

	tr := &internal.TestTransport{}
	tr.Body = `{
  "id": "string",
  "uniqueId": "test",
  "displayName": "Test",
  "email": "test@example.com",
  "emailVerified": true,
  "imageUrl": "https://example.com/test.png",
  "disabled": true
}`

	n := &organizationsImpl{transport: tr}

	res, err := n.Create(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, `POST`, tr.Request.Method())
	require.Equal(t, `/user/v1/organizations`, tr.Request.Path())

	res, err = n.Create(context.Background(), &OrganizationCreateInput{})
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestOrganizations_Get(t *testing.T) {
	t.Parallel()

	tr := &internal.TestTransport{}
	tr.Body = `{
  "id": "string",
  "uniqueId": "test",
  "displayName": "Test",
  "email": "test@example.com",
  "emailVerified": true,
  "imageUrl": "https://example.com/test.png",
  "disabled": true
}`

	n := &organizationsImpl{transport: tr}

	res, err := n.Get(context.Background(), "organizationId", nil)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, `GET`, tr.Request.Method())
	require.Equal(t, `/user/v1/organizations/organizationId`, tr.Request.Path())

	res, err = n.Get(context.Background(), "organizationId", &OrganizationGetInput{})
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestOrganizations_Update(t *testing.T) {
	t.Parallel()

	tr := &internal.TestTransport{}
	tr.Body = `{
  "id": "string",
  "uniqueId": "test",
  "displayName": "Test",
  "email": "test@example.com",
  "emailVerified": true,
  "imageUrl": "https://example.com/test.png",
  "disabled": true
}`

	n := &organizationsImpl{transport: tr}

	res, err := n.Update(context.Background(), "organizationId", nil)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, `PATCH`, tr.Request.Method())
	require.Equal(t, `/user/v1/organizations/organizationId`, tr.Request.Path())

	res, err = n.Update(context.Background(), "organizationId", &OrganizationUpdateInput{})
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestOrganizations_Delete(t *testing.T) {
	t.Parallel()

	tr := &internal.TestTransport{}
	tr.Body = `{
  "id": "string",
  "uniqueId": "test",
  "displayName": "Test",
  "email": "test@example.com",
  "emailVerified": true,
  "imageUrl": "https://example.com/test.png",
  "disabled": true
}`

	n := &organizationsImpl{transport: tr}

	res, err := n.Delete(context.Background(), "organizationId", nil)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, `DELETE`, tr.Request.Method())
	require.Equal(t, `/user/v1/organizations/organizationId`, tr.Request.Path())

	res, err = n.Delete(context.Background(), "organizationId", &OrganizationDeleteInput{})
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestOrganizations_Leave(t *testing.T) {
	t.Parallel()

	tr := &internal.TestTransport{}
	tr.Body = `{}`

	n := &organizationsImpl{transport: tr}

	res, err := n.Leave(context.Background(), "organizationId", nil)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, `DELETE`, tr.Request.Method())
	require.Equal(t, `/user/v1/organizations/organizationId:leave`, tr.Request.Path())

	res, err = n.Leave(context.Background(), "organizationId", &OrganizationLeaveInput{})
	require.NoError(t, err)
	require.NotNil(t, res)
}
