// Code generated. DO NOT EDIT.

package adminapi

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/userhubdev/go-sdk/adminv1"
	"github.com/userhubdev/go-sdk/apiv1"
	"github.com/userhubdev/go-sdk/commonv1"
	"github.com/userhubdev/go-sdk/internal"
	"github.com/userhubdev/go-sdk/types"
)

type Organizations interface {
	// Lists organizations.
	List(ctx context.Context, input *OrganizationListInput) (*adminv1.ListOrganizationsResponse, error)
	// Creates a new organization.
	Create(ctx context.Context, input *OrganizationCreateInput) (*adminv1.Organization, error)
	// Retrieves specified organization.
	Get(ctx context.Context, organizationId string, input *OrganizationGetInput) (*adminv1.Organization, error)
	// Updates specified organization.
	Update(ctx context.Context, organizationId string, input *OrganizationUpdateInput) (*adminv1.Organization, error)
	// Marks specified organization for deletion.
	Delete(ctx context.Context, organizationId string, input *OrganizationDeleteInput) (*adminv1.Organization, error)
	// Un-marks specified organization for deletion.
	Undelete(ctx context.Context, organizationId string, input *OrganizationUndeleteInput) (*adminv1.Organization, error)
	// Connect specified organization to external account.
	Connect(ctx context.Context, organizationId string, input *OrganizationConnectInput) (*adminv1.Organization, error)
	// Disconnect specified organization from external account.
	//
	// This will delete all the data associated with the connected account, including
	// payment methods, invoices, and subscriptions.
	//
	// If the delete external account option is enabled it will also attempt to delete
	// the external account (e.g. Stripe Customer object).
	//
	// WARNING: This can irreversibly destroy data and should be
	// used with extreme caution.
	Disconnect(ctx context.Context, organizationId string, input *OrganizationDisconnectInput) (*adminv1.Organization, error)
	// Lists organization members.
	ListMembers(ctx context.Context, organizationId string, input *OrganizationListMembersInput) (*adminv1.ListMembersResponse, error)
	// Creates a new organization member.
	AddMember(ctx context.Context, organizationId string, input *OrganizationAddMemberInput) (*adminv1.Member, error)
	// Retrieves specified organization member.
	GetMember(ctx context.Context, organizationId string, userId string, input *OrganizationGetMemberInput) (*adminv1.Member, error)
	// Updates specified organization member.
	UpdateMember(ctx context.Context, organizationId string, userId string, input *OrganizationUpdateMemberInput) (*adminv1.Member, error)
	// Deletes specified organization member.
	RemoveMember(ctx context.Context, organizationId string, userId string, input *OrganizationRemoveMemberInput) (*apiv1.EmptyResponse, error)
}

type organizationsImpl struct {
	transport internal.Transport
}

// OrganizationListInput is the input param for the List method.
type OrganizationListInput struct {
	// Filter the results by display name.
	//
	// To enable prefix filtering append `*` to the end of the value
	// and ensure you provide at least 3 characters excluding the
	// wildcard.
	//
	// This filter is case-insensitivity.
	DisplayName string
	// Filter the results by email address.
	//
	// To enable prefix filtering append `*` to the end of the value
	// and ensure you provide at least 3 characters excluding the
	// wildcard.
	//
	// This filter is case-insensitivity.
	Email string
	// The maximum number of organizations to return. The API may return fewer than
	// this value.
	//
	// If unspecified, at most 20 organizations will be returned.
	// The maximum value is 100; values above 100 will be coerced to 100.
	PageSize int32
	// A page token, received from a previous list organizations call.
	// Provide this to retrieve the subsequent page.
	//
	// When paginating, all other parameters provided to list organizations must match
	// the call that provided the page token.
	PageToken string
	// A comma-separated list of fields to order by.
	//
	// Supports:
	// - `displayName asc`
	// - `email asc`
	// - `signupTime desc`
	// - `createTime desc`
	// - `deleteTime desc`
	OrderBy string
	// Whether to show deleted organizations.
	ShowDeleted bool
	// The organization view to return in the results.
	//
	// This defaults to the `BASIC` view.
	View string
}

func (n *organizationsImpl) List(ctx context.Context, input *OrganizationListInput) (*adminv1.ListOrganizationsResponse, error) {
	req := internal.NewRequest(
		"admin.organizations.list",
		"GET",
		"/admin/v1/organizations",
	)
	req.SetIdempotent(true)

	if input != nil {
		if !internal.IsEmpty(input.DisplayName) {
			req.SetQuery("displayName", input.DisplayName)
		}
		if !internal.IsEmpty(input.Email) {
			req.SetQuery("email", input.Email)
		}
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

	model := &adminv1.ListOrganizationsResponse{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// OrganizationCreateInput is the input param for the Create method.
type OrganizationCreateInput struct {
	// The client defined unique identifier of the organization account.
	//
	// It is restricted to letters, numbers, underscores, and hyphens,
	// with the first character a letter or a number, and a 255
	// character maximum.
	//
	// ID's starting with `org_` are reserved.
	UniqueId string
	// The human-readable display name of the organization.
	//
	// The maximum length is 200 characters.
	DisplayName string
	// The email address of the organization.
	//
	// The maximum length is 320 characters.
	Email string
	// Whether the organization's email address has been verified.
	EmailVerified bool
	// The E164 phone number for the organization (e.g. `+12125550123`).
	PhoneNumber string
	// Whether the organization's phone number has been verified.
	PhoneNumberVerified bool
	// The photo/avatar URL of the organization.
	//
	// The maximum length is 2000 characters.
	ImageUrl string
	// The default ISO-4217 currency code for the organization (e.g. `USD`).
	CurrencyCode string
	// The IETF BCP-47 language code for the organization (e.g. `en`).
	LanguageCode string
	// The country/region code for the organization (e.g. `US`).
	RegionCode string
	// The IANA time zone for the organization (e.g. `America/New_York`).
	TimeZone string
	// The billing address for the organization.
	Address *commonv1.Address
	// The sign-up time for the organization.
	SignupTime time.Time
	// Whether the organization is disabled.
	Disabled bool
}

func (n *organizationsImpl) Create(ctx context.Context, input *OrganizationCreateInput) (*adminv1.Organization, error) {
	req := internal.NewRequest(
		"admin.organizations.create",
		"POST",
		"/admin/v1/organizations",
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

	model := &adminv1.Organization{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// OrganizationGetInput is the input param for the Get method.
type OrganizationGetInput struct {
}

func (n *organizationsImpl) Get(ctx context.Context, organizationId string, input *OrganizationGetInput) (*adminv1.Organization, error) {
	req := internal.NewRequest(
		"admin.organizations.get",
		"GET",
		fmt.Sprintf("/admin/v1/organizations/%s",
			url.PathEscape(organizationId),
		),
	)
	req.SetIdempotent(true)

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &adminv1.Organization{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// OrganizationUpdateInput is the input param for the Update method.
type OrganizationUpdateInput struct {
	// The client defined unique identifier of the organization account.
	//
	// It is restricted to letters, numbers, underscores, and hyphens,
	// with the first character a letter or a number, and a 255
	// character maximum.
	//
	// ID's starting with `org_` are reserved.
	UniqueId types.Optional[string]
	// The human-readable display name of the organization.
	//
	// The maximum length is 200 characters.
	DisplayName types.Optional[string]
	// The email address of the organization.
	//
	// The maximum length is 320 characters.
	Email types.Optional[string]
	// Whether the organization's email address has been verified.
	EmailVerified types.Optional[bool]
	// The E164 phone number for the organization (e.g. `+12125550123`).
	PhoneNumber types.Optional[string]
	// Whether the organization's phone number has been verified.
	PhoneNumberVerified types.Optional[bool]
	// The photo/avatar URL of the organization.
	//
	// The maximum length is 2000 characters.
	ImageUrl types.Optional[string]
	// The default ISO-4217 currency code for the organization (e.g. `USD`).
	CurrencyCode types.Optional[string]
	// The IETF BCP-47 language code for the organization (e.g. `en`).
	LanguageCode types.Optional[string]
	// The country/region code for the organization (e.g. `US`).
	RegionCode types.Optional[string]
	// The IANA time zone for the organization (e.g. `America/New_York`).
	TimeZone types.Optional[string]
	// The billing address for the organization.
	Address types.Optional[*commonv1.Address]
	// The sign-up time for the organization.
	SignupTime types.Optional[time.Time]
	// Whether the organization is disabled.
	Disabled types.Optional[bool]

	// If set to true, and the organization is not found, a new organization will be created.
	//
	// You must use the unique identifier for the identifier when this is true.
	AllowMissing bool
}

func (n *organizationsImpl) Update(ctx context.Context, organizationId string, input *OrganizationUpdateInput) (*adminv1.Organization, error) {
	req := internal.NewRequest(
		"admin.organizations.update",
		"PATCH",
		fmt.Sprintf("/admin/v1/organizations/%s",
			url.PathEscape(organizationId),
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

	model := &adminv1.Organization{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// OrganizationDeleteInput is the input param for the Delete method.
type OrganizationDeleteInput struct {
}

func (n *organizationsImpl) Delete(ctx context.Context, organizationId string, input *OrganizationDeleteInput) (*adminv1.Organization, error) {
	req := internal.NewRequest(
		"admin.organizations.delete",
		"DELETE",
		fmt.Sprintf("/admin/v1/organizations/%s",
			url.PathEscape(organizationId),
		),
	)
	req.SetIdempotent(true)

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &adminv1.Organization{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// OrganizationUndeleteInput is the input param for the Undelete method.
type OrganizationUndeleteInput struct {
}

func (n *organizationsImpl) Undelete(ctx context.Context, organizationId string, input *OrganizationUndeleteInput) (*adminv1.Organization, error) {
	req := internal.NewRequest(
		"admin.organizations.undelete",
		"POST",
		fmt.Sprintf("/admin/v1/organizations/%s:undelete",
			url.PathEscape(organizationId),
		),
	)
	req.SetIdempotent(true)

	body := map[string]any{}

	req.SetBody(body)

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &adminv1.Organization{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// OrganizationConnectInput is the input param for the Connect method.
type OrganizationConnectInput struct {
	// The identifier of the connection.
	ConnectionId string
	// The external identifier of the organization to connect.
	//
	// On create if this is empty a new external organization will
	// be created if the connection supports it.
	ExternalId string
}

func (n *organizationsImpl) Connect(ctx context.Context, organizationId string, input *OrganizationConnectInput) (*adminv1.Organization, error) {
	req := internal.NewRequest(
		"admin.organizations.connect",
		"POST",
		fmt.Sprintf("/admin/v1/organizations/%s:connect",
			url.PathEscape(organizationId),
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

	model := &adminv1.Organization{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// OrganizationDisconnectInput is the input param for the Disconnect method.
type OrganizationDisconnectInput struct {
	// The identifier of the connection.
	ConnectionId string
	// Whether to attempt to delete the external account and all
	// it's associated data (e.g. Stripe Customer object).
	DeleteExternalAccount bool
}

func (n *organizationsImpl) Disconnect(ctx context.Context, organizationId string, input *OrganizationDisconnectInput) (*adminv1.Organization, error) {
	req := internal.NewRequest(
		"admin.organizations.disconnect",
		"POST",
		fmt.Sprintf("/admin/v1/organizations/%s:disconnect",
			url.PathEscape(organizationId),
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

	model := &adminv1.Organization{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// OrganizationListMembersInput is the input param for the ListMembers method.
type OrganizationListMembersInput struct {
	// Filter the results by display name.
	//
	// To enable prefix filtering append `*` to the end of the value
	// and ensure you provide at least 3 characters excluding the
	// wildcard.
	//
	// This filter is case-insensitivity.
	DisplayName string
	// Filter the results by email address.
	//
	// To enable prefix filtering append `*` to the end of the value
	// and ensure you provide at least 3 characters excluding the
	// wildcard.
	//
	// This filter is case-insensitivity.
	Email string
	// Filter the results by a role identifier.
	RoleId string
	// The maximum number of members to return. The API may return fewer than
	// this value.
	//
	// If unspecified, at most 20 members will be returned.
	// The maximum value is 100; values above 100 will be coerced to 100.
	PageSize int32
	// A page token, received from a previous list members call.
	// Provide this to retrieve the subsequent page.
	//
	// When paginating, all other parameters provided to list members must match
	// the call that provided the page token.
	PageToken string
	// A comma-separated list of fields to order by.
	//
	// Supports:
	// - `displayName asc`
	// - `email asc`
	// - `createTime desc`
	OrderBy string
}

func (n *organizationsImpl) ListMembers(ctx context.Context, organizationId string, input *OrganizationListMembersInput) (*adminv1.ListMembersResponse, error) {
	req := internal.NewRequest(
		"admin.organizations.listMembers",
		"GET",
		fmt.Sprintf("/admin/v1/organizations/%s/members",
			url.PathEscape(organizationId),
		),
	)
	req.SetIdempotent(true)

	if input != nil {
		if !internal.IsEmpty(input.DisplayName) {
			req.SetQuery("displayName", input.DisplayName)
		}
		if !internal.IsEmpty(input.Email) {
			req.SetQuery("email", input.Email)
		}
		if !internal.IsEmpty(input.RoleId) {
			req.SetQuery("roleId", input.RoleId)
		}
		if !internal.IsEmpty(input.PageSize) {
			req.SetQuery("pageSize", input.PageSize)
		}
		if !internal.IsEmpty(input.PageToken) {
			req.SetQuery("pageToken", input.PageToken)
		}
		if !internal.IsEmpty(input.OrderBy) {
			req.SetQuery("orderBy", input.OrderBy)
		}
	}

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &adminv1.ListMembersResponse{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// OrganizationAddMemberInput is the input param for the AddMember method.
type OrganizationAddMemberInput struct {
	// The identifier of the user.
	UserId string
	// The identifier of the role.
	RoleId string
}

func (n *organizationsImpl) AddMember(ctx context.Context, organizationId string, input *OrganizationAddMemberInput) (*adminv1.Member, error) {
	req := internal.NewRequest(
		"admin.organizations.addMember",
		"POST",
		fmt.Sprintf("/admin/v1/organizations/%s/members",
			url.PathEscape(organizationId),
		),
	)

	body := map[string]any{}

	if input != nil {
		if !internal.IsEmpty(input.UserId) {
			body["userId"] = input.UserId
		}
		if !internal.IsEmpty(input.RoleId) {
			body["roleId"] = input.RoleId
		}
	}

	req.SetBody(body)

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &adminv1.Member{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// OrganizationGetMemberInput is the input param for the GetMember method.
type OrganizationGetMemberInput struct {
}

func (n *organizationsImpl) GetMember(ctx context.Context, organizationId string, userId string, input *OrganizationGetMemberInput) (*adminv1.Member, error) {
	req := internal.NewRequest(
		"admin.organizations.getMember",
		"GET",
		fmt.Sprintf("/admin/v1/organizations/%s/members/%s",
			url.PathEscape(organizationId),
			url.PathEscape(userId),
		),
	)
	req.SetIdempotent(true)

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &adminv1.Member{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// OrganizationUpdateMemberInput is the input param for the UpdateMember method.
type OrganizationUpdateMemberInput struct {
	// The identifier of the role.
	RoleId types.Optional[string]

	// If set to true, and the member is not found, a new member will be created.
	AllowMissing bool
}

func (n *organizationsImpl) UpdateMember(ctx context.Context, organizationId string, userId string, input *OrganizationUpdateMemberInput) (*adminv1.Member, error) {
	req := internal.NewRequest(
		"admin.organizations.updateMember",
		"PATCH",
		fmt.Sprintf("/admin/v1/organizations/%s/members/%s",
			url.PathEscape(organizationId),
			url.PathEscape(userId),
		),
	)
	req.SetIdempotent(true)

	body := map[string]any{}

	if input != nil {
		if !internal.IsEmpty(input.AllowMissing) {
			req.SetQuery("allowMissing", input.AllowMissing)
		}
		if input.RoleId.Present {
			body["roleId"] = input.RoleId.Value
		}
	}

	req.SetBody(body)

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &adminv1.Member{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// OrganizationRemoveMemberInput is the input param for the RemoveMember method.
type OrganizationRemoveMemberInput struct {
}

func (n *organizationsImpl) RemoveMember(ctx context.Context, organizationId string, userId string, input *OrganizationRemoveMemberInput) (*apiv1.EmptyResponse, error) {
	req := internal.NewRequest(
		"admin.organizations.removeMember",
		"DELETE",
		fmt.Sprintf("/admin/v1/organizations/%s/members/%s",
			url.PathEscape(organizationId),
			url.PathEscape(userId),
		),
	)

	res, err := n.transport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	model := &apiv1.EmptyResponse{}

	err = res.DecodeBody(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}
