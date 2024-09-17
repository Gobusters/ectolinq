package ectoparallel

import (
	"sort"
	"sync"
	"testing"
)

func TestForEach(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	sum := 0
	var mu sync.Mutex

	ForEach(numbers, func(n int) {
		mu.Lock()
		defer mu.Unlock()
		sum += n
	})

	if sum != 15 {
		t.Errorf("Expected sum to be 15, got %d", sum)
	}
}

func TestMap(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	squared := Map(numbers, func(n int) int {
		return n * n
	})

	expected := []int{1, 4, 9, 16, 25}
	if !equalSlices(squared, expected) {
		t.Errorf("Expected %v, got %v", expected, squared)
	}
}

func TestFilter(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evens := Filter(numbers, func(n int) bool {
		return n%2 == 0
	})

	expected := []int{2, 4, 6, 8, 10}
	sort.Ints(evens) // Sort because Filter doesn't guarantee order
	if !equalSlices(evens, expected) {
		t.Errorf("Expected %v, got %v", expected, evens)
	}
}

func TestParallelExecution(t *testing.T) {
	numbers := make([]int, 1000)
	for i := range numbers {
		numbers[i] = i
	}

	var counter int64
	var mu sync.Mutex

	ForEach(numbers, func(n int) {
		mu.Lock()
		counter++
		mu.Unlock()
	})

	if int(counter) != len(numbers) {
		t.Errorf("Expected counter to be %d, got %d", len(numbers), counter)
	}
}

// Helper function to compare slices
func equalSlices[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
