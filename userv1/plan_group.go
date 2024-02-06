// Code generated. DO NOT EDIT.

package userv1

// Plan group is a collection of related plans.
//
// A plan group will generally describe a tier of plans offered
// (e.g. Basic vs Pro) which might contain multiple billing options
// (e.g. monthly vs yearly, USD vs EUR).
type PlanGroup struct {
	// The system-assigned identifier of the plan group.
	Id string `json:"id"`
	// The client defined unique identifier of the plan group (e.g. `pro`).
	UniqueId string `json:"uniqueId"`
	// The name of the plan group.
	DisplayName string `json:"displayName"`
	// The user facing description of the plan group.
	Description string `json:"description"`
	// Whether the plans are for organizations or users.
	AccountType string `json:"accountType"`
	// The trial settings.
	//
	// For authenticated requests, this will not be set when the account
	// isn't eligible for a trial.
	Trial *PlanGroupTrial `json:"trial"`
	// Whether the plan is consider an downgrade/upgrade.
	ChangePath *PlanGroupChangePath `json:"changePath"`
	// The plans associated with plan group.
	Plans []*Plan `json:"plans"`
}
