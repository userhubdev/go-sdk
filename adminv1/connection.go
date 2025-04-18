// Code generated. DO NOT EDIT.

package adminv1

import (
	"time"
)

// An integration that connects your tenant to an external system.
type Connection struct {
	// The system-assigned identifier of the connection.
	Id string `json:"id"`
	// The client defined unique identifier of the connection.
	//
	// It is restricted to letters, numbers, underscores, and hyphens,
	// with the first character a letter or a number, and a 255
	// character maximum.
	//
	// ID's starting with `conn_` are reserved.
	UniqueId string `json:"uniqueId"`
	// The human-readable display name of the connection.
	//
	// The maximum length is 200 characters.
	DisplayName string `json:"displayName"`
	// The current status of the connection.
	State string `json:"state"`
	// The code that best describes the reason for the state.
	StateReason string `json:"stateReason"`
	// The connection type.
	Type string `json:"type"`
	// The delegated connection.
	Delegate *ConnectionDelegate `json:"delegate"`
	// The connection providers.
	Providers []*ConnectionProvider `json:"providers"`
	// The connection view.
	View string `json:"view"`
	// The creation time of the connection.
	CreateTime time.Time `json:"createTime"`
	// The last update time of the connection.
	UpdateTime time.Time `json:"updateTime"`
	// The Amazon Cognito connection data.
	AmazonCognito *AmazonCognitoConnection `json:"amazonCognito"`
	// The Auth0 connection data.
	Auth0 *Auth0Connection `json:"auth0"`
	// The builtin email configuration data.
	BuiltinEmail *BuiltinEmailConnection `json:"builtinEmail"`
	// The Google Cloud Identity Platform (Firebase Auth) connection.
	GoogleCloudIdentityPlatform *GoogleCloudIdentityPlatformConnection `json:"googleCloudIdentityPlatform"`
	// The Postmark configuration data.
	Postmark *PostmarkConnection `json:"postmark"`
	// The Stripe billing configuration data.
	Stripe *StripeConnection `json:"stripe"`
	// The webhooks configuration data.
	Webhook *WebhookConnection `json:"webhook"`
}
