// Code generated. DO NOT EDIT.

package adminv1

// The subscription items.
type UpdateSubscriptionItemsRequestItem struct {
	// The product identifier.
	//
	// If this is empty and the user ID is set, the default
	// seat will be used.
	ProductId string `json:"productId"`
	// The member user ID of the organization member. This can
	// only be specified for seat items.
	UserId string `json:"userId"`
	// The quantity for the item.
	//
	// If this is `0` the item will be removed.
	Quantity int32 `json:"quantity"`
}
