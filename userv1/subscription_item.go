// Code generated. DO NOT EDIT.

package userv1

// The subscription items.
type SubscriptionItem struct {
	// The identifier of the item.
	Id string `json:"id"`
	// The details of the associated product.
	Product *Product `json:"product"`
	// The details of the associated price.
	Price *Price `json:"price"`
	// The quantity for the item.
	Quantity int32 `json:"quantity"`
}
