package ectolinq

import "sync"

// ConcurrentDictionary is a thread-safe dictionary
// utilizes a read-write mutex to allow multiple readers or a single writer
type ConcurrentDictionary[T any] struct {
	values *Dictionary[T]
	mutex  sync.RWMutex
}

// NewConcurrentDictionary creates a new dictionary
func NewConcurrentDictionary[T any]() *ConcurrentDictionary[T] {
	return &ConcurrentDictionary[T]{
		values: NewDictionary[T](),
		mutex:  sync.RWMutex{},
	}
}

// ToConcurrentDictionary creates a new dictionary from a map
// m: The map to create the dictionary from
func ToConcurrentDictionary[T any](m map[string]T) *ConcurrentDictionary[T] {
	return &ConcurrentDictionary[T]{
		values: ToDictionary[T](m),
		mutex:  sync.RWMutex{},
	}
}

// Get returns the value in the dictionary for the given key
// key: The key to get the value for
func (d *ConcurrentDictionary[T]) Get(key string) (T, bool) {
	d.mutex.RLock()
	defer d.mutex.RUnlock()
	return d.values.Get(key)
}

// Set sets the value in the dictionary for the given key
// key: The key to set the value for
// value: The value to set
func (d *ConcurrentDictionary[T]) Set(key string, value T) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	d.values.Set(key, value)
}

// Keys returns the keys in the dictionary
func (d *ConcurrentDictionary[T]) Keys() []string {
	d.mutex.RLock()
	defer d.mutex.RUnlock()
	return d.values.Keys()
}

// ContainsKey returns if the dictionary contains the given key
// key: The key to check for
func (d *ConcurrentDictionary[T]) ContainsKey(key string) bool {
	d.mutex.RLock()
	defer d.mutex.RUnlock()
	return d.values.ContainsKey(key)
}

// ContainsValue returns if the dictionary contains the given value
// value: The value to check for
func (d *ConcurrentDictionary[T]) ContainsValue(value T) bool {
	d.mutex.RLock()
	defer d.mutex.RUnlock()
	return d.values.ContainsValue(value)
}

// ContainsWhere returns if the dictionary contains a value that satisfies the given predicate
// fn: The predicate to check for
func (d *ConcurrentDictionary[T]) ContainsWhere(fn func(string, T) bool) bool {
	d.mutex.RLock()
	defer d.mutex.RUnlock()
	return d.values.ContainsWhere(fn)
}

// Remove removes the value in the dictionary for the given key
// key: The key to remove the value for
func (d *ConcurrentDictionary[T]) Remove(key string) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	d.values.Remove(key)
}

// RemoveValue removes the given value from the dictionary
// value: The value to remove
func (d *ConcurrentDictionary[T]) RemoveValue(value T) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	d.values.RemoveValue(value)
}

// RemoveWhere removes the value in the dictionary that satisfies the given predicate
// fn: The predicate to check for
func (d *ConcurrentDictionary[T]) RemoveWhere(fn func(string, T) bool) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	d.values.RemoveWhere(fn)
}

// Clear removes all values from the dictionary
func (d *ConcurrentDictionary[T]) Clear() {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	d.values.Clear()
}

// Count returns the number of values in the dictionary
func (d *ConcurrentDictionary[T]) Count() int {
	d.mutex.RLock()
	defer d.mutex.RUnlock()
	return d.values.Count()
}

// ToArray returns the values in the dictionary as an array
func (d *ConcurrentDictionary[T]) ToArray() []T {
	d.mutex.RLock()
	defer d.mutex.RUnlock()
	return d.values.ToArray()
}

// ToList returns the values in the dictionary as a list
func (d *ConcurrentDictionary[T]) ToList() List[T] {
	d.mutex.RLock()
	defer d.mutex.RUnlock()
	return d.values.ToList() // Use values.ToList() instead of ToArray()
}

// ToMap returns the dictionary as a map
func (d *ConcurrentDictionary[T]) ToMap() map[string]T {
	d.mutex.RLock()
	defer d.mutex.RUnlock()
	return d.values.ToMap()
}

// Merge merges the given dictionaries into the dictionary
// dicts: The dictionaries to merge
func (d *ConcurrentDictionary[T]) Merge(dicts ...*ConcurrentDictionary[T]) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	for _, dict := range dicts {
		dict.mutex.RLock()
		d.values.Merge(*dict.values)
		dict.mutex.RUnlock()
	}
}

// MergeMaps merges the given maps into the dictionary
// maps: The maps to merge
func (d *ConcurrentDictionary[T]) MergeMaps(maps ...map[string]T) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	d.values.MergeMaps(maps...)
}
