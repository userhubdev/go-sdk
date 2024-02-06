// Code generated. DO NOT EDIT.

package adminv1

// The request associated with event.
type EventRequest struct {
	// The IP address of the client/user.
	//
	// It's very likely this is the IP address of the
	// API client and not the end-user.
	IpAddress string `json:"ipAddress"`
	// The trace ID associated with the request.
	TraceId string `json:"traceId"`
}
