// Code generated. DO NOT EDIT.

package adminv1

// The subscription items.
type SubscriptionItem struct {
	// The identifier of the item.
	Id string `json:"id"`
	// The item product.
	Product *Product `json:"product"`
	// The item price.
	Price *Price `json:"price"`
	// The quantity of products.
	Quantity int32 `json:"quantity"`
}
