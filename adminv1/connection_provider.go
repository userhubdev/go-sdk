// Code generated. DO NOT EDIT.

package adminv1

// The functionality the connection provides (e.g. `BILLING`).
type ConnectionProvider struct {
	// The provider type.
	Type string `json:"type"`
	// Whether the connection is the default for the provider type.
	Default bool `json:"default"`
}
