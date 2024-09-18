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
	defer func() {
		recover()
	}()
	return reflect.DeepEqual(a, b)
}

// IsEmpty returns if a value is empty
// val: The value to check
func IsEmpty[T any](val T) bool {
	defer func() {
		recover()
	}()
	rv := reflect.ValueOf(val)
	switch rv.Kind() {
	case reflect.Slice, reflect.Map, reflect.String:
		return rv.Len() == 0
	case reflect.Chan:
		return rv.IsNil() || rv.Cap() == 0
	case reflect.Ptr, reflect.Func, reflect.Interface:
		return rv.IsNil()
	case reflect.Invalid:
		return true
	default:
		return reflect.DeepEqual(val, reflect.Zero(reflect.TypeOf(val)).Interface())
	}
}

// IfEmpty returns a default value if a value is empty
// val: The value to check
// def: The default value
func IfEmpty[T any](val T, def T) T {
	defer func() {
		if r := recover(); r != nil {
			val = def
		}
	}()
	if IsEmpty(val) {
		return def
	}
	return val
}

// Get returns the value of a field in a struct
// s: The struct to get the value from
// path: The path to the field
func Get(s any, path string) (any, error) {
	if path == "" {
		return nil, fmt.Errorf("path cannot be empty")
	}

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

// Set sets the value of a field in a struct
// s: The struct to set the value in
// path: The path to the field
// value: The value to set
func Set(s any, path string, value any) error {
	if path == "" {
		return fmt.Errorf("path cannot be empty")
	}

	r := reflect.ValueOf(s)
	if r.Kind() != reflect.Ptr || r.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("expected a pointer to a struct")
	}

	r = r.Elem()
	pathParts := strings.Split(path, ".")
	for i, part := range pathParts {
		if i == len(pathParts)-1 {
			field := r.FieldByName(part)
			if !field.IsValid() {
				return fmt.Errorf("field not found: %s", part)
			}
			if !field.CanSet() {
				return fmt.Errorf("cannot set field: %s", part)
			}
			field.Set(reflect.ValueOf(value))
			return nil
		}
		r = r.FieldByName(part)
		if !r.IsValid() {
			return fmt.Errorf("field not found in path: %s", part)
		}
		if r.Kind() == reflect.Ptr {
			r = r.Elem()
		}
	}
	return nil
}

// HasField checks if a struct has a specific field
// s: The struct to check
// fieldName: The name of the field to check
func HasField(s any, fieldName string) bool {
	r := reflect.ValueOf(s)
	if r.Kind() == reflect.Ptr {
		r = r.Elem()
	}
	if r.Kind() != reflect.Struct {
		return false
	}
	return r.FieldByName(fieldName).IsValid()
}

// GetFieldNames returns a slice of all field names in a struct
// s: The struct to get the field names from
func GetFieldNames(s any) []string {
	r := reflect.ValueOf(s)
	if r.Kind() == reflect.Ptr {
		r = r.Elem()
	}
	if r.Kind() != reflect.Struct {
		return nil
	}
	var names []string
	for i := 0; i < r.NumField(); i++ {
		names = append(names, r.Type().Field(i).Name)
	}
	return names
}

// ToMap converts a struct to a map[string]interface{}
func ToMap(s any) (map[string]interface{}, error) {
	r := reflect.ValueOf(s)
	if r.Kind() == reflect.Ptr {
		r = r.Elem()
	}
	if r.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected a struct or a pointer to a struct")
	}
	m := make(map[string]interface{})
	for i := 0; i < r.NumField(); i++ {
		m[r.Type().Field(i).Name] = r.Field(i).Interface()
	}
	return m, nil
}

// FromMap creates a struct from a map[string]interface{}
// m: The map to create the struct from
// s: The struct to set the values in
func FromMap(m map[string]interface{}, s any) error {
	r := reflect.ValueOf(s)
	if r.Kind() != reflect.Ptr || r.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("expected a pointer to a struct")
	}
	r = r.Elem()
	for k, v := range m {
		field := r.FieldByName(k)
		if !field.IsValid() {
			continue // Skip fields that don't exist in the struct
		}
		if !field.CanSet() {
			return fmt.Errorf("cannot set field: %s", k)
		}
		field.Set(reflect.ValueOf(v))
	}
	return nil
}
