// Code generated. DO NOT EDIT.

package adminapi

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
      "state": "ACTIVE",
      "stateReason": "DELETED",
      "uniqueId": "test",
      "displayName": "Test",
      "email": "test@example.com",
      "emailVerified": true,
      "phoneNumber": "+12125550123",
      "phoneNumberVerified": true,
      "imageUrl": "https://example.com/test.png",
      "currencyCode": "USD",
      "languageCode": "en",
      "regionCode": "US",
      "timeZone": "America/New_York",
      "signupTime": "2024-02-05T23:07:46.483Z",
      "memberCount": 1,
      "disabled": true,
      "createTime": "2024-02-05T23:07:46.483Z",
      "updateTime": "2024-02-05T23:07:46.483Z"
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
	require.Equal(t, `/admin/v1/organizations`, tr.Request.Path())

	res, err = n.List(context.Background(), &OrganizationListInput{})
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestOrganizations_Create(t *testing.T) {
	t.Parallel()

	tr := &internal.TestTransport{}
	tr.Body = `{
  "id": "string",
  "state": "ACTIVE",
  "stateReason": "DELETED",
  "uniqueId": "test",
  "displayName": "Test",
  "email": "test@example.com",
  "emailVerified": true,
  "phoneNumber": "+12125550123",
  "phoneNumberVerified": true,
  "imageUrl": "https://example.com/test.png",
  "currencyCode": "USD",
  "languageCode": "en",
  "regionCode": "US",
  "timeZone": "America/New_York",
  "address": {
    "lines": [
      "string"
    ],
    "city": "Brooklyn",
    "state": "string",
    "postalCode": "11222",
    "country": "US"
  },
  "accountConnections": [
    {
      "externalId": "string",
      "adminUrl": "https://example.com",
      "state": "ACTIVE",
      "stateReason": "DELETED",
      "balanceAmount": "string",
      "currencyCode": "USD",
      "pullTime": "2024-02-05T23:07:46.483Z",
      "pushTime": "2024-02-05T23:07:46.483Z",
      "createTime": "2024-02-05T23:07:46.483Z",
      "updateTime": "2024-02-05T23:07:46.483Z"
    }
  ],
  "subscription": {
    "id": "string",
    "state": "ACTIVE",
    "anchorTime": "2024-02-05T23:07:46.483Z",
    "plan": {
      "id": "string",
      "displayName": "Test"
    }
  },
  "signupTime": "2024-02-05T23:07:46.483Z",
  "memberCount": 1,
  "disabled": true,
  "createTime": "2024-02-05T23:07:46.483Z",
  "updateTime": "2024-02-05T23:07:46.483Z"
}`

	n := &organizationsImpl{transport: tr}

	res, err := n.Create(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, `POST`, tr.Request.Method())
	require.Equal(t, `/admin/v1/organizations`, tr.Request.Path())

	res, err = n.Create(context.Background(), &OrganizationCreateInput{})
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestOrganizations_Get(t *testing.T) {
	t.Parallel()

	tr := &internal.TestTransport{}
	tr.Body = `{
  "id": "string",
  "state": "ACTIVE",
  "stateReason": "DELETED",
  "uniqueId": "test",
  "displayName": "Test",
  "email": "test@example.com",
  "emailVerified": true,
  "phoneNumber": "+12125550123",
  "phoneNumberVerified": true,
  "imageUrl": "https://example.com/test.png",
  "currencyCode": "USD",
  "languageCode": "en",
  "regionCode": "US",
  "timeZone": "America/New_York",
  "address": {
    "lines": [
      "string"
    ],
    "city": "Brooklyn",
    "state": "string",
    "postalCode": "11222",
    "country": "US"
  },
  "accountConnections": [
    {
      "externalId": "string",
      "adminUrl": "https://example.com",
      "state": "ACTIVE",
      "stateReason": "DELETED",
      "balanceAmount": "string",
      "currencyCode": "USD",
      "pullTime": "2024-02-05T23:07:46.483Z",
      "pushTime": "2024-02-05T23:07:46.483Z",
      "createTime": "2024-02-05T23:07:46.483Z",
      "updateTime": "2024-02-05T23:07:46.483Z"
    }
  ],
  "subscription": {
    "id": "string",
    "state": "ACTIVE",
    "anchorTime": "2024-02-05T23:07:46.483Z",
    "plan": {
      "id": "string",
      "displayName": "Test"
    }
  },
  "signupTime": "2024-02-05T23:07:46.483Z",
  "memberCount": 1,
  "disabled": true,
  "createTime": "2024-02-05T23:07:46.483Z",
  "updateTime": "2024-02-05T23:07:46.483Z"
}`

	n := &organizationsImpl{transport: tr}

	res, err := n.Get(context.Background(), "organizationId", nil)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, `GET`, tr.Request.Method())
	require.Equal(t, `/admin/v1/organizations/organizationId`, tr.Request.Path())

	res, err = n.Get(context.Background(), "organizationId", &OrganizationGetInput{})
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestOrganizations_Update(t *testing.T) {
	t.Parallel()

	tr := &internal.TestTransport{}
	tr.Body = `{
  "id": "string",
  "state": "ACTIVE",
  "stateReason": "DELETED",
  "uniqueId": "test",
  "displayName": "Test",
  "email": "test@example.com",
  "emailVerified": true,
  "phoneNumber": "+12125550123",
  "phoneNumberVerified": true,
  "imageUrl": "https://example.com/test.png",
  "currencyCode": "USD",
  "languageCode": "en",
  "regionCode": "US",
  "timeZone": "America/New_York",
  "address": {
    "lines": [
      "string"
    ],
    "city": "Brooklyn",
    "state": "string",
    "postalCode": "11222",
    "country": "US"
  },
  "accountConnections": [
    {
      "externalId": "string",
      "adminUrl": "https://example.com",
      "state": "ACTIVE",
      "stateReason": "DELETED",
      "balanceAmount": "string",
      "currencyCode": "USD",
      "pullTime": "2024-02-05T23:07:46.483Z",
      "pushTime": "2024-02-05T23:07:46.483Z",
      "createTime": "2024-02-05T23:07:46.483Z",
      "updateTime": "2024-02-05T23:07:46.483Z"
    }
  ],
  "subscription": {
    "id": "string",
    "state": "ACTIVE",
    "anchorTime": "2024-02-05T23:07:46.483Z",
    "plan": {
      "id": "string",
      "displayName": "Test"
    }
  },
  "signupTime": "2024-02-05T23:07:46.483Z",
  "memberCount": 1,
  "disabled": true,
  "createTime": "2024-02-05T23:07:46.483Z",
  "updateTime": "2024-02-05T23:07:46.483Z"
}`

	n := &organizationsImpl{transport: tr}

	res, err := n.Update(context.Background(), "organizationId", nil)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, `PATCH`, tr.Request.Method())
	require.Equal(t, `/admin/v1/organizations/organizationId`, tr.Request.Path())

	res, err = n.Update(context.Background(), "organizationId", &OrganizationUpdateInput{})
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestOrganizations_Delete(t *testing.T) {
	t.Parallel()

	tr := &internal.TestTransport{}
	tr.Body = `{
  "id": "string",
  "state": "ACTIVE",
  "stateReason": "DELETED",
  "uniqueId": "test",
  "displayName": "Test",
  "email": "test@example.com",
  "emailVerified": true,
  "phoneNumber": "+12125550123",
  "phoneNumberVerified": true,
  "imageUrl": "https://example.com/test.png",
  "currencyCode": "USD",
  "languageCode": "en",
  "regionCode": "US",
  "timeZone": "America/New_York",
  "address": {
    "lines": [
      "string"
    ],
    "city": "Brooklyn",
    "state": "string",
    "postalCode": "11222",
    "country": "US"
  },
  "accountConnections": [
    {
      "externalId": "string",
      "adminUrl": "https://example.com",
      "state": "ACTIVE",
      "stateReason": "DELETED",
      "balanceAmount": "string",
      "currencyCode": "USD",
      "pullTime": "2024-02-05T23:07:46.483Z",
      "pushTime": "2024-02-05T23:07:46.483Z",
      "createTime": "2024-02-05T23:07:46.483Z",
      "updateTime": "2024-02-05T23:07:46.483Z"
    }
  ],
  "subscription": {
    "id": "string",
    "state": "ACTIVE",
    "anchorTime": "2024-02-05T23:07:46.483Z",
    "plan": {
      "id": "string",
      "displayName": "Test"
    }
  },
  "signupTime": "2024-02-05T23:07:46.483Z",
  "memberCount": 1,
  "disabled": true,
  "createTime": "2024-02-05T23:07:46.483Z",
  "updateTime": "2024-02-05T23:07:46.483Z"
}`

	n := &organizationsImpl{transport: tr}

	res, err := n.Delete(context.Background(), "organizationId", nil)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, `DELETE`, tr.Request.Method())
	require.Equal(t, `/admin/v1/organizations/organizationId`, tr.Request.Path())

	res, err = n.Delete(context.Background(), "organizationId", &OrganizationDeleteInput{})
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestOrganizations_Undelete(t *testing.T) {
	t.Parallel()

	tr := &internal.TestTransport{}
	tr.Body = `{
  "id": "string",
  "state": "ACTIVE",
  "stateReason": "DELETED",
  "uniqueId": "test",
  "displayName": "Test",
  "email": "test@example.com",
  "emailVerified": true,
  "phoneNumber": "+12125550123",
  "phoneNumberVerified": true,
  "imageUrl": "https://example.com/test.png",
  "currencyCode": "USD",
  "languageCode": "en",
  "regionCode": "US",
  "timeZone": "America/New_York",
  "address": {
    "lines": [
      "string"
    ],
    "city": "Brooklyn",
    "state": "string",
    "postalCode": "11222",
    "country": "US"
  },
  "accountConnections": [
    {
      "externalId": "string",
      "adminUrl": "https://example.com",
      "state": "ACTIVE",
      "stateReason": "DELETED",
      "balanceAmount": "string",
      "currencyCode": "USD",
      "pullTime": "2024-02-05T23:07:46.483Z",
      "pushTime": "2024-02-05T23:07:46.483Z",
      "createTime": "2024-02-05T23:07:46.483Z",
      "updateTime": "2024-02-05T23:07:46.483Z"
    }
  ],
  "subscription": {
    "id": "string",
    "state": "ACTIVE",
    "anchorTime": "2024-02-05T23:07:46.483Z",
    "plan": {
      "id": "string",
      "displayName": "Test"
    }
  },
  "signupTime": "2024-02-05T23:07:46.483Z",
  "memberCount": 1,
  "disabled": true,
  "createTime": "2024-02-05T23:07:46.483Z",
  "updateTime": "2024-02-05T23:07:46.483Z"
}`

	n := &organizationsImpl{transport: tr}

	res, err := n.Undelete(context.Background(), "organizationId", nil)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, `POST`, tr.Request.Method())
	require.Equal(t, `/admin/v1/organizations/organizationId:undelete`, tr.Request.Path())

	res, err = n.Undelete(context.Background(), "organizationId", &OrganizationUndeleteInput{})
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestOrganizations_Connect(t *testing.T) {
	t.Parallel()

	tr := &internal.TestTransport{}
	tr.Body = `{
  "id": "string",
  "state": "ACTIVE",
  "stateReason": "DELETED",
  "uniqueId": "test",
  "displayName": "Test",
  "email": "test@example.com",
  "emailVerified": true,
  "phoneNumber": "+12125550123",
  "phoneNumberVerified": true,
  "imageUrl": "https://example.com/test.png",
  "currencyCode": "USD",
  "languageCode": "en",
  "regionCode": "US",
  "timeZone": "America/New_York",
  "address": {
    "lines": [
      "string"
    ],
    "city": "Brooklyn",
    "state": "string",
    "postalCode": "11222",
    "country": "US"
  },
  "accountConnections": [
    {
      "externalId": "string",
      "adminUrl": "https://example.com",
      "state": "ACTIVE",
      "stateReason": "DELETED",
      "balanceAmount": "string",
      "currencyCode": "USD",
      "pullTime": "2024-02-05T23:07:46.483Z",
      "pushTime": "2024-02-05T23:07:46.483Z",
      "createTime": "2024-02-05T23:07:46.483Z",
      "updateTime": "2024-02-05T23:07:46.483Z"
    }
  ],
  "subscription": {
    "id": "string",
    "state": "ACTIVE",
    "anchorTime": "2024-02-05T23:07:46.483Z",
    "plan": {
      "id": "string",
      "displayName": "Test"
    }
  },
  "signupTime": "2024-02-05T23:07:46.483Z",
  "memberCount": 1,
  "disabled": true,
  "createTime": "2024-02-05T23:07:46.483Z",
  "updateTime": "2024-02-05T23:07:46.483Z"
}`

	n := &organizationsImpl{transport: tr}

	res, err := n.Connect(context.Background(), "organizationId", nil)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, `POST`, tr.Request.Method())
	require.Equal(t, `/admin/v1/organizations/organizationId:connect`, tr.Request.Path())

	res, err = n.Connect(context.Background(), "organizationId", &OrganizationConnectInput{})
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestOrganizations_Disconnect(t *testing.T) {
	t.Parallel()

	tr := &internal.TestTransport{}
	tr.Body = `{
  "id": "string",
  "state": "ACTIVE",
  "stateReason": "DELETED",
  "uniqueId": "test",
  "displayName": "Test",
  "email": "test@example.com",
  "emailVerified": true,
  "phoneNumber": "+12125550123",
  "phoneNumberVerified": true,
  "imageUrl": "https://example.com/test.png",
  "currencyCode": "USD",
  "languageCode": "en",
  "regionCode": "US",
  "timeZone": "America/New_York",
  "address": {
    "lines": [
      "string"
    ],
    "city": "Brooklyn",
    "state": "string",
    "postalCode": "11222",
    "country": "US"
  },
  "accountConnections": [
    {
      "externalId": "string",
      "adminUrl": "https://example.com",
      "state": "ACTIVE",
      "stateReason": "DELETED",
      "balanceAmount": "string",
      "currencyCode": "USD",
      "pullTime": "2024-02-05T23:07:46.483Z",
      "pushTime": "2024-02-05T23:07:46.483Z",
      "createTime": "2024-02-05T23:07:46.483Z",
      "updateTime": "2024-02-05T23:07:46.483Z"
    }
  ],
  "subscription": {
    "id": "string",
    "state": "ACTIVE",
    "anchorTime": "2024-02-05T23:07:46.483Z",
    "plan": {
      "id": "string",
      "displayName": "Test"
    }
  },
  "signupTime": "2024-02-05T23:07:46.483Z",
  "memberCount": 1,
  "disabled": true,
  "createTime": "2024-02-05T23:07:46.483Z",
  "updateTime": "2024-02-05T23:07:46.483Z"
}`

	n := &organizationsImpl{transport: tr}

	res, err := n.Disconnect(context.Background(), "organizationId", nil)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, `POST`, tr.Request.Method())
	require.Equal(t, `/admin/v1/organizations/organizationId:disconnect`, tr.Request.Path())

	res, err = n.Disconnect(context.Background(), "organizationId", &OrganizationDisconnectInput{})
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestOrganizations_ListMembers(t *testing.T) {
	t.Parallel()

	tr := &internal.TestTransport{}
	tr.Body = `{
  "members": [
    {
      "state": "ACTIVE",
      "createTime": "2024-02-05T23:07:46.483Z",
      "updateTime": "2024-02-05T23:07:46.483Z"
    }
  ],
  "nextPageToken": "string",
  "previousPageToken": "string"
}`

	n := &organizationsImpl{transport: tr}

	res, err := n.ListMembers(context.Background(), "organizationId", nil)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, `GET`, tr.Request.Method())
	require.Equal(t, `/admin/v1/organizations/organizationId/members`, tr.Request.Path())

	res, err = n.ListMembers(context.Background(), "organizationId", &OrganizationListMembersInput{})
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestOrganizations_AddMember(t *testing.T) {
	t.Parallel()

	tr := &internal.TestTransport{}
	tr.Body = `{
  "state": "ACTIVE",
  "user": {
    "id": "string",
    "state": "ACTIVE",
    "stateReason": "DELETED",
    "uniqueId": "test",
    "displayName": "Test",
    "email": "test@example.com",
    "emailVerified": true,
    "phoneNumber": "+12125550123",
    "phoneNumberVerified": true,
    "imageUrl": "https://example.com/test.png",
    "currencyCode": "USD",
    "languageCode": "en",
    "regionCode": "US",
    "timeZone": "America/New_York",
    "address": {
      "lines": [],
      "city": "Brooklyn",
      "state": "string",
      "postalCode": "11222",
      "country": "US"
    },
    "accountConnections": [],
    "subscription": {
      "id": "string",
      "state": "ACTIVE",
      "anchorTime": "2024-02-05T23:07:46.483Z"
    },
    "memberships": [],
    "signupTime": "2024-02-05T23:07:46.483Z",
    "disabled": true,
    "createTime": "2024-02-05T23:07:46.483Z",
    "updateTime": "2024-02-05T23:07:46.483Z"
  },
  "role": {
    "id": "string",
    "uniqueId": "test",
    "displayName": "Test",
    "type": "OWNER",
    "description": "string",
    "permissionSets": [
      "string"
    ],
    "default": true,
    "archived": true,
    "createTime": "2024-02-05T23:07:46.483Z",
    "updateTime": "2024-02-05T23:07:46.483Z"
  },
  "seat": {
    "product": {
      "id": "string",
      "uniqueId": "test",
      "displayName": "Test"
    }
  },
  "createTime": "2024-02-05T23:07:46.483Z",
  "updateTime": "2024-02-05T23:07:46.483Z"
}`

	n := &organizationsImpl{transport: tr}

	res, err := n.AddMember(context.Background(), "organizationId", nil)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, `POST`, tr.Request.Method())
	require.Equal(t, `/admin/v1/organizations/organizationId/members`, tr.Request.Path())

	res, err = n.AddMember(context.Background(), "organizationId", &OrganizationAddMemberInput{})
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestOrganizations_GetMember(t *testing.T) {
	t.Parallel()

	tr := &internal.TestTransport{}
	tr.Body = `{
  "state": "ACTIVE",
  "user": {
    "id": "string",
    "state": "ACTIVE",
    "stateReason": "DELETED",
    "uniqueId": "test",
    "displayName": "Test",
    "email": "test@example.com",
    "emailVerified": true,
    "phoneNumber": "+12125550123",
    "phoneNumberVerified": true,
    "imageUrl": "https://example.com/test.png",
    "currencyCode": "USD",
    "languageCode": "en",
    "regionCode": "US",
    "timeZone": "America/New_York",
    "address": {
      "lines": [],
      "city": "Brooklyn",
      "state": "string",
      "postalCode": "11222",
      "country": "US"
    },
    "accountConnections": [],
    "subscription": {
      "id": "string",
      "state": "ACTIVE",
      "anchorTime": "2024-02-05T23:07:46.483Z"
    },
    "memberships": [],
    "signupTime": "2024-02-05T23:07:46.483Z",
    "disabled": true,
    "createTime": "2024-02-05T23:07:46.483Z",
    "updateTime": "2024-02-05T23:07:46.483Z"
  },
  "role": {
    "id": "string",
    "uniqueId": "test",
    "displayName": "Test",
    "type": "OWNER",
    "description": "string",
    "permissionSets": [
      "string"
    ],
    "default": true,
    "archived": true,
    "createTime": "2024-02-05T23:07:46.483Z",
    "updateTime": "2024-02-05T23:07:46.483Z"
  },
  "seat": {
    "product": {
      "id": "string",
      "uniqueId": "test",
      "displayName": "Test"
    }
  },
  "createTime": "2024-02-05T23:07:46.483Z",
  "updateTime": "2024-02-05T23:07:46.483Z"
}`

	n := &organizationsImpl{transport: tr}

	res, err := n.GetMember(context.Background(), "organizationId", "userId", nil)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, `GET`, tr.Request.Method())
	require.Equal(t, `/admin/v1/organizations/organizationId/members/userId`, tr.Request.Path())

	res, err = n.GetMember(context.Background(), "organizationId", "userId", &OrganizationGetMemberInput{})
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestOrganizations_UpdateMember(t *testing.T) {
	t.Parallel()

	tr := &internal.TestTransport{}
	tr.Body = `{
  "state": "ACTIVE",
  "user": {
    "id": "string",
    "state": "ACTIVE",
    "stateReason": "DELETED",
    "uniqueId": "test",
    "displayName": "Test",
    "email": "test@example.com",
    "emailVerified": true,
    "phoneNumber": "+12125550123",
    "phoneNumberVerified": true,
    "imageUrl": "https://example.com/test.png",
    "currencyCode": "USD",
    "languageCode": "en",
    "regionCode": "US",
    "timeZone": "America/New_York",
    "address": {
      "lines": [],
      "city": "Brooklyn",
      "state": "string",
      "postalCode": "11222",
      "country": "US"
    },
    "accountConnections": [],
    "subscription": {
      "id": "string",
      "state": "ACTIVE",
      "anchorTime": "2024-02-05T23:07:46.483Z"
    },
    "memberships": [],
    "signupTime": "2024-02-05T23:07:46.483Z",
    "disabled": true,
    "createTime": "2024-02-05T23:07:46.483Z",
    "updateTime": "2024-02-05T23:07:46.483Z"
  },
  "role": {
    "id": "string",
    "uniqueId": "test",
    "displayName": "Test",
    "type": "OWNER",
    "description": "string",
    "permissionSets": [
      "string"
    ],
    "default": true,
    "archived": true,
    "createTime": "2024-02-05T23:07:46.483Z",
    "updateTime": "2024-02-05T23:07:46.483Z"
  },
  "seat": {
    "product": {
      "id": "string",
      "uniqueId": "test",
      "displayName": "Test"
    }
  },
  "createTime": "2024-02-05T23:07:46.483Z",
  "updateTime": "2024-02-05T23:07:46.483Z"
}`

	n := &organizationsImpl{transport: tr}

	res, err := n.UpdateMember(context.Background(), "organizationId", "userId", nil)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, `PATCH`, tr.Request.Method())
	require.Equal(t, `/admin/v1/organizations/organizationId/members/userId`, tr.Request.Path())

	res, err = n.UpdateMember(context.Background(), "organizationId", "userId", &OrganizationUpdateMemberInput{})
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestOrganizations_RemoveMember(t *testing.T) {
	t.Parallel()

	tr := &internal.TestTransport{}
	tr.Body = `{}`

	n := &organizationsImpl{transport: tr}

	res, err := n.RemoveMember(context.Background(), "organizationId", "userId", nil)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, `DELETE`, tr.Request.Method())
	require.Equal(t, `/admin/v1/organizations/organizationId/members/userId`, tr.Request.Path())

	res, err = n.RemoveMember(context.Background(), "organizationId", "userId", &OrganizationRemoveMemberInput{})
	require.NoError(t, err)
	require.NotNil(t, res)
}
