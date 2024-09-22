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

// Merge combines two or more maps into a new map
func Merge[K comparable, V any](maps ...map[K]V) map[K]V {
	result := make(map[K]V)
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

// CountEntries returns the number of key-value pairs in the map that satisfy the predicate
func CountEntries[K comparable, V any](m map[K]V, predicate func(key K, value V) bool) int {
	count := 0
	for k, v := range m {
		if predicate(k, v) {
			count++
		}
	}
	return count
}

// GroupBy groups the map entries by a key selector function
func GroupBy[K comparable, V any, G comparable](m map[K]V, keySelector func(K, V) G) map[G][]MapEntry[K, V] {
	result := make(map[G][]MapEntry[K, V])
	for k, v := range m {
		group := keySelector(k, v)
		result[group] = append(result[group], MapEntry[K, V]{Key: k, Value: v})
	}
	return result
}

// Invert swaps the keys and values of the map
func Invert[K comparable, V comparable](m map[K]V) map[V]K {
	result := make(map[V]K)
	for k, v := range m {
		result[v] = k
	}
	return result
}

// Pick creates a new map with only the specified keys
func Pick[K comparable, V any](m map[K]V, keys ...K) map[K]V {
	result := make(map[K]V)
	for _, k := range keys {
		if v, ok := m[k]; ok {
			result[k] = v
		}
	}
	return result
}

// Omit creates a new map without the specified keys
func Omit[K comparable, V any](m map[K]V, keys ...K) map[K]V {
	result := make(map[K]V)
	for k, v := range m {
		if !ContainsValue(keys, k) {
			result[k] = v
		}
	}
	return result
}

// ContainsValue checks if a slice contains a specific element
func ContainsValue[T comparable](slice []T, element T) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}
