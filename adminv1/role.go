// Code generated. DO NOT EDIT.

package adminv1

import (
	"time"
)

// A member's role within an organization.
type Role struct {
	// The system-assigned identifier of the role.
	Id string `json:"id"`
	// The client defined unique identifier of the role.
	//
	// It is restricted to letters, numbers, underscores, and hyphens,
	// with the first character a letter or a number, and a 255
	// character maximum.
	//
	// ID's starting with `role_` are reserved.
	UniqueId string `json:"uniqueId"`
	// The human-readable display name of the role.
	DisplayName string `json:"displayName"`
	// The role type.
	Type string `json:"type"`
	// The description of the role.
	//
	// The maximum length is 1000 characters.
	Description string `json:"description"`
	// The policy that defines how a member is assigned a seat.
	SeatPolicy string `json:"seatPolicy"`
	// The additional permissions allowed by the role.
	PermissionSets []string `json:"permissionSets"`
	// Whether the role is the default for the tenant.
	Default bool `json:"default"`
	// The archived status of the role.
	Archived bool `json:"archived"`
	// The role view.
	View string `json:"view"`
	// The creation time of the role.
	CreateTime time.Time `json:"createTime"`
	// The last update time of the role.
	UpdateTime time.Time `json:"updateTime"`
}
