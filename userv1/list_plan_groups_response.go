// Code generated. DO NOT EDIT.

package userv1

// Response message for ListPlanGroups.
type ListPlanGroupsResponse struct {
	// The list of plan groups.
	PlanGroups []*PlanGroup `json:"planGroups"`
	// A token, which can be sent as `pageToken` to retrieve the next page.
	// If this field is omitted, there are no subsequent pages.
	NextPageToken string `json:"nextPageToken"`
	// A token, which can be sent as `pageToken` to retrieve the previous page.
	// If this field is absent, there are no preceding pages. If this field is
	// an empty string then the previous page is the first result.
	PreviousPageToken string `json:"previousPageToken"`
}
