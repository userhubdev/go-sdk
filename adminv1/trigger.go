// Code generated. DO NOT EDIT.

package adminv1

import (
	"time"
)

// A trigger is a way to run connection functionality when specific events
// occur.
type Trigger struct {
	// The connection.
	Connection *Connection `json:"connection"`
	// The event type.
	EventType string `json:"eventType"`
	// The creation time of the trigger.
	CreateTime time.Time `json:"createTime"`
	// The last update time of the trigger.
	UpdateTime time.Time `json:"updateTime"`
}
