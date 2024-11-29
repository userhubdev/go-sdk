// Code generated. DO NOT EDIT.

package adminv1

import (
	"time"
)

// Plan group is a container for plan revisions and billing
// intervals.
//
// A plan group would generally describe a tier of plans offered
// (e.g. Pro) which might contain two options, a monthly and
// yearly plan.
type PlanGroup struct {
	// The system-assigned identifier of the plan group.
	Id string `json:"id"`
	// The client defined unique identifier of the plan group (e.g. `pro`).
	//
	// It is restricted to letters, numbers, underscores, and hyphens,
	// with the first character a letter or a number, and a 255
	// character maximum.
	//
	// ID's starting with `pg_r are reserved.
	UniqueId string `json:"uniqueId"`
	// The customer facing human-readable display name of the plan group.
	//
	// The maximum length is 200 characters.
	DisplayName string `json:"displayName"`
	// The admin-facing description of the plan group.
	//
	// The maximum length is 1000 characters.
	Description string `json:"description"`
	// The type of account the plan can be activated for.
	AccountType string `json:"accountType"`
	// The trial settings.
	Trial *PlanGroupTrial `json:"trial"`
	// The visibility of the plan group.
	Visibility string `json:"visibility"`
	// The archived status of the plan group.
	Archived bool `json:"archived"`
	// The current revision for the plan group.
	Revision *PlanGroupRevision `json:"revision"`
	// The creation time of the plan group.
	CreateTime time.Time `json:"createTime"`
	// The last update time of the plan group.
	UpdateTime time.Time `json:"updateTime"`
}
