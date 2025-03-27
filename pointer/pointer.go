package pointer

import "reflect"

// IsPointer checks if the value is a pointer
func IsPointer[T any](value T) bool {
	return reflect.TypeOf(value).Kind() == reflect.Pointer
}

// Safe returns the value if it is not nil, otherwise it returns the zero value
func Safe[T any](value *T) T {
	if value == nil {
		return *new(T)
	}
	return *value
}

// New returns a new pointer to the value
func New[T any](value T) *T {
	return &value
}
