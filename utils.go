package ectolinq

import "reflect"

func Ternary[T any](condition bool, trueVal, falseVal T) T {
	if condition {
		return trueVal
	}
	return falseVal
}

func Default[T any](value T, defaultValue T) T {
	if IsZero(value) {
		return defaultValue
	}
	return value
}

func IsZero[T any](value T) bool {
	return reflect.ValueOf(value).IsZero()
}
