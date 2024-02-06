// Code generated. DO NOT EDIT.

package connectionsv1

// Response message for listing custom users.
type ListCustomUsersResponse struct {
	// The list of users.
	Users []*CustomUser `json:"users"`
	// A token the webhook can set to indicate it has more results.
	//
	// This can be a page number, offset number, timestamp, or any value
	// the webhook handler wants to use for pagination.
	//
	// It must be encoded as a string.
	NextPageToken string `json:"nextPageToken"`
}
