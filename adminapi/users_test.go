// Code generated. DO NOT EDIT.

package adminapi

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/userhubdev/go-sdk/internal"
)

func TestUsers_List(t *testing.T) {
	t.Parallel()

	tr := &internal.TestTransport{}
	tr.Body = `{
  "users": [
    {
      "id": "string",
      "state": "ACTIVE",
      "stateReason": "UPDATING",
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
      "disabled": true,
      "view": "BASIC",
      "createTime": "2024-02-05T23:07:46.483Z",
      "updateTime": "2024-02-05T23:07:46.483Z"
    }
  ],
  "nextPageToken": "string",
  "previousPageToken": "string"
}`

	n := &usersImpl{transport: tr}

	res, err := n.List(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, `GET`, tr.Request.Method())
	require.Equal(t, `/admin/v1/users`, tr.Request.Path())

	res, err = n.List(context.Background(), &UserListInput{})
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestUsers_Create(t *testing.T) {
	t.Parallel()

	tr := &internal.TestTransport{}
	tr.Body = `{
  "id": "string",
  "state": "ACTIVE",
  "stateReason": "UPDATING",
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
      "stateReason": "UPDATING",
      "displayName": "Test",
      "email": "test@example.com",
      "emailVerified": true,
      "phoneNumber": "+12125550123",
      "phoneNumberVerified": true,
      "currencyCode": "USD",
      "balanceAmount": "10",
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
  "memberships": [
    {
      "createTime": "2024-02-05T23:07:46.483Z",
      "updateTime": "2024-02-05T23:07:46.483Z"
    }
  ],
  "signupTime": "2024-02-05T23:07:46.483Z",
  "disabled": true,
  "view": "BASIC",
  "createTime": "2024-02-05T23:07:46.483Z",
  "updateTime": "2024-02-05T23:07:46.483Z"
}`

	n := &usersImpl{transport: tr}

	res, err := n.Create(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, `POST`, tr.Request.Method())
	require.Equal(t, `/admin/v1/users`, tr.Request.Path())

	res, err = n.Create(context.Background(), &UserCreateInput{})
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestUsers_Get(t *testing.T) {
	t.Parallel()

	tr := &internal.TestTransport{}
	tr.Body = `{
  "id": "string",
  "state": "ACTIVE",
  "stateReason": "UPDATING",
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
      "stateReason": "UPDATING",
      "displayName": "Test",
      "email": "test@example.com",
      "emailVerified": true,
      "phoneNumber": "+12125550123",
      "phoneNumberVerified": true,
      "currencyCode": "USD",
      "balanceAmount": "10",
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
  "memberships": [
    {
      "createTime": "2024-02-05T23:07:46.483Z",
      "updateTime": "2024-02-05T23:07:46.483Z"
    }
  ],
  "signupTime": "2024-02-05T23:07:46.483Z",
  "disabled": true,
  "view": "BASIC",
  "createTime": "2024-02-05T23:07:46.483Z",
  "updateTime": "2024-02-05T23:07:46.483Z"
}`

	n := &usersImpl{transport: tr}

	res, err := n.Get(context.Background(), "userId", nil)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, `GET`, tr.Request.Method())
	require.Equal(t, `/admin/v1/users/userId`, tr.Request.Path())

	res, err = n.Get(context.Background(), "userId", &UserGetInput{})
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestUsers_Update(t *testing.T) {
	t.Parallel()

	tr := &internal.TestTransport{}
	tr.Body = `{
  "id": "string",
  "state": "ACTIVE",
  "stateReason": "UPDATING",
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
      "stateReason": "UPDATING",
      "displayName": "Test",
      "email": "test@example.com",
      "emailVerified": true,
      "phoneNumber": "+12125550123",
      "phoneNumberVerified": true,
      "currencyCode": "USD",
      "balanceAmount": "10",
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
  "memberships": [
    {
      "createTime": "2024-02-05T23:07:46.483Z",
      "updateTime": "2024-02-05T23:07:46.483Z"
    }
  ],
  "signupTime": "2024-02-05T23:07:46.483Z",
  "disabled": true,
  "view": "BASIC",
  "createTime": "2024-02-05T23:07:46.483Z",
  "updateTime": "2024-02-05T23:07:46.483Z"
}`

	n := &usersImpl{transport: tr}

	res, err := n.Update(context.Background(), "userId", nil)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, `PATCH`, tr.Request.Method())
	require.Equal(t, `/admin/v1/users/userId`, tr.Request.Path())

	res, err = n.Update(context.Background(), "userId", &UserUpdateInput{})
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestUsers_Delete(t *testing.T) {
	t.Parallel()

	tr := &internal.TestTransport{}
	tr.Body = `{
  "id": "string",
  "state": "ACTIVE",
  "stateReason": "UPDATING",
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
      "stateReason": "UPDATING",
      "displayName": "Test",
      "email": "test@example.com",
      "emailVerified": true,
      "phoneNumber": "+12125550123",
      "phoneNumberVerified": true,
      "currencyCode": "USD",
      "balanceAmount": "10",
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
  "memberships": [
    {
      "createTime": "2024-02-05T23:07:46.483Z",
      "updateTime": "2024-02-05T23:07:46.483Z"
    }
  ],
  "signupTime": "2024-02-05T23:07:46.483Z",
  "disabled": true,
  "view": "BASIC",
  "createTime": "2024-02-05T23:07:46.483Z",
  "updateTime": "2024-02-05T23:07:46.483Z"
}`

	n := &usersImpl{transport: tr}

	res, err := n.Delete(context.Background(), "userId", nil)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, `DELETE`, tr.Request.Method())
	require.Equal(t, `/admin/v1/users/userId`, tr.Request.Path())

	res, err = n.Delete(context.Background(), "userId", &UserDeleteInput{})
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestUsers_Undelete(t *testing.T) {
	t.Parallel()

	tr := &internal.TestTransport{}
	tr.Body = `{
  "id": "string",
  "state": "ACTIVE",
  "stateReason": "UPDATING",
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
      "stateReason": "UPDATING",
      "displayName": "Test",
      "email": "test@example.com",
      "emailVerified": true,
      "phoneNumber": "+12125550123",
      "phoneNumberVerified": true,
      "currencyCode": "USD",
      "balanceAmount": "10",
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
  "memberships": [
    {
      "createTime": "2024-02-05T23:07:46.483Z",
      "updateTime": "2024-02-05T23:07:46.483Z"
    }
  ],
  "signupTime": "2024-02-05T23:07:46.483Z",
  "disabled": true,
  "view": "BASIC",
  "createTime": "2024-02-05T23:07:46.483Z",
  "updateTime": "2024-02-05T23:07:46.483Z"
}`

	n := &usersImpl{transport: tr}

	res, err := n.Undelete(context.Background(), "userId", nil)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, `POST`, tr.Request.Method())
	require.Equal(t, `/admin/v1/users/userId:undelete`, tr.Request.Path())

	res, err = n.Undelete(context.Background(), "userId", &UserUndeleteInput{})
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestUsers_Purge(t *testing.T) {
	t.Parallel()

	tr := &internal.TestTransport{}
	tr.Body = `{}`

	n := &usersImpl{transport: tr}

	res, err := n.Purge(context.Background(), "userId", nil)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, `POST`, tr.Request.Method())
	require.Equal(t, `/admin/v1/users/userId:purge`, tr.Request.Path())

	res, err = n.Purge(context.Background(), "userId", &UserPurgeInput{})
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestUsers_Connect(t *testing.T) {
	t.Parallel()

	tr := &internal.TestTransport{}
	tr.Body = `{
  "id": "string",
  "state": "ACTIVE",
  "stateReason": "UPDATING",
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
      "stateReason": "UPDATING",
      "displayName": "Test",
      "email": "test@example.com",
      "emailVerified": true,
      "phoneNumber": "+12125550123",
      "phoneNumberVerified": true,
      "currencyCode": "USD",
      "balanceAmount": "10",
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
  "memberships": [
    {
      "createTime": "2024-02-05T23:07:46.483Z",
      "updateTime": "2024-02-05T23:07:46.483Z"
    }
  ],
  "signupTime": "2024-02-05T23:07:46.483Z",
  "disabled": true,
  "view": "BASIC",
  "createTime": "2024-02-05T23:07:46.483Z",
  "updateTime": "2024-02-05T23:07:46.483Z"
}`

	n := &usersImpl{transport: tr}

	res, err := n.Connect(context.Background(), "userId", nil)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, `POST`, tr.Request.Method())
	require.Equal(t, `/admin/v1/users/userId:connect`, tr.Request.Path())

	res, err = n.Connect(context.Background(), "userId", &UserConnectInput{})
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestUsers_Disconnect(t *testing.T) {
	t.Parallel()

	tr := &internal.TestTransport{}
	tr.Body = `{
  "id": "string",
  "state": "ACTIVE",
  "stateReason": "UPDATING",
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
      "stateReason": "UPDATING",
      "displayName": "Test",
      "email": "test@example.com",
      "emailVerified": true,
      "phoneNumber": "+12125550123",
      "phoneNumberVerified": true,
      "currencyCode": "USD",
      "balanceAmount": "10",
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
  "memberships": [
    {
      "createTime": "2024-02-05T23:07:46.483Z",
      "updateTime": "2024-02-05T23:07:46.483Z"
    }
  ],
  "signupTime": "2024-02-05T23:07:46.483Z",
  "disabled": true,
  "view": "BASIC",
  "createTime": "2024-02-05T23:07:46.483Z",
  "updateTime": "2024-02-05T23:07:46.483Z"
}`

	n := &usersImpl{transport: tr}

	res, err := n.Disconnect(context.Background(), "userId", nil)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, `POST`, tr.Request.Method())
	require.Equal(t, `/admin/v1/users/userId:disconnect`, tr.Request.Path())

	res, err = n.Disconnect(context.Background(), "userId", &UserDisconnectInput{})
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestUsers_ImportAccount(t *testing.T) {
	t.Parallel()

	tr := &internal.TestTransport{}
	tr.Body = `{
  "id": "string",
  "state": "ACTIVE",
  "stateReason": "UPDATING",
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
      "stateReason": "UPDATING",
      "displayName": "Test",
      "email": "test@example.com",
      "emailVerified": true,
      "phoneNumber": "+12125550123",
      "phoneNumberVerified": true,
      "currencyCode": "USD",
      "balanceAmount": "10",
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
  "memberships": [
    {
      "createTime": "2024-02-05T23:07:46.483Z",
      "updateTime": "2024-02-05T23:07:46.483Z"
    }
  ],
  "signupTime": "2024-02-05T23:07:46.483Z",
  "disabled": true,
  "view": "BASIC",
  "createTime": "2024-02-05T23:07:46.483Z",
  "updateTime": "2024-02-05T23:07:46.483Z"
}`

	n := &usersImpl{transport: tr}

	res, err := n.ImportAccount(context.Background(), "userId", nil)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, `POST`, tr.Request.Method())
	require.Equal(t, `/admin/v1/users/userId:import`, tr.Request.Path())

	res, err = n.ImportAccount(context.Background(), "userId", &UserImportAccountInput{})
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestUsers_CreateApiSession(t *testing.T) {
	t.Parallel()

	tr := &internal.TestTransport{}
	tr.Body = `{
  "accessToken": "string",
  "expireTime": "2024-02-05T23:07:46.483Z"
}`

	n := &usersImpl{transport: tr}

	res, err := n.CreateApiSession(context.Background(), "userId", nil)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, `POST`, tr.Request.Method())
	require.Equal(t, `/admin/v1/users/userId:createApiSession`, tr.Request.Path())

	res, err = n.CreateApiSession(context.Background(), "userId", &UserCreateApiSessionInput{})
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestUsers_CreatePortalSession(t *testing.T) {
	t.Parallel()

	tr := &internal.TestTransport{}
	tr.Body = `{
  "redirectUrl": "https://example.com"
}`

	n := &usersImpl{transport: tr}

	res, err := n.CreatePortalSession(context.Background(), "userId", nil)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, `POST`, tr.Request.Method())
	require.Equal(t, `/admin/v1/users/userId:createPortalSession`, tr.Request.Path())

	res, err = n.CreatePortalSession(context.Background(), "userId", &UserCreatePortalSessionInput{})
	require.NoError(t, err)
	require.NotNil(t, res)
}
