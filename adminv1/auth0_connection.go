// Code generated. DO NOT EDIT.

package adminv1

// The Auth0 connection data.
type Auth0Connection struct {
	// The Auth0 domain.
	Domain string `json:"domain"`
	// The Auth0 client ID.
	ClientId string `json:"clientId"`
	// The Auth0 client secret.
	ClientSecret string `json:"clientSecret"`
	// OpenID Connect (OIDC) configuration.
	//
	// If configured, this can be used instead of implementing a
	// Portal callback.
	Oidc *OidcConfig `json:"oidc"`
}
