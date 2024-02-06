// Code generated. DO NOT EDIT.

package adminv1

// The plan products.
type PlanItem struct {
	// The product associated with the item.
	Product *Product `json:"product"`
	// The price associated with them item.
	Price *Price `json:"price"`
	// The plan item type.
	Type string `json:"type"`
}
