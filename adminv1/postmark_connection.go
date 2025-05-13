// Code generated. DO NOT EDIT.

package adminv1

import (
	"github.com/userhubdev/go-sdk/commonv1"
)

// The Postmark specific connection data.
type PostmarkConnection struct {
	// The Postmark server token (e.g. `942faf79-bf10-4dc1-830a-dc7943f43f35`).
	ServerToken string `json:"serverToken"`
	// The Postmark server ID.
	//
	// This is automatically populated when the server token is updated.
	ServerId string `json:"serverId"`
	// The from email address.
	//
	// The Postmark account must be allowed to send from this email
	// address.
	From *commonv1.Email `json:"from"`
	// The reply to email address.
	ReplyTo *commonv1.Email `json:"replyTo"`
}
