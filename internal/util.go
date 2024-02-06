package internal

func IsEmpty[T comparable](s T) bool {
	var empty T
	if s == empty {
		return true
	}
	if e, ok := any(s).(interface {
		IsZero() bool
	}); ok {
		return e.IsZero()
	}
	return false
}
