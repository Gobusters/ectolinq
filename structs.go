package ectolinq

import (
	"fmt"
	"reflect"
	"strings"
)

// Equals returns if two values are equal
// a: The first value
// b: The second value
func Equals[T any](a T, b T) bool {
	return reflect.DeepEqual(a, b)
}

// IsEmpty returns if a value is empty
// val: The value to check
func IsEmpty[T any](val T) bool {
	return reflect.DeepEqual(val, reflect.Zero(reflect.TypeOf(val)).Interface())
}

// IfEmpty returns a default value if a value is empty
// val: The value to check
// def: The default value
func IfEmpty[T any](val T, def T) T {
	if IsEmpty(val) {
		return def
	}

	return val
}

// Get returns the value of a field in a struct
// s: The struct to get the value from
// path: The path to the field
func Get(s any, path string) (any, error) {
	r := reflect.ValueOf(s)

	// If s is a pointer, get the value it points to
	if r.Kind() == reflect.Ptr {
		r = r.Elem()
	}

	// Check if we indeed have a struct
	if r.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected a struct or a pointer to a struct")
	}

	pathParts := strings.Split(path, ".")
	for _, part := range pathParts {
		r = r.FieldByName(part)
		if !r.IsValid() {
			return nil, fmt.Errorf("field not found in path: %s", part)
		}

		// If it's a pointer, resolve it
		if r.Kind() == reflect.Ptr {
			r = r.Elem()
		}
	}

	return r.Interface(), nil
}
