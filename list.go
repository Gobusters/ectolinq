package ectolinq

type List[T any] []T

// NewList returns a new List
func NewList[T any]() List[T] {
	return make(List[T], 0)
}

// From returns a List from an array
// items: The array to convert
func From[T any](items []T) List[T] {
	return items
}

// ForEach executes an action for each element in the array in parallel
// items: The array to iterate
// action: The action to perform on each element
func (l List[T]) ForEach(fn func(T)) List[T] {
	ForEach(l, fn)
	return l
}

// Find returns the first element in the array that satisfies the predicate
// items: The array to search
// predicate: The predicate to test each element against
func (l List[T]) Find(fn func(T) bool) T {
	return Find(l, fn)
}

// Reverse reverses the order of the elements in the array
// items: The array to reverse
func (l List[T]) Reverse(items []T) List[T] {
	return Reverse(l)
}

// FindIndex returns the index of the first occurrence of a value in the array
// items: The array to search
// value: The value to locate in the array
func (l List[T]) FindIndex(items []T, value T) int {
	return FindIndex(items, value)
}

// FindIndexWith returns the index of the first element in the array that satisfies the predicate
// items: The array to search
// predicate: The predicate to test each element against
func (l List[T]) FindIndexWith(items []T, predicate func(T) bool) int {
	return FindIndexWith(items, predicate)
}

// FindLast returns the last element in the array that satisfies the predicate
// items: The array to search
// predicate: The predicate to test each element against
func (l List[T]) FindLast(items []T, val T) T {
	return FindLast(items, val)
}

// FindLastWith returns the last element in the array that satisfies the predicate
// items: The array to search
// predicate: The predicate to test each element against
func (l List[T]) FindLastWith(items []T, predicate func(T) bool) T {
	return FindLastWith(items, predicate)
}

// FindLastIndex returns the index of the last occurrence of a value in the array
// items: The array to search
// value: The value to locate in the array
func (l List[T]) FindLastIndex(items []T, value T) int {
	return FindLastIndex(items, value)
}

// FindLastIndexWith returns the index of the last element in the array that satisfies the predicate
// items: The array to search
// predicate: The predicate to test each element against
func (l List[T]) FindLastIndexWith(items []T, predicate func(T) bool) int {
	return FindLastIndexWith(items, predicate)
}

// IndexOf returns the index of the first occurrence of a value in the array
// items: The array to search
// value: The value to locate in the array
func (l List[T]) IndexOf(items []T, value T) int {
	return IndexOf(items, value)
}

// LastIndexOf returns the index of the last occurrence of a value in the array
// items: The array to search
// value: The value to locate in the array
func (l List[T]) LastIndexOf(items []T, value T) int {
	return LastIndexOf(items, value)
}

// Contains determines whether an array contains a specific value
// items: The array to search
// value: The value to locate in the array
func (l List[T]) Contains(items []T, value T) bool {
	return Contains(items, value)
}

// All determines whether all elements of an array satisfy a condition
// items: The array to search
// predicate: The predicate to test each element against
func (l List[T]) All(items []T, predicate func(T) bool) bool {
	return All(items, predicate)
}

// Any determines whether any element of an array satisfies a condition
// items: The array to search
// predicate: The predicate to test each element against
func (l List[T]) Any(items []T, predicate func(T) bool) bool {
	return Any(items, predicate)
}

// Count returns the number of elements in an array that satisfy a condition
// items: The array to search
// predicate: The predicate to test each element against
func (l List[T]) Count(items []T, predicate func(T) bool) int {
	return Count(items, predicate)
}

// Distinct returns distinct elements from an array
// items: The array to search
func (l List[T]) Distinct(items []T) List[T] {
	return Distinct(items)
}

// Except returns the elements of an array that do not appear in a second array
// items: The array to search
// other: The array whose elements that also occur in the first array will cause those elements to be removed from the returned array
func (l List[T]) Except(items []T, other []T) List[T] {
	return Except(items, other)
}

// Intersect returns the elements that appear in two arrays
// items: The array to search
// other: The array whose distinct elements that also appear in the first array will be returned
func (l List[T]) Intersect(items []T, other []T) List[T] {
	return Intersect(items, other)
}

// Union returns the elements that appear in either of two arrays
// items: The first array to search
// other: The second array to search
func (l List[T]) Union(items []T, other []T) List[T] {
	return Union(items, other)
}

// SequenceEqual determines whether two arrays are equal
// items: The first array to compare
// other: The second array to compare
func (l List[T]) SequenceEqual(items []T, other []T) bool {
	return SequenceEqual(items, other)
}

// Reduce applies an accumulator function over an array
// items: The array to reduce
// accumulator: The accumulator function to use
func (l List[T]) Reduce(items []T, accumulator func(T, T) T) T {
	return Reduce(items, accumulator)
}

// ReduceWith applies an accumulator function over an array. Starts with the specified value
// items: The array to reduce
// initialValue: The value to start with
// accumulator: The accumulator function to use
func (l List[T]) ReduceWith(items []T, initialValue T, accumulator func(T, T) T) T {
	return ReduceWith(items, initialValue, accumulator)
}

// Filter removes all elements from an array that satisfy the predicate
// items: The array to filter
// predicate: The predicate to test each element against
func (l List[T]) Filter(items []T, predicate func(T) bool) List[T] {
	return Filter(items, predicate)
}

// RemoveAt removes an element from an array at the specified index
// items: The array to remove elements from
func (l List[T]) RemoveAt(items []T, index int) List[T] {
	return RemoveAt(items, index)
}

// Remove removes the first occurrence of a specific object from an array
// items: The array to remove elements from
// value: The value to remove from the array
func (l List[T]) Remove(items []T, value T) List[T] {
	return Remove(items, value)
}

// RemoveWith removes the first element in the array that satisfies the predicate
// items: The array to remove elements from
// predicate: The predicate to test each element against
func (l List[T]) RemoveWith(items []T, predicate func(T) bool) List[T] {
	return RemoveWith(items, predicate)
}

// Randomize returns a new array with the elements in a random order
// items: The array to randomize
func (l List[T]) Randomize(items []T) List[T] {
	return Randomize(items)
}

// Slice returns a new array with the elements from the specified start index to the specified end index
// items: The array to slice
// start: The index to start at
// end: The index to end at
func (l List[T]) Slice(items []T, start int, end int) List[T] {
	return Slice(items, start, end)
}

// SortWith sorts the elements of an array in place using the specified comparer function
// items: The array to sort
// comparer: The comparer function to use
func (l List[T]) SortWith(items []T, comparer func(T, T) bool) List[T] {
	return SortWith(items, comparer)
}

// Take returns a new array with the specified number of elements from the start of the array
// items: The array to take elements from
// count: The number of elements to take
func (l List[T]) Take(items []T, count int) List[T] {
	return Take(items, count)
}

// TakeLast returns a new array with the specified number of elements from the end of the array
// items: The array to take elements from
// count: The number of elements to take
func (l List[T]) TakeLast(items []T, count int) List[T] {
	return TakeLast(items, count)
}

// TakeWhile returns a new array with elements from the start of the array while the predicate returns true
// items: The array to take elements from
// predicate: The predicate to test each element against
func (l List[T]) TakeWhile(items []T, predicate func(T) bool) List[T] {
	return TakeWhile(items, predicate)
}

// Skip returns a new array with the specified number of elements removed from the start of the array
// items: The array to skip elements from
// count: The number of elements to skip
func (l List[T]) Skip(items []T, count int) List[T] {
	return Skip(items, count)
}

// SkipLast returns a new array with the specified number of elements removed from the end of the array
// items: The array to skip elements from
// count: The number of elements to skip
func (l List[T]) SkipLast(items []T, count int) List[T] {
	return SkipLast(items, count)
}

// SkipWhile returns a new array with elements removed from the start of the array while the predicate returns true
// items: The array to skip elements from
// predicate: The predicate to test each element against
func (l List[T]) SkipWhile(items []T, predicate func(T) bool) List[T] {
	return SkipWhile(items, predicate)
}

// Chunk returns a new array with elements grouped into chunks of the specified size
// items: The array to chunk
// size: The size of each chunk
func (l List[T]) Chunk(items []T, size int) []List[T] {
	var genericLists []List[T]
	chunks := Chunk(items, size)
	for _, slice := range chunks {
		genericLists = append(genericLists, List[T](slice))
	}
	return genericLists
}

// Flatten returns a new array with all sub-array elements concatenated
// items: The array to flatten
func (l List[T]) Flatten(items [][]T) List[T] {
	return Flatten(items)
}

// ReplaceAll replaces all occurrences of a value in an array with another value
// items: The array to replace values in
// oldValue: The value to replace
// newValue: The value to replace with
func (l List[T]) ReplaceAll(items []T, oldValue T, newValue T) List[T] {
	return ReplaceAll(items, oldValue, newValue)
}

// ReplaceAllWith replaces all occurrences of a value in an array with the result of the selector function
// items: The array to replace values in
// value: The value to replace with
// predicate: The selector function to use
func (l List[T]) ReplaceAllWith(items []T, value T, predicate func(T) bool) List[T] {
	return ReplaceAllWith(items, value, predicate)
}

// Replace replaces the first occurrence of a value in an array with another value
// items: The array to replace values in
// oldValue: The value to replace
// newValue: The value to replace with
func (l List[T]) Replace(items []T, oldValue T, newValue T) List[T] {
	return Replace(items, oldValue, newValue)
}

// ReplaceWith replaces the first occurrence of a value in an array with the result of the selector function
// items: The array to replace values in
// value: The value to replace with
// predicate: The selector function to use
func (l List[T]) ReplaceWith(items []T, value T, predicate func(T) bool) List[T] {
	return ReplaceWith(items, value, predicate)
}
