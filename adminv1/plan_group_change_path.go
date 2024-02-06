// Code generated. DO NOT EDIT.

package adminv1

import (
	"time"
)

// This defines the upgrade/downgrade paths for a plan group.
type PlanGroupChangePath struct {
	// The target plan group for this change path.
	Target *PlanGroup `json:"target"`
	// Whether the change is considered an upgrade or
	// a downgrade.
	Direction string `json:"direction"`
	// The visibility of the change path.
	Visibility string `json:"visibility"`
	// The creation time of the plan group change path.
	CreateTime time.Time `json:"createTime"`
	// The last update time of the plan group change path.
	UpdateTime time.Time `json:"updateTime"`
}
