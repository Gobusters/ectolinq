package pointer

import "reflect"

func IsPointer[T any](value T) bool {
	return reflect.TypeOf(value).Kind() == reflect.Pointer
}

func Safe[T any](value *T) T {
	if value == nil {
		return *new(T)
	}
	return *value
}

func New[T any](value T) *T {
	return &value
}
