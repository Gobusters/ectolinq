package ectolinq

type List[T any] []T

// NewList returns a new List
func NewList[T any]() List[T] {
	return make(List[T], 0)
}

// Length returns the length of the array
func (l List[T]) Length() int {
	return len(l)
}

// ToList returns a List from an array
func ToList[T any](items []T) List[T] {
	return items
}

// ForEach executes an action for each element in the array in parallel
// action: The action to perform on each element
func (l List[T]) ForEach(fn func(T)) List[T] {
	ForEach(l, fn)
	return l
}

// Find returns the first element in the array that satisfies the predicate
// predicate: The predicate to test each element against
func (l List[T]) Find(fn func(T) bool) T {
	return Find(l, fn)
}

// Reverse reverses the order of the elements in the array
func (l List[T]) Reverse() List[T] {
	return Reverse(l)
}

// FindIndex returns the index of the first occurrence of a value in the array
// value: The value to locate in the array
func (l List[T]) FindIndex(value T) int {
	return FindIndex(l, value)
}

// FindIndexWhere returns the index of the first element in the array that satisfies the predicate
// predicate: The predicate to test each element against
func (l List[T]) FindIndexWhere(predicate func(T) bool) int {
	return FindIndexWhere(l, predicate)
}

// FindLast returns the last element in the array that satisfies the predicate
// predicate: The predicate to test each element against
func (l List[T]) FindLast(val T) T {
	return FindLast(l, val)
}

// FindLastWhere returns the last element in the array that satisfies the predicate
// predicate: The predicate to test each element against
func (l List[T]) FindLastWhere(predicate func(T) bool) T {
	return FindLastWhere(l, predicate)
}

// FindLastIndex returns the index of the last occurrence of a value in the array
// value: The value to locate in the array
func (l List[T]) FindLastIndex(value T) int {
	return FindLastIndex(l, value)
}

// FindLastIndexWhere returns the index of the last element in the array that satisfies the predicate
// predicate: The predicate to test each element against
func (l List[T]) FindLastIndexWhere(predicate func(T) bool) int {
	return FindLastIndexWhere(l, predicate)
}

// Contains determines whether an array contains a specific value
// value: The value to locate in the array
func (l List[T]) Contains(value T) bool {
	return Any(l, func(item T) bool {
		return Equals(item, value)
	})
}

// All determines whether all elements of an array satisfy a condition
// predicate: The predicate to test each element against
func (l List[T]) All(predicate func(T) bool) bool {
	return All(l, predicate)
}

// Any determines whether any element of an array satisfies a condition
// predicate: The predicate to test each element against
func (l List[T]) Any(predicate func(T) bool) bool {
	return Any(l, predicate)
}

// Count returns the number of elements in an array that satisfy a condition
// predicate: The predicate to test each element against
func (l List[T]) Count(predicate func(T) bool) int {
	return Count(l, predicate)
}

// Distinct returns distinct elements from an array
func (l List[T]) Distinct() List[T] {
	var result List[T]
	for _, item := range l {
		if !result.Contains(item) {
			result = append(result, item)
		}
	}
	return result
}

// Except returns the elements of an array that do not appear in a second array
// other: The array whose elements that also occur in the first array will cause those elements to be removed from the returned array
func (l List[T]) Except(other []T) List[T] {
	otherList := ToList(other)
	return Filter(l, func(item T) bool {
		return !otherList.Contains(item)
	})
}

// Intersect returns the elements that appear in two arrays
// other: The array whose distinct elements that also appear in the first array will be returned
func (l List[T]) Intersect(other []T) List[T] {
	otherList := ToList(other)
	return Filter(l, func(item T) bool {
		return otherList.Contains(item)
	})
}

// Union returns the elements that appear in either of two arrays
// other: The second array to search
func (l List[T]) Union(other []T) List[T] {
	otherList := ToList(other)
	return l.Concat(otherList).Distinct()
}

// Concat concatenates two lists
func (l List[T]) Concat(other List[T]) List[T] {
	return append(l, other...)
}

// SequenceEqual determines whether two arrays are equal
// other: The second array to compare
func (l List[T]) SequenceEqual(other []T) bool {
	return SequenceEqual(l, other)
}

// Reduce applies an accumulator function over an array
// accumulator: The accumulator function to use
func (l List[T]) Reduce(accumulator func(T, T) T) T {
	return Reduce(l, accumulator)
}

// ReduceWhere applies an accumulator function over an array. Starts with the specified value
// initialValue: The value to start with
// accumulator: The accumulator function to use
func (l List[T]) ReduceWhere(initialValue T, accumulator func(T, T) T) T {
	return ReduceWhere(l, initialValue, accumulator)
}

// Filter removes all elements from an array that satisfy the predicate
// predicate: The predicate to test each element against
func (l List[T]) Filter(predicate func(T) bool) List[T] {
	return Filter(l, predicate)
}

// RemoveAt removes an element from an array at the specified index
// index: The index to remove at
func (l List[T]) RemoveAt(index int) List[T] {
	return RemoveAt(l, index)
}

// Remove removes the first occurrence of a specific object from an array
// value: The value to remove from the array
func (l List[T]) Remove(value T) List[T] {
	return Remove(l, value)
}

// RemoveWhere removes the first element in the array that satisfies the predicate
// predicate: The predicate to test each element against
func (l List[T]) RemoveWhere(predicate func(T) bool) List[T] {
	return RemoveWhere(l, predicate)
}

// Randomize returns a new array with the elements in a random order
func (l List[T]) Randomize() List[T] {
	return Randomize(l)
}

// Slice returns a new array with the elements from the specified start index to the specified end index
// start: The index to start at
// end: The index to end at
func (l List[T]) Slice(start int, end int) List[T] {
	return Slice(l, start, end)
}

// SortWhere sorts the elements of an array in place using the specified comparer function
// comparer: The comparer function to use
func (l List[T]) SortWhere(comparer func(T, T) bool) List[T] {
	return SortWhere(l, comparer)
}

// Take returns a new array with the specified number of elements from the start of the array
// count: The number of elements to take
func (l List[T]) Take(count int) List[T] {
	return Take(l, count)
}

// TakeLast returns a new array with the specified number of elements from the end of the array
// count: The number of elements to take
func (l List[T]) TakeLast(count int) List[T] {
	return TakeLast(l, count)
}

// TakeWhile returns a new array with elements from the start of the array while the predicate returns true
// predicate: The predicate to test each element against
func (l List[T]) TakeWhile(predicate func(T) bool) List[T] {
	return TakeWhile(l, predicate)
}

// Skip returns a new array with the specified number of elements removed from the start of the array
// count: The number of elements to skip
func (l List[T]) Skip(count int) List[T] {
	return Skip(l, count)
}

// SkipLast returns a new array with the specified number of elements removed from the end of the array
// count: The number of elements to skip
func (l List[T]) SkipLast(count int) List[T] {
	return SkipLast(l, count)
}

// SkipWhile returns a new array with elements removed from the start of the array while the predicate returns true
// predicate: The predicate to test each element against
func (l List[T]) SkipWhile(predicate func(T) bool) List[T] {
	return SkipWhile(l, predicate)
}

// Chunk returns a new array with elements grouped into chunks of the specified size
// size: The size of each chunk
func (l List[T]) Chunk(size int) []List[T] {
	var genericLists []List[T]
	chunks := Chunk(l, size)
	for _, slice := range chunks {
		genericLists = append(genericLists, List[T](slice))
	}
	return genericLists
}

// ReplaceAll replaces all occurrences of a value in an array with another value
// oldValue: The value to replace
// newValue: The value to replace with
func (l List[T]) ReplaceAll(oldValue T, newValue T) List[T] {
	return ReplaceAll(l, oldValue, newValue)
}

// ReplaceAllWhere replaces all occurrences of a value in an array with the result of the selector function
// value: The value to replace with
// predicate: The selector function to use
func (l List[T]) ReplaceAllWhere(value T, predicate func(T) bool) List[T] {
	return ReplaceAllWhere(l, value, predicate)
}

// Replace replaces the first occurrence of a value in an array with another value
// oldValue: The value to replace
// newValue: The value to replace with
func (l List[T]) Replace(oldValue T, newValue T) List[T] {
	return Replace(l, oldValue, newValue)
}

// ReplaceWhere replaces the first occurrence of a value in an array with the result of the selector function
// value: The value to replace with
// predicate: The selector function to use
func (l List[T]) ReplaceWhere(value T, predicate func(T) bool) List[T] {
	return ReplaceWhere(l, value, predicate)
}

// Push adds an element to the end of the array
// value: The value to push
func (l List[T]) Push(value T) List[T] {
	l = Push(l, value)
	return l
}

// Pop removes the last element from an array and returns it
func (l List[T]) Pop() (T, List[T]) {
	return Pop(l)
}

// Unshift adds an element to the start of the array
// value: The value to unshift
func (l List[T]) Unshift(value T) List[T] {
	return Unshift(l, value)
}

// Shift removes the first element from an array and returns it
func (l List[T]) Shift() (T, List[T]) {
	return Shift(l)
}

// Last returns the last element in the array
func (l List[T]) Last() T {
	return Last(l)
}

// First returns the first element in the array
func (l List[T]) First() T {
	return First(l)
}

// Partition splits an array into two arrays based on a predicate
func (l List[T]) Partition(predicate func(T) bool) (List[T], List[T]) {
	return Partition(l, predicate)
}

// Zip combines two arrays into a new array using a selector function
func (l List[T]) Zip(other []T, selector func(T, T) T) List[T] {
	return Zip(l, other, selector)
}
