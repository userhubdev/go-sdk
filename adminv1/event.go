// Code generated. DO NOT EDIT.

package adminv1

import (
	"time"
)

// Event describes a service a tenant provides. Multiple
// events can be multiple and sold using a plan.
type Event struct {
	// The system-assigned identifier of the event.
	Id string `json:"id"`
	// The type of the event.
	Type string `json:"type"`
	// The time of the event.
	Time time.Time `json:"time"`
	// The entity associated with the event.
	Entity *EventEntity `json:"entity"`
	// The connection associated with the event.
	Connection *EventConnection `json:"connection"`
	// The organization associated with the event.
	Organization *EventOrganization `json:"organization"`
	// The user associated with the event.
	User *EventUser `json:"user"`
	// The API key associated with the event.
	ApiKey *EventApiKey `json:"apiKey"`
	// The actor associated with the event.
	Actor *EventActor `json:"actor"`
	// The request associated with the event.
	Request *EventRequest `json:"request"`
}
