// Code generated. DO NOT EDIT.

package adminv1

// The subscription seat.
type SubscriptionSeatInfo struct {
	// The seat product.
	Product *Product `json:"product"`
	// The quantity which has been invoiced for the current billing period.
	//
	// This might be less than the total quantity while a subscription change
	// is pending or if the subscription is over-provisioned.
	CurrentPeriodQuantity int32 `json:"currentPeriodQuantity"`
	// The quantity scheduled to appear on the next invoice.
	//
	// This will only be set when different from current period quantity.
	NextPeriodQuantity int32 `json:"nextPeriodQuantity"`
	// The quantity currently in use.
	AssignedQuantity int32 `json:"assignedQuantity"`
	// The quantity available for use.
	UnassignedQuantity int32 `json:"unassignedQuantity"`
	// The sum of the assigned and unassigned quantities.
	TotalQuantity int32 `json:"totalQuantity"`
}
