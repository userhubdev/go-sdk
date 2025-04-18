// Code generated. DO NOT EDIT.

package adminapi

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/userhubdev/go-sdk/internal"
)

func TestInvoices_List(t *testing.T) {
	t.Parallel()

	tr := &internal.TestTransport{}
	tr.Body = `{
  "invoices": [
    {
      "id": "string",
      "state": "DRAFT",
      "stateReason": "UPDATING",
      "stateTime": "2024-02-05T23:07:46.483Z",
      "externalId": "string",
      "number": "string",
      "currencyCode": "USD",
      "description": "string",
      "effectiveTime": "2024-02-05T23:07:46.483Z",
      "subtotalAmount": "10",
      "discountAmount": "0",
      "taxAmount": "0",
      "totalAmount": "10",
      "dueAmount": "10",
      "remainingDueAmount": "0",
      "dueTime": "2024-02-05T23:07:46.483Z",
      "paidAmount": "10",
      "paymentState": "PAYMENT_METHOD_REQUIRED",
      "pullTime": "2024-02-05T23:07:46.483Z",
      "createTime": "2024-02-05T23:07:46.483Z",
      "updateTime": "2024-02-05T23:07:46.483Z"
    }
  ],
  "nextPageToken": "string",
  "previousPageToken": "string"
}`

	n := &invoicesImpl{transport: tr}

	res, err := n.List(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, `GET`, tr.Request.Method())
	require.Equal(t, `/admin/v1/invoices`, tr.Request.Path())

	res, err = n.List(context.Background(), &InvoiceListInput{})
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestInvoices_Get(t *testing.T) {
	t.Parallel()

	tr := &internal.TestTransport{}
	tr.Body = `{
  "id": "string",
  "state": "DRAFT",
  "stateReason": "UPDATING",
  "stateTime": "2024-02-05T23:07:46.483Z",
  "connection": {
    "id": "string",
    "uniqueId": "test",
    "displayName": "Test",
    "state": "ACTIVE",
    "stateReason": "UPDATING",
    "type": "AMAZON_COGNITO",
    "delegate": {
      "id": "string",
      "uniqueId": "test",
      "displayName": "Test",
      "state": "ACTIVE",
      "stateReason": "UPDATING",
      "type": "AMAZON_COGNITO"
    },
    "providers": [],
    "view": "BASIC",
    "createTime": "2024-02-05T23:07:46.483Z",
    "updateTime": "2024-02-05T23:07:46.483Z",
    "amazonCognito": {
      "userPoolId": "string",
      "region": "string",
      "accessKeyId": "string",
      "accessKeySecret": "string"
    },
    "auth0": {
      "domain": "string",
      "clientId": "string",
      "clientSecret": "string"
    },
    "builtinEmail": {
      "allowedEmails": []
    },
    "googleCloudIdentityPlatform": {
      "credentials": "string",
      "projectId": "string"
    },
    "postmark": {
      "serverToken": "string",
      "serverId": "string",
      "allowedEmails": []
    },
    "stripe": {
      "accountId": "string",
      "live": true
    },
    "webhook": {
      "url": "https://example.com",
      "headers": {}
    }
  },
  "externalId": "string",
  "number": "string",
  "currencyCode": "USD",
  "description": "string",
  "account": {
    "fullName": "Jane Doe",
    "email": "test@example.com",
    "phoneNumber": "+12125550123",
    "address": {
      "lines": [],
      "city": "Brooklyn",
      "state": "string",
      "postalCode": "11222",
      "country": "US"
    }
  },
  "effectiveTime": "2024-02-05T23:07:46.483Z",
  "period": {
    "startTime": "2024-02-05T23:07:46.483Z",
    "endTime": "2024-02-05T23:07:46.483Z"
  },
  "subtotalAmount": "10",
  "discountAmount": "0",
  "balance": {
    "startAmount": "10",
    "endAmount": "10",
    "appliedAmount": "10"
  },
  "taxAmount": "0",
  "totalAmount": "10",
  "dueAmount": "10",
  "remainingDueAmount": "0",
  "dueTime": "2024-02-05T23:07:46.483Z",
  "paidAmount": "10",
  "paymentState": "PAYMENT_METHOD_REQUIRED",
  "paymentIntent": {
    "stripe": {
      "accountId": "string",
      "live": true,
      "clientSecret": "string"
    }
  },
  "items": [
    {
      "id": "string",
      "quantity": 1,
      "subtotalAmount": "10",
      "discountAmount": "0",
      "description": "string",
      "externalId": "string",
      "proration": true
    }
  ],
  "changes": [
    {
      "time": "2024-02-05T23:07:46.483Z",
      "description": "string",
      "subtotalAmount": "10",
      "discountAmount": "0",
      "startQuantity": 1,
      "endQuantity": 1,
      "startItemIds": [],
      "endItemIds": []
    }
  ],
  "pullTime": "2024-02-05T23:07:46.483Z",
  "createTime": "2024-02-05T23:07:46.483Z",
  "updateTime": "2024-02-05T23:07:46.483Z"
}`

	n := &invoicesImpl{transport: tr}

	res, err := n.Get(context.Background(), "invoiceId", nil)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, `GET`, tr.Request.Method())
	require.Equal(t, `/admin/v1/invoices/invoiceId`, tr.Request.Path())

	res, err = n.Get(context.Background(), "invoiceId", &InvoiceGetInput{})
	require.NoError(t, err)
	require.NotNil(t, res)
}
