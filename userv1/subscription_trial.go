// Code generated. DO NOT EDIT.

package userv1

import (
	"time"
)

// The trial information for the subscription.
type SubscriptionTrial struct {
	// The time the trial started.
	StartTime time.Time `json:"startTime"`
	// The time the trial ended/ends.
	EndTime time.Time `json:"endTime"`
	// The number of days in the trial.
	//
	// This number is rounded to the nearest whole number
	// of days.
	Days int32 `json:"days"`
	// The number of days remaining in the trial.
	//
	// This number is rounded down, so will generally be
	// less than days. It will be zero on the last day
	// of the trial and null when the trial expires.
	RemainingDays int32 `json:"remainingDays"`
}
