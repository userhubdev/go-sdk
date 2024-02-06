// Code generated. DO NOT EDIT.

package adminv1

// Member input parameters.
type MemberInput struct {
	// The identifier of the user.
	UserId *string `json:"userId"`
	// The identifier of the role.
	//
	// This is currently limited to `member`, `admin`, and `owner`.
	RoleId *string `json:"roleId"`
}
