// Code generated. DO NOT EDIT.

package userv1

// The products which the plan includes.
type PlanItem struct {
	// The product associated with the item.
	Product *Product `json:"product"`
	// The price associated with them item.
	Price *Price `json:"price"`
	// The plan item type.
	Type string `json:"type"`
}
