package ectolinq

type Dictionary[T any] map[string]T

// Get returns the value in the dictionary for the given key
// key: The key to get the value for
func (d Dictionary[T]) Get(key string) T {
	var value T
	value, _ = d[key]

	return value
}

// Set sets the value in the dictionary for the given key
// key: The key to set the value for
// value: The value to set
func (d Dictionary[T]) Set(key string, value T) {
	d[key] = value
}

// Keys returns the keys in the dictionary
func (d Dictionary[T]) Keys() []string {
	keys := make([]string, len(d))
	i := 0
	for key := range d {
		keys[i] = key
		i++
	}

	return keys
}

// ContainsKey returns if the dictionary contains the given key
// key: The key to check for
func (d Dictionary[T]) ContainsKey(key string) bool {
	_, ok := d[key]
	return ok
}

// ContainsValue returns if the dictionary contains the given value
// value: The value to check for
func (d Dictionary[T]) ContainsValue(value T) bool {
	for _, v := range d {
		if Equals(v, value) {
			return true
		}
	}

	return false
}

// ContainsWhere returns if the dictionary contains a value that satisfies the given predicate
// fn: The predicate to check for
func (d Dictionary[T]) ContainsWhere(fn func(string, T) bool) bool {
	for key, value := range d {
		if fn(key, value) {
			return true
		}
	}

	return false
}

// Remove removes the value in the dictionary for the given key
// key: The key to remove the value for
func (d Dictionary[T]) Remove(key string) {
	delete(d, key)
}

// RemoveValue removes the given value from the dictionary
// value: The value to remove
func (d Dictionary[T]) RemoveValue(value T) {
	for key, v := range d {
		if Equals(v, value) {
			delete(d, key)
		}
	}
}

// RemoveWhere removes the value in the dictionary that satisfies the given predicate
// fn: The predicate to check for
func (d Dictionary[T]) RemoveWhere(fn func(string, T) bool) {
	for key, value := range d {
		if fn(key, value) {
			delete(d, key)
		}
	}
}

// Clear removes all values from the dictionary
func (d Dictionary[T]) Clear() {
	for key := range d {
		delete(d, key)
	}
}

// Count returns the number of values in the dictionary
func (d Dictionary[T]) Count() int {
	return len(d)
}

// ToArray returns the values in the dictionary as an array
func (d Dictionary[T]) ToArray() []T {
	values := make([]T, len(d))
	i := 0
	for _, value := range d {
		values[i] = value
		i++
	}

	return values
}

// ToList returns the values in the dictionary as a list
func (d Dictionary[T]) ToList() List[T] {
	return d.ToArray()
}
