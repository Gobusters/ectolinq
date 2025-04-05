package ectolinq

import (
	"math/rand"
	"sort"
)

// Find returns the first element in the slice that satisfies the predicate
// items: The slice to search
// predicate: The predicate to test each element against
func Find[T any](items []T, predicate func(T) bool) T {
	var result T

	for _, item := range items {
		if predicate(item) {
			return item
		}
	}

	return result
}

// Reverse reverses the order of the elements in the slice in-place
// items: The slice to reverse
func Reverse[T any](items []T) []T {
	// Consider using a more efficient in-place reversal
	for i := 0; i < len(items)/2; i++ {
		j := len(items) - 1 - i
		items[i], items[j] = items[j], items[i]
	}
	return items
}

// FindIndex returns the index of the first occurrence of a value in the slice
// items: The slice to search
// value: The value to locate in the slice
func FindIndex[T any](items []T, value T) int {
	for index, item := range items {
		if Equals(item, value) {
			return index
		}
	}
	return -1
}

// FindIndexWhere returns the index of the first element in the slice that satisfies the predicate
// items: The slice to search
// predicate: The predicate to test each element against
func FindIndexWhere[T any](items []T, predicate func(T) bool) int {
	for index, item := range items {
		if predicate(item) {
			return index
		}
	}

	return -1
}

// FindLast returns the last element in the slice that satisfies the predicate
// items: The slice to search
// predicate: The predicate to test each element against
func FindLast[T any](items []T, val T) T {
	// Instead of reversing the slice, consider iterating from the end
	for i := len(items) - 1; i >= 0; i-- {
		if Equals(items[i], val) {
			return items[i]
		}
	}
	var zero T
	return zero
}

// FindLastWhere returns the last element in the slice that satisfies the predicate
// items: The slice to search
// predicate: The predicate to test each element against
func FindLastWhere[T any](items []T, predicate func(T) bool) T {
	var item T

	for _, item = range Reverse(items) {
		if predicate(item) {
			return item
		}
	}

	return item
}

// FindLastIndex returns the index of the last occurrence of a value in the slice
// items: The slice to search
// value: The value to locate in the slice
func FindLastIndex[T any](items []T, value T) int {
	for i := len(items) - 1; i >= 0; i-- {
		if Equals(items[i], value) {
			return i
		}
	}
	return -1
}

// FindLastIndexWhere returns the index of the last element in the slice that satisfies the predicate
// items: The slice to search
// predicate: The predicate to test each element against
func FindLastIndexWhere[T any](items []T, predicate func(T) bool) int {
	for i := len(items) - 1; i >= 0; i-- {
		if predicate(items[i]) {
			return i
		}
	}
	return -1
}

// Contains determines whether an slice contains a specific value
// items: The slice to search
// value: The value to locate in the slice
func Contains[T comparable](items []T, value T) bool {
	for _, item := range items {
		if item == value {
			return true
		}
	}
	return false
}

// Distinct returns distinct elements from an slice
// items: The slice to search
func Distinct[T comparable](items []T) []T {
	seen := make(map[T]struct{}, len(items))
	results := make([]T, 0, len(items))

	for _, item := range items {
		if _, ok := seen[item]; !ok {
			seen[item] = struct{}{}
			results = append(results, item)
		}
	}

	return results
}

// DistinctBy returns distinct elements from an slice based on a key selector
// items: The slice to search
// keySelector: The function to extract the key from each element
func DistinctBy[T any, U comparable](items []T, keySelector func(T) U) []T {
	seen := make(map[U]struct{}, len(items))
	results := make([]T, 0, len(items))

	for _, item := range items {
		key := keySelector(item)
		if _, ok := seen[key]; !ok {
			seen[key] = struct{}{}
			results = append(results, item)
		}
	}

	return results
}

// Except returns the elements of an slice that do not appear in a second slice
// items: The slice to search
// other: The slice whose elements that also occur in the first slice will cause those elements to be removed from the returned slice
func Except[T comparable](items []T, other []T) []T {
	otherSet := make(map[T]struct{}, len(other))
	for _, item := range other {
		otherSet[item] = struct{}{}
	}

	var except []T
	for _, item := range items {
		if _, ok := otherSet[item]; !ok {
			except = append(except, item)
		}
	}

	return except
}

// Intersect returns the elements that appear in two slices
// items: The slice to search
// other: The slice whose distinct elements that also appear in the first slice will be returned
func Intersect[T comparable](items []T, other []T) []T {
	otherSet := make(map[T]struct{}, len(other))
	for _, item := range other {
		otherSet[item] = struct{}{}
	}

	var intersect []T
	seen := make(map[T]struct{})
	for _, item := range items {
		if _, ok := otherSet[item]; ok {
			if _, alreadySeen := seen[item]; !alreadySeen {
				intersect = append(intersect, item)
				seen[item] = struct{}{}
			}
		}
	}

	return intersect
}

// Union returns the elements that appear in either of two slices
// items: The first slice to search
// other: The second slice to search
func Union[T comparable](items []T, other []T) []T {
	return Distinct(append(items, other...))
}

// SequenceEqual determines whether two slices are equal
// items: The first slice to compare
// other: The second slice to compare
func SequenceEqual[T any](items []T, other []T) bool {
	if len(items) != len(other) {
		return false
	}

	for index, item := range items {
		if !Equals(item, other[index]) {
			return false
		}
	}

	return true
}

// Reduce applies an accumulator function over an slice
// items: The slice to reduce
// accumulator: The accumulator function to use
func Reduce[T any, U any](items []T, accumulator func(U, T) U, initialValue ...U) U {
	var zero U

	if len(initialValue) > 0 {
		zero = initialValue[0]
	}

	for _, item := range items {
		zero = accumulator(zero, item)
	}

	return zero
}

// Map projects each element of an slice into a new form
// items: The slice to map
// selector: The selector function to use
func Map[T any, U any](items []T, selector func(T) U) []U {
	var mapped []U

	for _, item := range items {
		mapped = append(mapped, selector(item))
	}

	return mapped
}

// Filter removes all elements from an slice that satisfy the predicate
// items: The slice to filter
// predicate: The predicate to test each element against
func Filter[T any](items []T, predicate func(T) bool) []T {
	var filtered []T

	for _, item := range items {
		if predicate(item) {
			filtered = append(filtered, item)
		}
	}

	return filtered
}

// ForEach performs the specified action on each element of an slice
// items: The slice to iterate
// action: The action to perform on each element
func ForEach[T any](items []T, action func(T)) {
	for _, item := range items {
		action(item)
	}
}

// RemoveAt removes an element from an slice at the specified index
// items: The slice to remove elements from
func RemoveAt[T any](items []T, index int) []T {
	return append(items[:index], items[index+1:]...)
}

// Remove removes the first occurrence of a specific object from an slice
// items: The slice to remove elements from
// value: The value to remove from the slice
func Remove[T any](items []T, value T) []T {
	index := FindIndex(items, value)
	if index != -1 {
		return RemoveAt(items, index)
	}
	return items
}

// RemoveWhere removes the first element in the slice that satisfies the predicate
// items: The slice to remove elements from
// predicate: The predicate to test each element against
func RemoveWhere[T any](items []T, predicate func(T) bool) []T {
	index := FindIndexWhere(items, predicate)
	if index != -1 {
		return RemoveAt(items, index)
	}
	return items
}

// KeyWhere returns the a Map of the slice where the key is the result of the selector function
// items: The slice to convert to a Map
// selector: The selector function to use
func KeyWhere[T any, U comparable](items []T, selector func(T) U) map[U]T {
	var key = make(map[U]T)

	for _, item := range items {
		key[selector(item)] = item
	}

	return key
}

// Key returns the a Map of the slice where the key is the value of the field at the specified path
// items: The slice to convert to a Map
// path: The path to the field to use as the key. If the field is not found, the item will not be added to the Map
func Key[T any, U comparable](items []T, path string) map[U]T {
	return KeyWhere(items, func(item T) U {
		value, _ := Get(item, path)
		casted, _ := value.(U)
		return casted
	})
}

// GroupWhere returns the a Map of the slice where the key is the result of the selector function and the value is an slice of all the elements that match the key
// items: The slice to convert to a Map
// selector: The selector function to use
func GroupWhere[T any, U comparable](items []T, selector func(T) U) map[U][]T {
	var key = make(map[U][]T)

	for _, item := range items {
		key[selector(item)] = append(key[selector(item)], item)
	}

	return key
}

// Group returns the a Map of the slice where the key is the value of the field at the specified path and the value is an slice of all the elements that match the key
// items: The slice to convert to a Map
// path: The path to the field to use as the key. If the field is not found, the item will not be added to the Map
func Group[T any, U comparable](items []T, path string) map[U][]T {
	return GroupWhere(items, func(item T) U {
		value, _ := Get(item, path)
		casted, _ := value.(U)
		return casted
	})
}

// Randomize returns a new slice with the elements in a random order
// items: The slice to randomize
func Randomize[T any](items []T) []T {
	rand.Shuffle(len(items), func(i, j int) {
		items[i], items[j] = items[j], items[i]
	})

	return items
}

// Slice returns a new slice with the elements from the specified start index to the specified end index
// items: The slice to slice
// start: The index to start at
// end: The index to end at
func Slice[T any](items []T, start int, end int) []T {
	return items[start:end]
}

// SortWhere sorts the elements of an slice in place using the specified comparer function
// items: The slice to sort
// comparer: The comparer function to use
func SortWhere[T any](items []T, comparer func(T, T) bool) []T {
	sort.Slice(items, func(i, j int) bool {
		return comparer(items[i], items[j])
	})
	return items
}

// Take returns a new slice with the specified number of elements from the start of the slice
// items: The slice to take elements from
// count: The number of elements to take
func Take[T any](items []T, count int) []T {
	if len(items) == 0 {
		return []T{}
	}
	if count > len(items) {
		count = len(items)
	}
	return items[:count]
}

// TakeLast returns a new slice with the specified number of elements from the end of the slice
// items: The slice to take elements from
// count: The number of elements to take
func TakeLast[T any](items []T, count int) []T {
	return items[len(items)-count:]
}

// TakeWhile returns a new slice with elements from the start of the slice while the predicate returns true
// items: The slice to take elements from
// predicate: The predicate to test each element against
func TakeWhile[T any](items []T, predicate func(T) bool) []T {
	var taken []T

	for _, item := range items {
		if !predicate(item) {
			break
		}

		taken = append(taken, item)
	}

	return taken
}

// Skip returns a new slice with the specified number of elements removed from the start of the slice
// items: The slice to skip elements from
// count: The number of elements to skip
func Skip[T any](items []T, count int) []T {
	if len(items) == 0 {
		return []T{}
	}
	if count > len(items) {
		count = len(items)
	}
	return items[count:]
}

// SkipLast returns a new slice with the specified number of elements removed from the end of the slice
// items: The slice to skip elements from
// count: The number of elements to skip
func SkipLast[T any](items []T, count int) []T {
	return items[:len(items)-count]
}

// SkipWhile returns a new slice with elements removed from the start of the slice while the predicate returns true
// items: The slice to skip elements from
// predicate: The predicate to test each element against
func SkipWhile[T any](items []T, predicate func(T) bool) []T {
	var skipped []T

	for _, item := range items {
		if !predicate(item) {
			skipped = append(skipped, item)
		}
	}

	return skipped
}

// Chunk returns a new slice with elements grouped into chunks of the specified size
// items: The slice to chunk
// size: The size of each chunk
func Chunk[T any](items []T, size int) [][]T {
	var chunks [][]T

	for i := 0; i < len(items); i += size {
		end := i + size

		if end > len(items) {
			end = len(items)
		}

		chunks = append(chunks, items[i:end])
	}

	return chunks
}

// Flatten returns a new slice with all sub-slice elements concatenated
// items: The slice to flatten
func Flatten[T any](items [][]T) []T {
	var flattened []T

	for _, item := range items {
		flattened = append(flattened, item...)
	}

	return flattened
}

// ReplaceAll replaces all occurrences of a value in an slice with another value
// items: The slice to replace values in
// oldValue: The value to replace
// newValue: The value to replace with
func ReplaceAll[T any](items []T, oldValue T, newValue T) []T {
	for index, item := range items {
		if Equals(item, oldValue) {
			items[index] = newValue
		}
	}

	return items
}

// ReplaceAllWhere replaces all occurrences of a value in an slice with the result of the selector function
// items: The slice to replace values in
// value: The value to replace with
// predicate: The selector function to use
func ReplaceAllWhere[T any](items []T, value T, predicate func(T) bool) []T {
	for index, item := range items {
		if predicate(item) {
			items[index] = value
		}
	}

	return items
}

// Replace replaces the first occurrence of a value in an slice with another value
// items: The slice to replace values in
// oldValue: The value to replace
// newValue: The value to replace with
func Replace[T any](items []T, oldValue T, newValue T) []T {
	index := FindIndex(items, oldValue)
	if index != -1 {
		items[index] = newValue
	}
	return items
}

// ReplaceWhere replaces the first occurrence of a value in an slice with the result of the selector function
// items: The slice to replace values in
// value: The value to replace with
// predicate: The selector function to use
func ReplaceWhere[T any](items []T, value T, predicate func(T) bool) []T {
	index := FindIndexWhere(items, predicate)
	if index != -1 {
		items[index] = value
	}
	return items
}

// Push adds an element to the end of an slice
// items: The slice to add the element to
// item: The element to add
func Push[T any](items []T, item T) []T {
	return append(items, item)
}

// Pop removes the last element from an slice and returns it
// items: The slice to remove the element from
func Pop[T any](items []T) (T, []T) {
	if len(items) == 0 {
		var zero T
		return zero, items
	}
	return items[len(items)-1], items[:len(items)-1]
}

// Unshift adds an element to the start of an slice
// items: The slice to add the element to
// item: The element to add
func Unshift[T any](items []T, item T) []T {
	return append([]T{item}, items...)
}

// Shift removes the first element from an slice and returns it
// items: The slice to remove the element from
func Shift[T any](items []T) (T, []T) {
	if len(items) == 0 {
		var zero T
		return zero, items
	}
	return items[0], items[1:]
}

// Last returns the last element in an slice
// items: The slice to get the last element from
func Last[T any](items []T) T {
	if len(items) == 0 {
		var zero T
		return zero
	}
	return items[len(items)-1]
}

// First returns the first element in an slice
// items: The slice to get the first element from
func First[T any](items []T) T {
	var item T
	if len(items) > 0 {
		item = items[0]
	}
	return item
}

// Sum returns the sum of all elements in an slice
// items: The slice to sum
func Sum[T int | int32 | int64 | float32 | float64](items []T) T {
	var sum T
	for _, item := range items {
		sum += item
	}
	return sum
}

// Average returns the average of all elements in an slice
// items: The slice to average
func Average[T int | int32 | int64 | float32 | float64](items []T) float64 {
	if len(items) == 0 {
		return 0
	}
	sum := Sum(items)
	return float64(sum) / float64(len(items))
}

// Min returns the minimum value in an slice
// items: The slice to get the minimum value from
func Min[T int | int32 | int64 | float32 | float64](items []T) T {
	if len(items) == 0 {
		var zero T
		return zero
	}
	min := items[0]
	for _, item := range items[1:] {
		if item < min {
			min = item
		}
	}
	return min
}

// Max returns the maximum value in an slice
// items: The slice to get the maximum value from
func Max[T int | int32 | int64 | float32 | float64](items []T) T {
	if len(items) == 0 {
		var zero T
		return zero
	}
	max := items[0]
	for _, item := range items[1:] {
		if item > max {
			max = item
		}
	}
	return max
}

// Partition splits an slice into two slices based on a predicate
// items: The slice to partition
// predicate: The predicate to test each element against
func Partition[T any](items []T, predicate func(T) bool) ([]T, []T) {
	var trueItems, falseItems []T
	for _, item := range items {
		if predicate(item) {
			trueItems = append(trueItems, item)
		} else {
			falseItems = append(falseItems, item)
		}
	}
	return trueItems, falseItems
}

// Zip combines two slices into a new slice using a selector function
// items: The slice to combine
// other: The slice to combine
// selector: The selector function to use
func Zip[T any, U any, V any](items []T, other []U, selector func(T, U) V) []V {
	minLen := len(items)
	if len(other) < minLen {
		minLen = len(other)
	}
	zipped := make([]V, minLen)
	for i := 0; i < minLen; i++ {
		zipped[i] = selector(items[i], other[i])
	}
	return zipped
}

// All determines whether all elements of an slice satisfy a condition
// items: The slice to search
// predicate: The predicate to test each element against
func All[T any](items []T, predicate func(T) bool) bool {
	for _, item := range items {
		if !predicate(item) {
			return false
		}
	}
	return true
}

// Any determines whether any element of an slice satisfies a condition
// items: The slice to search
// predicate: The predicate to test each element against
func Any[T any](items []T, predicate func(T) bool) bool {
	for _, item := range items {
		if predicate(item) {
			return true
		}
	}
	return false
}

// Count returns the number of elements in an slice that satisfy a condition
// items: The slice to search
// predicate: The predicate to test each element against
func Count[T any](items []T, predicate func(T) bool) int {
	count := 0
	for _, item := range items {
		if predicate(item) {
			count++
		}
	}
	return count
}

// At returns the value at the provided index or the default value if the index is out of bounds
// items: the slice
// index: the index
func At[T any](items []T, index int) T {
	var zero T
	if len(items) <= index || index < 0 || len(items) == 0 {
		return zero
	}

	return items[index]
}
