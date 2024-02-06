// Code generated. DO NOT EDIT.

package adminv1

// Response message for ListOrganizations.
type ListOrganizationsResponse struct {
	// The list of organizations.
	Organizations []*Organization `json:"organizations"`
	// A token, which can be sent as `pageToken` to retrieve the next page.
	// If this field is omitted, there are no subsequent pages.
	NextPageToken string `json:"nextPageToken"`
	// A token, which can be sent as `pageToken` to retrieve the previous page.
	// If this field is absent, there are no preceding pages. If this field is
	// an empty string then the previous page is the first result.
	PreviousPageToken string `json:"previousPageToken"`
}
