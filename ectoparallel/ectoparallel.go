package ectoparallel

import (
	"runtime"
	"sync"
)

// ForEach executes an action for each element in the array in parallel
// items: The array to iterate
// action: The action to perform on each element
func ForEach[T any](slice []T, fn func(T)) {
	var wg sync.WaitGroup
	numWorkers := runtime.GOMAXPROCS(0)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(start int) {
			defer wg.Done()
			for j := start; j < len(slice); j += numWorkers {
				fn(slice[j])
			}
		}(i)
	}

	wg.Wait()
}

// Map projects each element of an array into a new form in parallel
// items: The array to map
// selector: The selector function to use
func Map[T any, U any](slice []T, fn func(T) U) []U {
	mapped := make([]U, len(slice))
	var wg sync.WaitGroup
	numWorkers := runtime.GOMAXPROCS(0)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(start int) {
			defer wg.Done()
			for j := start; j < len(slice); j += numWorkers {
				mapped[j] = fn(slice[j])
			}
		}(i)
	}

	wg.Wait()
	return mapped
}

// Filter removes all elements from an array that satisfy the predicate in parallel
// items: The array to filter
// predicate: The predicate to test each element against
func Filter[T any](slice []T, fn func(T) bool) []T {
	filtered := make([]T, 0, len(slice))
	var mu sync.Mutex
	var wg sync.WaitGroup
	numWorkers := runtime.GOMAXPROCS(0)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(start int) {
			defer wg.Done()
			localFiltered := make([]T, 0)
			for j := start; j < len(slice); j += numWorkers {
				if fn(slice[j]) {
					localFiltered = append(localFiltered, slice[j])
				}
			}
			mu.Lock()
			filtered = append(filtered, localFiltered...)
			mu.Unlock()
		}(i)
	}

	wg.Wait()
	return filtered
}
