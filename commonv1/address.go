// Code generated. DO NOT EDIT.

package commonv1

// A postal address (billing, mailing, etc...).
type Address struct {
	// The address lines.
	Lines []string `json:"lines"`
	// The city, district, suburb, town, or village.
	City string `json:"city"`
	// The state, country, province, or region.
	State string `json:"state"`
	// The ZIP or postal code.
	PostalCode string `json:"postalCode"`
	// The 2-letter country code.
	Country string `json:"country"`
}
