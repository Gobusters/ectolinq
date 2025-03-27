package ectolinq

import "reflect"

// Ternary returns the trueVal if the condition is true, otherwise it returns the falseVal
func Ternary[T any](condition bool, trueVal, falseVal T) T {
	if condition {
		return trueVal
	}
	return falseVal
}

// Default returns the value if it is not zero, otherwise it returns the default value
func Default[T any](value T, defaultValue T) T {
	if IsZero(value) {
		return defaultValue
	}
	return value
}

// IsZero checks if the value is zero
func IsZero[T any](value T) bool {
	return reflect.ValueOf(value).IsZero()
}
