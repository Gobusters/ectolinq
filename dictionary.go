package ectolinq

// Dictionary is a wrapper around a map that provides additional functionality
type Dictionary[T any] struct {
	values map[string]T
}

// NewDictionary creates a new dictionary
func NewDictionary[T any]() *Dictionary[T] {
	return &Dictionary[T]{
		values: make(map[string]T),
	}
}

// Add a constructor that takes an initial capacity
func NewDictionaryWithCapacity[T any](capacity int) *Dictionary[T] {
	return &Dictionary[T]{
		values: make(map[string]T, capacity),
	}
}

// ToDictionary creates a new dictionary from a map
// m: The map to create the dictionary from
func ToDictionary[T any](m map[string]T) *Dictionary[T] {
	return &Dictionary[T]{
		values: m,
	}
}

// Get returns the value in the dictionary for the given key
// key: The key to get the value for
func (d *Dictionary[T]) Get(key string) (T, bool) {
	value, ok := d.values[key]

	return value, ok
}

// Set sets the value in the dictionary for the given key
// key: The key to set the value for
// value: The value to set
func (d *Dictionary[T]) Set(key string, value T) {
	d.values[key] = value
}

// Keys returns the keys in the dictionary
func (d *Dictionary[T]) Keys() []string {
	keys := make([]string, len(d.values))
	i := 0
	for key := range d.values {
		keys[i] = key
		i++
	}

	return keys
}

// ContainsKey returns if the dictionary contains the given key
// key: The key to check for
func (d *Dictionary[T]) ContainsKey(key string) bool {
	_, ok := d.values[key]
	return ok
}

// ContainsValue returns if the dictionary contains the given value
// value: The value to check for
func (d *Dictionary[T]) ContainsValue(value T) bool {
	for _, v := range d.values {
		if Equals(v, value) {
			return true
		}
	}

	return false
}

// ContainsWhere returns if the dictionary contains a value that satisfies the given predicate
// fn: The predicate to check for
func (d *Dictionary[T]) ContainsWhere(fn func(string, T) bool) bool {
	for key, value := range d.values {
		if fn(key, value) {
			return true
		}
	}

	return false
}

// Remove removes the value in the dictionary for the given key
// key: The key to remove the value for
func (d *Dictionary[T]) Remove(key string) {
	delete(d.values, key)
}

// RemoveValue removes the given value from the dictionary
// value: The value to remove
func (d *Dictionary[T]) RemoveValue(value T) {
	for key, v := range d.values {
		if Equals(v, value) {
			delete(d.values, key)
		}
	}
}

// RemoveWhere removes the value in the dictionary that satisfies the given predicate
// fn: The predicate to check for
func (d *Dictionary[T]) RemoveWhere(fn func(string, T) bool) {
	for key, value := range d.values {
		if fn(key, value) {
			delete(d.values, key)
		}
	}
}

// Modify the Merge method to return the modified dictionary
func (d *Dictionary[T]) Merge(dicts ...Dictionary[T]) *Dictionary[T] {
	for _, dict := range dicts {
		for key, value := range dict.values {
			d.values[key] = value
		}
	}
	return d
}

// MergeMaps merges the given maps into the dictionary
// maps: The maps to merge
func (d *Dictionary[T]) MergeMaps(maps ...map[string]T) {
	for _, m := range maps {
		for key, value := range m {
			d.values[key] = value
		}
	}
}

// Optimize the Clear method
func (d *Dictionary[T]) Clear() {
	d.values = make(map[string]T)
}

// Count returns the number of values in the dictionary
func (d *Dictionary[T]) Count() int {
	return len(d.values)
}

// ToArray returns the values in the dictionary as an array
func (d *Dictionary[T]) ToArray() []T {
	values := make([]T, len(d.values))
	i := 0
	for _, value := range d.values {
		values[i] = value
		i++
	}

	return values
}

// ToList returns the values in the dictionary as a list
func (d *Dictionary[T]) ToList() List[T] {
	return d.ToArray()
}

// ToMap returns the dictionary as a map
func (d *Dictionary[T]) ToMap() map[string]T {
	return d.values
}

// Add a method to get all values
func (d *Dictionary[T]) Values() []T {
	return d.ToArray()
}
