// Code generated. DO NOT EDIT.

package apiv1

// The field view options..
type FieldView struct {
	// Whether the field is included in the view.
	//
	// THis overrides the message default.
	Include string `json:"include"`
	// Whether the field is excluded from the view.
	//
	// THis overrides the message default.
	Exclude string `json:"exclude"`
	// The referenced type's view.
	Type string `json:"type"`
}
