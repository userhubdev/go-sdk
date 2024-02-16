// Code generated. DO NOT EDIT.

package userapi

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/userhubdev/go-sdk/internal"
)

func TestInvoices_List(t *testing.T) {
	tr := &internal.TestTransport{}
	tr.Body = `{
  "invoices": [
    {
      "id": "string",
      "state": "OPEN",
      "stateTime": "2024-02-05T23:07:46.483Z",
      "number": "string",
      "currencyCode": "USD",
      "description": "string",
      "effectiveTime": "2024-02-05T23:07:46.483Z",
      "subtotalAmount": "string",
      "discountAmount": "string",
      "taxAmount": "string",
      "totalAmount": "string",
      "dueAmount": "string",
      "remainingDueAmount": "string",
      "dueTime": "2024-02-05T23:07:46.483Z",
      "paidAmount": "string",
      "paymentState": "ACTION_REQUIRED",
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
	require.Equal(t, `/user/v1/invoices`, tr.Request.Path())

	res, err = n.List(context.Background(), &InvoiceListInput{})
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestInvoices_Get(t *testing.T) {
	tr := &internal.TestTransport{}
	tr.Body = `{
  "id": "string",
  "state": "OPEN",
  "stateTime": "2024-02-05T23:07:46.483Z",
  "number": "string",
  "currencyCode": "USD",
  "description": "string",
  "account": {
    "fullName": "Test",
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
  "subtotalAmount": "string",
  "discountAmount": "string",
  "balance": {
    "startAmount": "string",
    "endAmount": "string",
    "appliedAmount": "string"
  },
  "taxAmount": "string",
  "totalAmount": "string",
  "dueAmount": "string",
  "remainingDueAmount": "string",
  "dueTime": "2024-02-05T23:07:46.483Z",
  "paidAmount": "string",
  "paymentState": "ACTION_REQUIRED",
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
      "subtotalAmount": "string",
      "discountAmount": "string",
      "description": "string",
      "proration": true
    }
  ],
  "changes": [
    {
      "time": "2024-02-05T23:07:46.483Z",
      "description": "string",
      "subtotalAmount": "string",
      "discountAmount": "string",
      "startQuantity": 1,
      "endQuantity": 1,
      "startItemIds": [],
      "endItemIds": []
    }
  ],
  "createTime": "2024-02-05T23:07:46.483Z",
  "updateTime": "2024-02-05T23:07:46.483Z"
}`

	n := &invoicesImpl{transport: tr}

	res, err := n.Get(context.Background(), "invoiceId", nil)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, `GET`, tr.Request.Method())
	require.Equal(t, `/user/v1/invoices/invoiceId`, tr.Request.Path())

	res, err = n.Get(context.Background(), "invoiceId", &InvoiceGetInput{})
	require.NoError(t, err)
	require.NotNil(t, res)
}
