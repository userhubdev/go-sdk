// Code generated. DO NOT EDIT.

package adminv1

// The Google Cloud Identity Platform/Firebase specific connection data.
type GoogleCloudIdentityPlatformConnection struct {
	// The service account key in JSON format.
	Credentials string `json:"credentials"`
	// The Google Cloud Identity Platform/Firebase project ID.
	//
	// This will default to the project ID in the service account key if
	// not specified.
	ProjectId string `json:"projectId"`
}
