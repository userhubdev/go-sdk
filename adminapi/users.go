// Code generated. DO NOT EDIT.

package adminapi

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/userhubdev/go-sdk/adminv1"
	"github.com/userhubdev/go-sdk/commonv1"
	"github.com/userhubdev/go-sdk/internal"
	"github.com/userhubdev/go-sdk/types"
)

type Users interface {
	// Lists users.
	List(ctx context.Context, input *UserListInput) (*adminv1.ListUsersResponse, error)
	// Creates a new user.
	Create(ctx context.Context, input *UserCreateInput) (*adminv1.User, error)
	// Retrieves specified user.
	Get(ctx context.Context, userId string, input *UserGetInput) (*adminv1.User, error)
	// Updates specified user.
	Update(ctx context.Context, userId string, input *UserUpdateInput) (*adminv1.User, error)
	// Marks specified user for deletion.
	Delete(ctx context.Context, userId string, input *UserDeleteInput) (*adminv1.User, error)
	// Un-marks specified user for deletion.
	Undelete(ctx context.Context, userId string, input *UserUndeleteInput) (*adminv1.User, error)
	// Connect specified user to external account.
	Connect(ctx context.Context, userId string, input *UserConnectInput) (*adminv1.User, error)
	// Disconnect specified user from external account.
	//
	// This will delete all the data associated with the connected account, including
	// payment methods, invoices, and subscriptions.
	//
	// If the delete external account option is enabled it will also attempt to delete
	// the external account (e.g. Stripe Customer object).
	//
	// WARNING: This can irreversibly destroy data and should be
	// used with extreme caution.
	Disconnect(ctx context.Context, userId string, input *UserDisconnectInput) (*adminv1.User, error)
	// Import user from external identity provider if they don't already
	// exist.
	//
	// If the user already exists in UserHub, this is a no-op.
	ImportAccount(ctx context.Context, userId string, input *UserImportAccountInput) (*adminv1.User, error)
	// Create a User API session.
	CreateApiSession(ctx context.Context, userId string, input *UserCreateApiSessionInput) (*adminv1.CreateApiSessionResponse, error)
	// Create Portal session.
	CreatePortalSession(ctx context.Context, userId string, input *UserCreatePortalSessionInput) (*adminv1.CreatePortalSessionResponse, error)
}

type usersImpl struct {
	transport internal.Transport
}

// UserListInput is the input param for the List method.
type UserListInput struct {
	// The maximum number of users to return. The API may return fewer than
	// this value.
	//
	// If unspecified, at most 20 users will be returned.
	// The maximum value is 100; values above 100 will be coerced to 100.
	PageSize int32
	// A page token, received from a previous list users call.
	// Provide this to retrieve the subsequent page.
	//
	// When paginating, all other parameters provided to list users must match
	// the call that provided the page token.
	PageToken string
	// A comma-separated list of fields to order by, sorted in ascending order.
	// Use `desc` after a field name for descending.
	//
	// Supported fields:
	// - `displayName`
	// - `email`
	// - `createTime`
	// - `deleteTime`
	OrderBy string
	// Whether to show deleted users.
	ShowDeleted bool
	// The User view to return in the results.
	//
	// This defaults to the `BASIC` view.
	View string
}

func (n *usersImpl) List(ctx context.Context, input *UserListInput) (*adminv1.ListUsersResponse, error) {
	req := internal.NewRequest(
		"admin.users.list",
		"GET",
		"/admin/v1/users",
	)
	req.SetIdempotent(true)

	if input != nil {
		if !internal.IsEmpty(input.PageSize) {
			req.SetQuery("pageSize", input.PageSize)
		}
		if !internal.IsEmpty(input.PageToken) {
			req.SetQuery("pageToken", input.PageToken)
		}
		if !internal.IsEmpty(input.OrderBy) {
			req.SetQuery("orderBy", input.OrderBy)
		}
		if !internal.IsEmpty(input.ShowDeleted) {
			req.SetQuery("showDeleted", input.ShowDeleted)
		}
		if !internal.IsEmpty(input.View) {
			req.SetQuery("view", input.View)
		}
	}

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &adminv1.ListUsersResponse{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// UserCreateInput is the input param for the Create method.
type UserCreateInput struct {
	// The client defined unique identifier of the user account.
	//
	// It is restricted to letters, numbers, underscores, and hyphens,
	// with the first character a letter or a number, and a 255
	// character maximum.
	//
	// ID's starting with `usr_` are reserved.
	UniqueId string
	// The human-readable display name of the user.
	//
	// The maximum length is 200 characters.
	DisplayName string
	// The email address of the user.
	//
	// The maximum length is 320 characters.
	Email string
	// Whether the user's email address has been verified.
	EmailVerified bool
	// The E164 phone number for the user (e.g. `+12125550123`).
	PhoneNumber string
	// Whether the user's phone number has been verified.
	PhoneNumberVerified bool
	// The photo/avatar URL of the user.
	//
	// The maximum length is 2000 characters.
	ImageUrl string
	// The default ISO-4217 currency code for the user (e.g. `USD`).
	CurrencyCode string
	// The IETF BCP-47 language code for the user (e.g. `en`).
	LanguageCode string
	// The country/region code for the user (e.g. `US`).
	RegionCode string
	// The IANA time zone for the user (e.g. `America/New_York`).
	TimeZone string
	// The billing address for the user.
	Address *commonv1.Address
	// The sign-up time for the user.
	SignupTime time.Time
	// Whether the user is disabled.
	Disabled bool
}

func (n *usersImpl) Create(ctx context.Context, input *UserCreateInput) (*adminv1.User, error) {
	req := internal.NewRequest(
		"admin.users.create",
		"POST",
		"/admin/v1/users",
	)

	body := map[string]any{}

	if input != nil {
		if !internal.IsEmpty(input.UniqueId) {
			body["uniqueId"] = input.UniqueId
		}
		if !internal.IsEmpty(input.DisplayName) {
			body["displayName"] = input.DisplayName
		}
		if !internal.IsEmpty(input.Email) {
			body["email"] = input.Email
		}
		if !internal.IsEmpty(input.EmailVerified) {
			body["emailVerified"] = input.EmailVerified
		}
		if !internal.IsEmpty(input.PhoneNumber) {
			body["phoneNumber"] = input.PhoneNumber
		}
		if !internal.IsEmpty(input.PhoneNumberVerified) {
			body["phoneNumberVerified"] = input.PhoneNumberVerified
		}
		if !internal.IsEmpty(input.ImageUrl) {
			body["imageUrl"] = input.ImageUrl
		}
		if !internal.IsEmpty(input.CurrencyCode) {
			body["currencyCode"] = input.CurrencyCode
		}
		if !internal.IsEmpty(input.LanguageCode) {
			body["languageCode"] = input.LanguageCode
		}
		if !internal.IsEmpty(input.RegionCode) {
			body["regionCode"] = input.RegionCode
		}
		if !internal.IsEmpty(input.TimeZone) {
			body["timeZone"] = input.TimeZone
		}
		if !internal.IsEmpty(input.Address) {
			body["address"] = input.Address
		}
		if !internal.IsEmpty(input.SignupTime) {
			body["signupTime"] = input.SignupTime
		}
		if !internal.IsEmpty(input.Disabled) {
			body["disabled"] = input.Disabled
		}
	}

	req.SetBody(body)

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &adminv1.User{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// UserGetInput is the input param for the Get method.
type UserGetInput struct {
}

func (n *usersImpl) Get(ctx context.Context, userId string, input *UserGetInput) (*adminv1.User, error) {
	req := internal.NewRequest(
		"admin.users.get",
		"GET",
		fmt.Sprintf("/admin/v1/users/%s",
			url.PathEscape(userId),
		),
	)
	req.SetIdempotent(true)

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &adminv1.User{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// UserUpdateInput is the input param for the Update method.
type UserUpdateInput struct {
	// The client defined unique identifier of the user account.
	//
	// It is restricted to letters, numbers, underscores, and hyphens,
	// with the first character a letter or a number, and a 255
	// character maximum.
	//
	// ID's starting with `usr_` are reserved.
	UniqueId types.Optional[string]
	// The human-readable display name of the user.
	//
	// The maximum length is 200 characters.
	DisplayName types.Optional[string]
	// The email address of the user.
	//
	// The maximum length is 320 characters.
	Email types.Optional[string]
	// Whether the user's email address has been verified.
	EmailVerified types.Optional[bool]
	// The E164 phone number for the user (e.g. `+12125550123`).
	PhoneNumber types.Optional[string]
	// Whether the user's phone number has been verified.
	PhoneNumberVerified types.Optional[bool]
	// The photo/avatar URL of the user.
	//
	// The maximum length is 2000 characters.
	ImageUrl types.Optional[string]
	// The default ISO-4217 currency code for the user (e.g. `USD`).
	CurrencyCode types.Optional[string]
	// The IETF BCP-47 language code for the user (e.g. `en`).
	LanguageCode types.Optional[string]
	// The country/region code for the user (e.g. `US`).
	RegionCode types.Optional[string]
	// The IANA time zone for the user (e.g. `America/New_York`).
	TimeZone types.Optional[string]
	// The billing address for the user.
	Address types.Optional[*commonv1.Address]
	// The sign-up time for the user.
	SignupTime types.Optional[time.Time]
	// Whether the user is disabled.
	Disabled types.Optional[bool]

	// If set to true, and the user is not found, a new user will be created.
	//
	// You must use the unique identifier for the identifier when this is true.
	AllowMissing bool
}

func (n *usersImpl) Update(ctx context.Context, userId string, input *UserUpdateInput) (*adminv1.User, error) {
	req := internal.NewRequest(
		"admin.users.update",
		"PATCH",
		fmt.Sprintf("/admin/v1/users/%s",
			url.PathEscape(userId),
		),
	)
	req.SetIdempotent(true)

	body := map[string]any{}

	if input != nil {
		if !internal.IsEmpty(input.AllowMissing) {
			req.SetQuery("allowMissing", input.AllowMissing)
		}
		if input.UniqueId.Present {
			body["uniqueId"] = input.UniqueId.Value
		}
		if input.DisplayName.Present {
			body["displayName"] = input.DisplayName.Value
		}
		if input.Email.Present {
			body["email"] = input.Email.Value
		}
		if input.EmailVerified.Present {
			body["emailVerified"] = input.EmailVerified.Value
		}
		if input.PhoneNumber.Present {
			body["phoneNumber"] = input.PhoneNumber.Value
		}
		if input.PhoneNumberVerified.Present {
			body["phoneNumberVerified"] = input.PhoneNumberVerified.Value
		}
		if input.ImageUrl.Present {
			body["imageUrl"] = input.ImageUrl.Value
		}
		if input.CurrencyCode.Present {
			body["currencyCode"] = input.CurrencyCode.Value
		}
		if input.LanguageCode.Present {
			body["languageCode"] = input.LanguageCode.Value
		}
		if input.RegionCode.Present {
			body["regionCode"] = input.RegionCode.Value
		}
		if input.TimeZone.Present {
			body["timeZone"] = input.TimeZone.Value
		}
		if input.Address.Present {
			body["address"] = input.Address.Value
		}
		if input.SignupTime.Present {
			body["signupTime"] = input.SignupTime.Value
		}
		if input.Disabled.Present {
			body["disabled"] = input.Disabled.Value
		}
	}

	req.SetBody(body)

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &adminv1.User{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// UserDeleteInput is the input param for the Delete method.
type UserDeleteInput struct {
}

func (n *usersImpl) Delete(ctx context.Context, userId string, input *UserDeleteInput) (*adminv1.User, error) {
	req := internal.NewRequest(
		"admin.users.delete",
		"DELETE",
		fmt.Sprintf("/admin/v1/users/%s",
			url.PathEscape(userId),
		),
	)

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &adminv1.User{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// UserUndeleteInput is the input param for the Undelete method.
type UserUndeleteInput struct {
}

func (n *usersImpl) Undelete(ctx context.Context, userId string, input *UserUndeleteInput) (*adminv1.User, error) {
	req := internal.NewRequest(
		"admin.users.undelete",
		"POST",
		fmt.Sprintf("/admin/v1/users/%s:undelete",
			url.PathEscape(userId),
		),
	)

	body := map[string]any{}

	req.SetBody(body)

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &adminv1.User{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// UserConnectInput is the input param for the Connect method.
type UserConnectInput struct {
	// The identifier of the connection.
	ConnectionId string
	// The external identifier of the user to connect.
	//
	// On create if this is empty a new external user will
	// be created if the connection supports it.
	ExternalId string
}

func (n *usersImpl) Connect(ctx context.Context, userId string, input *UserConnectInput) (*adminv1.User, error) {
	req := internal.NewRequest(
		"admin.users.connect",
		"POST",
		fmt.Sprintf("/admin/v1/users/%s:connect",
			url.PathEscape(userId),
		),
	)

	body := map[string]any{}

	if input != nil {
		if !internal.IsEmpty(input.ConnectionId) {
			body["connectionId"] = input.ConnectionId
		}
		if !internal.IsEmpty(input.ExternalId) {
			body["externalId"] = input.ExternalId
		}
	}

	req.SetBody(body)

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &adminv1.User{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// UserDisconnectInput is the input param for the Disconnect method.
type UserDisconnectInput struct {
	// The identifier of the connection.
	ConnectionId string
	// Whether to attempt to delete the external account and all
	// it's associated data (e.g. Stripe Customer object).
	DeleteExternalAccount bool
}

func (n *usersImpl) Disconnect(ctx context.Context, userId string, input *UserDisconnectInput) (*adminv1.User, error) {
	req := internal.NewRequest(
		"admin.users.disconnect",
		"POST",
		fmt.Sprintf("/admin/v1/users/%s:disconnect",
			url.PathEscape(userId),
		),
	)

	body := map[string]any{}

	if input != nil {
		if !internal.IsEmpty(input.ConnectionId) {
			body["connectionId"] = input.ConnectionId
		}
		if !internal.IsEmpty(input.DeleteExternalAccount) {
			body["deleteExternalAccount"] = input.DeleteExternalAccount
		}
	}

	req.SetBody(body)

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &adminv1.User{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// UserImportAccountInput is the input param for the ImportAccount method.
type UserImportAccountInput struct {
}

func (n *usersImpl) ImportAccount(ctx context.Context, userId string, input *UserImportAccountInput) (*adminv1.User, error) {
	req := internal.NewRequest(
		"admin.users.importAccount",
		"POST",
		fmt.Sprintf("/admin/v1/users/%s:import",
			url.PathEscape(userId),
		),
	)
	req.SetIdempotent(true)

	body := map[string]any{}

	req.SetBody(body)

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &adminv1.User{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// UserCreateApiSessionInput is the input param for the CreateApiSession method.
type UserCreateApiSessionInput struct {
}

func (n *usersImpl) CreateApiSession(ctx context.Context, userId string, input *UserCreateApiSessionInput) (*adminv1.CreateApiSessionResponse, error) {
	req := internal.NewRequest(
		"admin.users.createApiSession",
		"POST",
		fmt.Sprintf("/admin/v1/users/%s:createApiSession",
			url.PathEscape(userId),
		),
	)
	req.SetIdempotent(true)

	body := map[string]any{}

	req.SetBody(body)

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &adminv1.CreateApiSessionResponse{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// UserCreatePortalSessionInput is the input param for the CreatePortalSession method.
type UserCreatePortalSessionInput struct {
	// The portal URL, this is the target URL on the portal site.
	//
	// If not defined the root URL for the portal will be used.
	PortalUrl string
	// The return URL, this is where the user should be sent after they
	// exit the Portal.
	//
	// If not defined the app URL will be used.
	ReturnUrl string
}

func (n *usersImpl) CreatePortalSession(ctx context.Context, userId string, input *UserCreatePortalSessionInput) (*adminv1.CreatePortalSessionResponse, error) {
	req := internal.NewRequest(
		"admin.users.createPortalSession",
		"POST",
		fmt.Sprintf("/admin/v1/users/%s:createPortalSession",
			url.PathEscape(userId),
		),
	)
	req.SetIdempotent(true)

	body := map[string]any{}

	if input != nil {
		if !internal.IsEmpty(input.PortalUrl) {
			body["portalUrl"] = input.PortalUrl
		}
		if !internal.IsEmpty(input.ReturnUrl) {
			body["returnUrl"] = input.ReturnUrl
		}
	}

	req.SetBody(body)

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &adminv1.CreatePortalSessionResponse{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}
