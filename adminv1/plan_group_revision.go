// Code generated. DO NOT EDIT.

package adminv1

import (
	"time"
)

// Plan group revisions track the configuration options for a plan group.
type PlanGroupRevision struct {
	// The system-assigned identifier of the plan group revision.
	Id string `json:"id"`
	// The default status of the revision.
	//
	// When this is true, the revision will be used as the default when not
	// explicitly specified.
	Default bool `json:"default"`
	// The supported currency codes (e.g. `USD`).
	CurrencyCodes []string `json:"currencyCodes"`
	// The plans associated with plan group revision.
	Plans []*PlanGroupRevisionPlan `json:"plans"`
	// The items associated with plan group revision.
	Items []*PlanGroupRevisionItem `json:"items"`
	// Whether the revision has been committed.
	//
	// After the revision has been committed, it is available for use, but
	// can no longer be edited.
	Committed bool `json:"committed"`
	// The tags associated with the revision.
	//
	// Tags are restricted to letters, numbers, underscores, and hyphens,
	// with the first character a letter or a number, and a 255
	// character maximum.
	//
	// A revision is limited to 10 tags.
	Tags []string `json:"tags"`
	// The revised plan group revision identifier.
	//
	// This allows you to track the revision lineage.
	SourceRevisionId string `json:"sourceRevisionId"`
	// The creation time of the plan group revision.
	CreateTime time.Time `json:"createTime"`
	// The last update time of the plan group revision.
	UpdateTime time.Time `json:"updateTime"`
}
