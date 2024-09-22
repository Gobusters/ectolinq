package ectolinq

// Keys Returns a slice of keys from the map
func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// Returns a slice of values from the map
func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

// MapEntry is a key-value pair of a map
type MapEntry[K comparable, V any] struct {
	Key   K
	Value V
}

// Entries Returns a slice of key-value pairs from the map
func Entries[K comparable, V any](m map[K]V) []MapEntry[K, V] {
	entries := make([]MapEntry[K, V], 0, len(m))
	for k, v := range m {
		entries = append(entries, MapEntry[K, V]{Key: k, Value: v})
	}
	return entries
}

// MapFromEntries Returns a map from a slice of key-value pairs
func MapFromEntries[K comparable, V any](entries []MapEntry[K, V]) map[K]V {
	m := make(map[K]V)
	for _, entry := range entries {
		m[entry.Key] = entry.Value
	}
	return m
}

// FilterMap Returns a map with the elements that satisfy the predicate
func FilterMap[K comparable, V any](m map[K]V, predicate func(key K, value V) bool) map[K]V {
	filtered := make(map[K]V)

	for k, v := range m {
		if predicate(k, v) {
			filtered[k] = v
		}
	}
	return filtered
}

// MapValues returns a new map with the function applied to each value
func MapValues[K comparable, V any](m map[K]V, f func(V) V) map[K]V {
	newMap := make(map[K]V)
	for k, v := range m {
		newMap[k] = f(v)
	}
	return newMap
}

// MapKeys returns a new map with the function applied to each key
func MapKeys[K comparable, V any](m map[K]V, f func(K) K) map[K]V {
	newMap := make(map[K]V)
	for k, v := range m {
		newMap[f(k)] = v
	}
	return newMap
}

// MapEntries returns a new map with the entries mapped by the function
func MapEntries[K comparable, V any](m map[K]V, f func(key K, value V) (K, V)) map[K]V {
	newMap := make(map[K]V)
	for k, v := range m {
		newK, newV := f(k, v)
		newMap[newK] = newV
	}
	return newMap
}

// FindValue returns the value of the first element that satisfies the predicate
func FindValue[K comparable, V any](m map[K]V, predicate func(key K, value V) bool) (V, bool) {
	var zero V
	for k, v := range m {
		if predicate(k, v) {
			return v, true
		}
	}
	return zero, false
}

// FindKey returns the key of the first element that satisfies the predicate
func FindKey[K comparable, V any](m map[K]V, predicate func(key K, value V) bool) (K, bool) {
	var zero K
	for k, v := range m {
		if predicate(k, v) {
			return k, true
		}
	}

	return zero, false
}

// ReduceEntries reduces the map to a single value
func ReduceEntries[K comparable, V any, R any](m map[K]V, initial R, f func(acc R, key K, value V) R) R {
	result := initial
	for k, v := range m {
		result = f(result, k, v)
	}
	return result
}

// ForEachEntry executes a provided function once for each map entry
func ForEachEntry[K comparable, V any](m map[K]V, f func(key K, value V)) {
	for k, v := range m {
		f(k, v)
	}
}
