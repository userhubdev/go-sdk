// Code generated. DO NOT EDIT.

package adminv1

// The products included in the plan.
type PlanItem struct {
	// The plan item type.
	Type string `json:"type"`
	// The product associated with the item.
	Product *Product `json:"product"`
	// The price associated with the item.
	Price *Price `json:"price"`
}
