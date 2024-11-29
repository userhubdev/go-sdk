// Code generated. DO NOT EDIT.

package adminv1

// The Amazon Cognito connection data.
type AmazonCognitoConnection struct {
	// The Amazon Cognito user pool ID.
	UserPoolId string `json:"userPoolId"`
	// The Amazon region.
	Region string `json:"region"`
	// The Amazon access key ID.
	AccessKeyId string `json:"accessKeyId"`
	// The Amazon access key secret.
	AccessKeySecret string `json:"accessKeySecret"`
	// OpenID Connect (OIDC) configuration.
	//
	// If configured, this can be used instead of implementing a
	// Portal callback.
	Oidc *OidcConfig `json:"oidc"`
}
