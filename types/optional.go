package types

// Optional is an optional value.
type Optional[T any] struct {
	Present bool
	Value   T
}
