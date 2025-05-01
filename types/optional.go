package types

import (
	"bytes"
	"encoding/json"
)

// Optional is an optional value.
type Optional[T any] struct {
	Present bool
	Value   T
}

// UnmarshalJSON unmarshal JSON data and sets Present to false when the input is `null`.
func (o *Optional[T]) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		var empty T
		o.Value = empty
		o.Present = false
		return nil
	}

	o.Present = true
	return json.Unmarshal(data, &o.Value)
}

// MarshalJSON returns `null` if the value isn't present.
func (o Optional[T]) MarshalJSON() ([]byte, error) {
	if !o.Present {
		return []byte("null"), nil
	}

	return json.Marshal(o.Value)
}
