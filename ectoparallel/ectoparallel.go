package ectoparallel

import "sync"

// ForEach executes an action for each element in the array in parallel
// items: The array to iterate
// action: The action to perform on each element
func ForEach[T any](slice []T, fn func(T)) {
	var wg sync.WaitGroup

	for _, item := range slice {
		wg.Add(1)
		go func(it T) {
			defer wg.Done()
			fn(it)
		}(item)
	}

	wg.Wait() // Wait for all goroutines to finish
}

// Map projects each element of an array into a new form in parallel
// items: The array to map
// selector: The selector function to use
func Map[T any, U any](slice []T, fn func(T) U) []U {
	var wg sync.WaitGroup
	var mapped []U

	for _, item := range slice {
		wg.Add(1)
		go func(it T) {
			defer wg.Done()
			mapped = append(mapped, fn(it))
		}(item)
	}

	wg.Wait() // Wait for all goroutines to finish

	return mapped
}

// Filter removes all elements from an array that satisfy the predicate in parallel
// items: The array to filter
// predicate: The predicate to test each element against
func Filter[T any](slice []T, fn func(T) bool) []T {
	var wg sync.WaitGroup
	var filtered []T

	for _, item := range slice {
		wg.Add(1)
		go func(it T) {
			defer wg.Done()
			if fn(it) {
				filtered = append(filtered, it)
			}
		}(item)
	}

	wg.Wait() // Wait for all goroutines to finish

	return filtered
}
