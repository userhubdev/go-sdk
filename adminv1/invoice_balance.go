// Code generated. DO NOT EDIT.

package adminv1

// The account balance as of the time the invoice
// was finalized.
type InvoiceBalance struct {
	// The starting balance of the account.
	StartAmount string `json:"startAmount"`
	// The ending balance of the account.
	EndAmount string `json:"endAmount"`
	// The amount applied to the invoice from the balance.
	//
	// A negative amount means a debit from the account balance.
	// A positive amount means a credit to the account balance.
	AppliedAmount string `json:"appliedAmount"`
}
