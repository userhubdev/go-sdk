// Code generated. DO NOT EDIT.

package adminv1

// The subscription seat.
type SubscriptionSeatInfo struct {
	// Whether a seat can be assigned or reserved.
	State string `json:"state"`
	// The code that best describes the reason for the state.
	StateReason string `json:"stateReason"`
	// The seat product.
	Product *Product `json:"product"`
	// The number of seats expected to be invoiced for the current billing period.
	//
	// This might be less than the total quantity while a subscription change
	// is pending or if the subscription is over-provisioned.
	CurrentPeriodQuantity int32 `json:"currentPeriodQuantity"`
	// The number of seats scheduled to appear on the next invoice.
	//
	// This will only be set when different from current period quantity.
	NextPeriodQuantity int32 `json:"nextPeriodQuantity"`
	// The number of seats currently assigned.
	AssignedQuantity int32 `json:"assignedQuantity"`
	// The number of seats not assigned.
	UnassignedQuantity int32 `json:"unassignedQuantity"`
	// The number of seats currently reserved because of pending invitations.
	//
	// These can be made available by cancelling outstanding invitations.
	ReservedQuantity int32 `json:"reservedQuantity"`
	// The number of seats which can be assigned or reserved.
	AvailableQuantity int32 `json:"availableQuantity"`
	// The total number of seats for the current period.
	TotalQuantity int32 `json:"totalQuantity"`
}
