// Code generated. DO NOT EDIT.

package connectionsv1

// Request message for listing custom users.
type ListCustomUsersRequest struct {
	// The maximum number of users to return. The webhook is allowed to
	// return fewer than this value, but it should never return more.
	PageSize int32 `json:"pageSize"`
	// A page token, this is from the response of the previous list
	// request.
	//
	// This should be used to determine the next page of results.
	PageToken string `json:"pageToken"`
}
