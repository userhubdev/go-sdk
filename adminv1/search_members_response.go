// Code generated. DO NOT EDIT.

package adminv1

// Response message for SearchMembers.
type SearchMembersResponse struct {
	// The search of members.
	Members []*Member `json:"members"`
	// A token, which can be sent as `pageToken` to retrieve the next page.
	// If this field is omitted, there are no subsequent pages.
	NextPageToken string `json:"nextPageToken"`
	// A token, which can be sent as `pageToken` to retrieve the previous page.
	// If this field is absent, there are no preceding pages. If this field is
	// an empty string then the previous page is the first result.
	PreviousPageToken string `json:"previousPageToken"`
	// The estimated total count of matched members irrespective of pagination.
	//
	// This field is ignored if `show_total_size` is not true or `pageToken`
	// is not empty.
	TotalSize int32 `json:"totalSize"`
}
