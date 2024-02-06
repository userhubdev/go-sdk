// Code generated. DO NOT EDIT.

package userv1

import (
	"time"
)

// Information about the current billing period.
type SubscriptionCurrentPeriod struct {
	// The time the current billing period started.
	StartTime time.Time `json:"startTime"`
	// The time the current billing period ends.
	EndTime time.Time `json:"endTime"`
}
