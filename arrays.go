package ectolinq

import (
	"math/rand"
)

// Find returns the first element in the array that satisfies the predicate
// items: The array to search
// predicate: The predicate to test each element against
func Find[T any](items []T, predicate func(T) bool) T {
	var item T

	for _, item = range items {
		if predicate(item) {
			return item
		}
	}

	return item
}

// Reverse reverses the order of the elements in the array
// items: The array to reverse
func Reverse[T any](items []T) []T {
	var reversed []T

	for index := len(items) - 1; index >= 0; index-- {
		reversed = append(reversed, items[index])
	}

	return reversed
}

// FindIndex returns the index of the first occurrence of a value in the array
// items: The array to search
// value: The value to locate in the array
func FindIndex[T any](items []T, value T) int {
	for index, item := range items {
		if Equals(item, value) {
			return index
		}
	}

	return -1
}

// FindIndexWhere returns the index of the first element in the array that satisfies the predicate
// items: The array to search
// predicate: The predicate to test each element against
func FindIndexWhere[T any](items []T, predicate func(T) bool) int {
	for index, item := range items {
		if predicate(item) {
			return index
		}
	}

	return -1
}

// FindLast returns the last element in the array that satisfies the predicate
// items: The array to search
// predicate: The predicate to test each element against
func FindLast[T any](items []T, val T) T {
	return FindLastWhere(items, func(item T) bool {
		return Equals(item, val)
	})
}

// FindLastWhere returns the last element in the array that satisfies the predicate
// items: The array to search
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

// FindLastIndex returns the index of the last occurrence of a value in the array
// items: The array to search
// value: The value to locate in the array
func FindLastIndex[T any](items []T, value T) int {
	return FindLastIndexWhere(items, func(item T) bool {
		return Equals(item, value)
	})
}

// FindLastIndexWhere returns the index of the last element in the array that satisfies the predicate
// items: The array to search
// predicate: The predicate to test each element against
func FindLastIndexWhere[T any](items []T, predicate func(T) bool) int {
	for index, item := range Reverse(items) {
		if predicate(item) {
			return index
		}
	}

	return -1
}

// IndexOf returns the index of the first occurrence of a value in the array
// items: The array to search
// value: The value to locate in the array
func IndexOf[T any](items []T, value T) int {
	for index, item := range items {
		if Equals(item, value) {
			return index
		}
	}

	return -1
}

// LastIndexOf returns the index of the last occurrence of a value in the array
// items: The array to search
// value: The value to locate in the array
func LastIndexOf[T any](items []T, value T) int {
	for index, item := range Reverse(items) {
		if Equals(item, value) {
			return index
		}
	}

	return -1
}

// Contains determines whether an array contains a specific value
// items: The array to search
// value: The value to locate in the array
func Contains[T any](items []T, value T) bool {
	return IndexOf(items, value) != -1
}

// All determines whether all elements of an array satisfy a condition
// items: The array to search
// predicate: The predicate to test each element against
func All[T any](items []T, predicate func(T) bool) bool {
	for _, item := range items {
		if !predicate(item) {
			return false
		}
	}

	return true
}

// Any determines whether any element of an array satisfies a condition
// items: The array to search
// predicate: The predicate to test each element against
func Any[T any](items []T, predicate func(T) bool) bool {
	for _, item := range items {
		if predicate(item) {
			return true
		}
	}

	return false
}

// Count returns the number of elements in an array that satisfy a condition
// items: The array to search
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

// Distinct returns distinct elements from an array
// items: The array to search
func Distinct[T any](items []T) []T {
	var distinct []T

	for _, item := range items {
		if !Contains(distinct, item) {
			distinct = append(distinct, item)
		}
	}

	return distinct
}

// Except returns the elements of an array that do not appear in a second array
// items: The array to search
// other: The array whose elements that also occur in the first array will cause those elements to be removed from the returned array
func Except[T any](items []T, other []T) []T {
	var except []T

	for _, item := range items {
		if !Contains(other, item) {
			except = append(except, item)
		}
	}

	return except
}

// Intersect returns the elements that appear in two arrays
// items: The array to search
// other: The array whose distinct elements that also appear in the first array will be returned
func Intersect[T any](items []T, other []T) []T {
	var intersect []T

	for _, item := range items {
		if Contains(other, item) {
			intersect = append(intersect, item)
		}
	}

	return intersect
}

// Union returns the elements that appear in either of two arrays
// items: The first array to search
// other: The second array to search
func Union[T any](items []T, other []T) []T {
	return Distinct(append(items, other...))
}

// SequenceEqual determines whether two arrays are equal
// items: The first array to compare
// other: The second array to compare
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

// Reduce applies an accumulator function over an array. Starts with the default value of the type
// items: The array to reduce
// accumulator: The accumulator function to use
func Reduce[T any](items []T, accumulator func(T, T) T) T {
	var reduced T

	for _, item := range items {
		reduced = accumulator(reduced, item)
	}

	return reduced
}

// ReduceWhere applies an accumulator function over an array. Starts with the specified value
// items: The array to reduce
// initialValue: The value to start with
// accumulator: The accumulator function to use
func ReduceWhere[T any](items []T, initialValue T, accumulator func(T, T) T) T {
	var reduced T

	for _, item := range items {
		reduced = accumulator(reduced, item)
	}

	return reduced
}

// Map projects each element of an array into a new form
// items: The array to map
// selector: The selector function to use
func Map[T any, U any](items []T, selector func(T) U) []U {
	var mapped []U

	for _, item := range items {
		mapped = append(mapped, selector(item))
	}

	return mapped
}

// Filter removes all elements from an array that satisfy the predicate
// items: The array to filter
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

// ForEach performs the specified action on each element of an array
// items: The array to iterate
// action: The action to perform on each element
func ForEach[T any](items []T, action func(T)) {
	for _, item := range items {
		action(item)
	}
}

// RemoveAt removes an element from an array at the specified index
// items: The array to remove elements from
func RemoveAt[T any](items []T, index int) []T {
	return append(items[:index], items[index+1:]...)
}

// Remove removes the first occurrence of a specific object from an array
// items: The array to remove elements from
// value: The value to remove from the array
func Remove[T any](items []T, value T) []T {
	return RemoveAt(items, IndexOf(items, value))
}

// RemoveWhere removes the first element in the array that satisfies the predicate
// items: The array to remove elements from
// predicate: The predicate to test each element against
func RemoveWhere[T any](items []T, predicate func(T) bool) []T {
	return RemoveAt(items, FindIndexWhere(items, predicate))
}

// KeyWhere returns the a Map of the array where the key is the result of the selector function
// items: The array to convert to a Map
// selector: The selector function to use
func KeyWhere[T any, U comparable](items []T, selector func(T) U) map[U]T {
	var key = make(map[U]T)

	for _, item := range items {
		key[selector(item)] = item
	}

	return key
}

// Key returns the a Map of the array where the key is the value of the field at the specified path
// items: The array to convert to a Map
// path: The path to the field to use as the key. If the field is not found, the item will not be added to the Map
func Key[T any, U comparable](items []T, path string) map[U]T {
	return KeyWhere(items, func(item T) U {
		value, _ := Get(item, path)
		casted, _ := value.(U)
		return casted
	})
}

// GroupWhere returns the a Map of the array where the key is the result of the selector function and the value is an array of all the elements that match the key
// items: The array to convert to a Map
// selector: The selector function to use
func GroupWhere[T any, U comparable](items []T, selector func(T) U) map[U][]T {
	var key = make(map[U][]T)

	for _, item := range items {
		key[selector(item)] = append(key[selector(item)], item)
	}

	return key
}

// Group returns the a Map of the array where the key is the value of the field at the specified path and the value is an array of all the elements that match the key
// items: The array to convert to a Map
// path: The path to the field to use as the key. If the field is not found, the item will not be added to the Map
func Group[T any, U comparable](items []T, path string) map[U][]T {
	return GroupWhere(items, func(item T) U {
		value, _ := Get(item, path)
		casted, _ := value.(U)
		return casted
	})
}

// Randomize returns a new array with the elements in a random order
// items: The array to randomize
func Randomize[T any](items []T) []T {
	rand.Shuffle(len(items), func(i, j int) {
		items[i], items[j] = items[j], items[i]
	})

	return items
}

// Slice returns a new array with the elements from the specified start index to the specified end index
// items: The array to slice
// start: The index to start at
// end: The index to end at
func Slice[T any](items []T, start int, end int) []T {
	return items[start:end]
}

// SortWhere sorts the elements of an array in place using the specified comparer function
// items: The array to sort
// comparer: The comparer function to use
func SortWhere[T any](items []T, comparer func(T, T) bool) []T {
	for i := 0; i < len(items)-1; i++ {
		for j := i + 1; j < len(items); j++ {
			if comparer(items[i], items[j]) {
				items[i], items[j] = items[j], items[i]
			}
		}
	}

	return items
}

// Take returns a new array with the specified number of elements from the start of the array
// items: The array to take elements from
// count: The number of elements to take
func Take[T any](items []T, count int) []T {
	return items[:count]
}

// TakeLast returns a new array with the specified number of elements from the end of the array
// items: The array to take elements from
// count: The number of elements to take
func TakeLast[T any](items []T, count int) []T {
	return items[len(items)-count:]
}

// TakeWhile returns a new array with elements from the start of the array while the predicate returns true
// items: The array to take elements from
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

// Skip returns a new array with the specified number of elements removed from the start of the array
// items: The array to skip elements from
// count: The number of elements to skip
func Skip[T any](items []T, count int) []T {
	return items[count:]
}

// SkipLast returns a new array with the specified number of elements removed from the end of the array
// items: The array to skip elements from
// count: The number of elements to skip
func SkipLast[T any](items []T, count int) []T {
	return items[:len(items)-count]
}

// SkipWhile returns a new array with elements removed from the start of the array while the predicate returns true
// items: The array to skip elements from
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

// Chunk returns a new array with elements grouped into chunks of the specified size
// items: The array to chunk
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

// Flatten returns a new array with all sub-array elements concatenated
// items: The array to flatten
func Flatten[T any](items [][]T) []T {
	var flattened []T

	for _, item := range items {
		flattened = append(flattened, item...)
	}

	return flattened
}

// ReplaceAll replaces all occurrences of a value in an array with another value
// items: The array to replace values in
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

// ReplaceAllWhere replaces all occurrences of a value in an array with the result of the selector function
// items: The array to replace values in
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

// Replace replaces the first occurrence of a value in an array with another value
// items: The array to replace values in
// oldValue: The value to replace
// newValue: The value to replace with
func Replace[T any](items []T, oldValue T, newValue T) []T {
	items[IndexOf(items, oldValue)] = newValue
	return items
}

// ReplaceWhere replaces the first occurrence of a value in an array with the result of the selector function
// items: The array to replace values in
// value: The value to replace with
// predicate: The selector function to use
func ReplaceWhere[T any](items []T, value T, predicate func(T) bool) []T {
	items[FindIndexWhere(items, predicate)] = value
	return items
}

// Push adds an element to the end of an array
// items: The array to add the element to
// item: The element to add
func Push[T any](items []T, item T) []T {
	return append(items, item)
}

// Pop removes the last element from an array and returns it
// items: The array to remove the element from
func Pop[T any](items []T) (T, []T) {
	return items[len(items)-1], items[:len(items)-1]
}

// Unshift adds an element to the start of an array
// items: The array to add the element to
// item: The element to add
func Unshift[T any](items []T, item T) []T {
	return append([]T{item}, items...)
}

// Shift removes the first element from an array and returns it
// items: The array to remove the element from
func Shift[T any](items []T) (T, []T) {
	return items[0], items[1:]
}

// Last returns the last element in an array
// items: The array to get the last element from
func Last[T any](items []T) T {
	return items[len(items)-1]
}

// First returns the first element in an array
// items: The array to get the first element from
func First[T any](items []T) T {
	var item T
	if len(items) > 0 {
		item = items[0]
	}
	return item
}
