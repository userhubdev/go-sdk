// Code generated. DO NOT EDIT.

package operationsv1

import (
	"time"

	"github.com/userhubdev/go-sdk/commonv1"
)

// Operation is a long running background task.
type Operation struct {
	// The system-assigned identifier of the operation.
	Id string `json:"id"`
	// If the value is `false`, it means the operation is still in progress.
	// If `true`, the operation is completed, and either `error` or `response` is
	// available.
	Done bool `json:"done"`
	// The error result of the operation in case of failure.
	Error *OperationError `json:"error"`
	// The normal response of the operation in case of success.
	Response *commonv1.Any `json:"response"`
	// The creation time of the operation.
	CreateTime time.Time `json:"createTime"`
	// The last update time of the operation.
	UpdateTime time.Time `json:"updateTime"`
	// The deletion time of the operation.
	DeleteTime time.Time `json:"deleteTime"`
}
