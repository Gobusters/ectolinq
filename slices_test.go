package ectolinq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	t.Run("Find existing element", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5}
		result := Find(items, func(i int) bool { return i > 3 })
		assert.Equal(t, 4, result, "Should find first element greater than 3")
	})

	t.Run("Find non-existing element", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5}
		result := Find(items, func(i int) bool { return i > 10 })
		assert.Equal(t, 0, result, "Should return zero value when element not found")
	})

	t.Run("Find in empty slice", func(t *testing.T) {
		var items []int
		result := Find(items, func(i int) bool { return i > 0 })
		assert.Equal(t, 0, result, "Should return zero value for empty slice")
	})
}

func TestReverse(t *testing.T) {
	t.Run("Reverse non-empty slice", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5}
		result := Reverse(items)
		assert.Equal(t, []int{5, 4, 3, 2, 1}, result, "Should reverse the slice")
	})

	t.Run("Reverse empty slice", func(t *testing.T) {
		var items []int
		result := Reverse(items)
		assert.Empty(t, result, "Should return empty slice for empty input")
	})

	t.Run("Reverse single element slice", func(t *testing.T) {
		items := []int{1}
		result := Reverse(items)
		assert.Equal(t, []int{1}, result, "Should return same slice for single element")
	})
}

func TestFindIndex(t *testing.T) {
	t.Run("Find existing element", func(t *testing.T) {
		items := []string{"apple", "banana", "cherry"}
		result := FindIndex(items, "banana")
		assert.Equal(t, 1, result, "Should return correct index for existing element")
	})

	t.Run("Find non-existing element", func(t *testing.T) {
		items := []string{"apple", "banana", "cherry"}
		result := FindIndex(items, "date")
		assert.Equal(t, -1, result, "Should return -1 for non-existing element")
	})

	t.Run("Find in empty slice", func(t *testing.T) {
		var items []string
		result := FindIndex(items, "apple")
		assert.Equal(t, -1, result, "Should return -1 for empty slice")
	})
}

func TestDistinct(t *testing.T) {
	t.Run("Distinct with duplicates", func(t *testing.T) {
		items := []int{1, 2, 2, 3, 3, 4, 5, 5}
		result := Distinct(items)
		assert.Equal(t, []int{1, 2, 3, 4, 5}, result, "Should return unique elements")
	})

	t.Run("Distinct without duplicates", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5}
		result := Distinct(items)
		assert.Equal(t, []int{1, 2, 3, 4, 5}, result, "Should return same slice when no duplicates")
	})

	t.Run("Distinct empty slice", func(t *testing.T) {
		var items []int
		result := Distinct(items)
		assert.Empty(t, result, "Should return empty slice for empty input")
	})
}

func TestIntersect(t *testing.T) {
	t.Run("Intersect with common elements", func(t *testing.T) {
		items1 := []int{1, 2, 3, 4, 5}
		items2 := []int{4, 5, 6, 7, 8}
		result := Intersect(items1, items2)
		assert.Equal(t, []int{4, 5}, result, "Should return common elements")
	})

	t.Run("Intersect without common elements", func(t *testing.T) {
		items1 := []int{1, 2, 3}
		items2 := []int{4, 5, 6}
		result := Intersect(items1, items2)
		assert.Empty(t, result, "Should return empty slice when no common elements")
	})

	t.Run("Intersect with empty slice", func(t *testing.T) {
		items1 := []int{1, 2, 3}
		var items2 []int
		result := Intersect(items1, items2)
		assert.Empty(t, result, "Should return empty slice when one input is empty")
	})
}

func TestMap(t *testing.T) {
	t.Run("Map integers", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5}
		result := Map(items, func(i int) int { return i * 2 })
		assert.Equal(t, []int{2, 4, 6, 8, 10}, result, "Should apply function to all elements")
	})

	t.Run("Map strings", func(t *testing.T) {
		items := []string{"a", "b", "c"}
		result := Map(items, func(s string) string { return s + s })
		assert.Equal(t, []string{"aa", "bb", "cc"}, result, "Should apply function to all elements")
	})

	t.Run("Map empty slice", func(t *testing.T) {
		var items []int
		result := Map(items, func(i int) int { return i * 2 })
		assert.Empty(t, result, "Should return empty slice for empty input")
	})
}

func TestFilter(t *testing.T) {
	t.Run("Filter even numbers", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5}
		result := Filter(items, func(i int) bool { return i%2 == 0 })
		assert.Equal(t, []int{2, 4}, result, "Should return even numbers")
	})

	t.Run("Filter all elements", func(t *testing.T) {
		items := []int{1, 3, 5, 7, 9}
		result := Filter(items, func(i int) bool { return i > 0 })
		assert.Equal(t, items, result, "Should return all elements when all satisfy predicate")
	})

	t.Run("Filter no elements", func(t *testing.T) {
		items := []int{1, 3, 5, 7, 9}
		result := Filter(items, func(i int) bool { return i > 10 })
		assert.Empty(t, result, "Should return empty slice when no elements satisfy predicate")
	})
}

func TestChunk(t *testing.T) {
	t.Run("Chunk with exact division", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5, 6}
		result := Chunk(items, 2)
		assert.Equal(t, [][]int{{1, 2}, {3, 4}, {5, 6}}, result, "Should split into equal chunks")
	})

	t.Run("Chunk with remainder", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5, 6, 7}
		result := Chunk(items, 3)
		assert.Equal(t, [][]int{{1, 2, 3}, {4, 5, 6}, {7}}, result, "Should split with last chunk smaller")
	})

	t.Run("Chunk empty slice", func(t *testing.T) {
		var items []int
		result := Chunk(items, 2)
		assert.Empty(t, result, "Should return empty slice for empty input")
	})
}

func TestSum(t *testing.T) {
	t.Run("Sum positive numbers", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5}
		result := Sum(items)
		assert.Equal(t, 15, result, "Should return sum of all elements")
	})

	t.Run("Sum with negative numbers", func(t *testing.T) {
		items := []int{-1, 2, -3, 4, -5}
		result := Sum(items)
		assert.Equal(t, -3, result, "Should return correct sum with negative numbers")
	})

	t.Run("Sum empty slice", func(t *testing.T) {
		var items []int
		result := Sum(items)
		assert.Equal(t, 0, result, "Should return 0 for empty slice")
	})
}

func TestAverage(t *testing.T) {
	t.Run("Average of integers", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5}
		result := Average(items)
		assert.Equal(t, 3.0, result, "Should return correct average")
	})

	t.Run("Average of floats", func(t *testing.T) {
		items := []float64{1.5, 2.5, 3.5}
		result := Average(items)
		assert.Equal(t, 2.5, result, "Should return correct average for floats")
	})

	t.Run("Average of empty slice", func(t *testing.T) {
		var items []int
		result := Average(items)
		assert.Equal(t, 0.0, result, "Should return 0 for empty slice")
	})
}

func TestZip(t *testing.T) {
	t.Run("Zip equal length slices", func(t *testing.T) {
		items1 := []int{1, 2, 3}
		items2 := []string{"one", "two", "three"}
		result := Zip(items1, items2, func(i int, s string) string {
			return string(rune('0'+i)) + "-" + s
		})
		assert.Equal(t, []string{"1-one", "2-two", "3-three"}, result, "Should combine elements correctly")
	})

	t.Run("Zip different length slices", func(t *testing.T) {
		items1 := []int{1, 2, 3, 4}
		items2 := []string{"one", "two"}
		result := Zip(items1, items2, func(i int, s string) string {
			return string(rune('0'+i)) + "-" + s
		})
		assert.Equal(t, []string{"1-one", "2-two"}, result, "Should zip up to the shorter slice length")
	})

	t.Run("Zip with empty slice", func(t *testing.T) {
		items1 := []int{1, 2, 3}
		var items2 []string
		result := Zip(items1, items2, func(i int, s string) string {
			return string(rune('0'+i)) + "-" + s
		})
		assert.Empty(t, result, "Should return empty slice when one input is empty")
	})
}

func TestFindLast(t *testing.T) {
	t.Run("Find last existing element", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 3, 5}
		result := FindLast(items, 3)
		assert.Equal(t, 3, result, "Should find last occurrence of 3")
	})

	t.Run("Find last non-existing element", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5}
		result := FindLast(items, 6)
		assert.Equal(t, 0, result, "Should return zero value when element not found")
	})

	t.Run("Find last in empty slice", func(t *testing.T) {
		var items []int
		result := FindLast(items, 1)
		assert.Equal(t, 0, result, "Should return zero value for empty slice")
	})
}

func TestFindLastIndex(t *testing.T) {
	t.Run("Find last index of existing element", func(t *testing.T) {
		items := []string{"apple", "banana", "cherry", "banana"}
		result := FindLastIndex(items, "banana")
		assert.Equal(t, 3, result, "Should return last index of 'banana'")
	})

	t.Run("Find last index of non-existing element", func(t *testing.T) {
		items := []string{"apple", "banana", "cherry"}
		result := FindLastIndex(items, "date")
		assert.Equal(t, -1, result, "Should return -1 for non-existing element")
	})

	t.Run("Find last index in empty slice", func(t *testing.T) {
		var items []string
		result := FindLastIndex(items, "apple")
		assert.Equal(t, -1, result, "Should return -1 for empty slice")
	})
}

func TestContains(t *testing.T) {
	t.Run("Contains existing element", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5}
		result := Contains(items, 3)
		assert.True(t, result, "Should return true for existing element")
	})

	t.Run("Contains non-existing element", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5}
		result := Contains(items, 6)
		assert.False(t, result, "Should return false for non-existing element")
	})

	t.Run("Contains in empty slice", func(t *testing.T) {
		var items []int
		result := Contains(items, 1)
		assert.False(t, result, "Should return false for empty slice")
	})
}

func TestExcept(t *testing.T) {
	t.Run("Except with overlapping elements", func(t *testing.T) {
		items1 := []int{1, 2, 3, 4, 5}
		items2 := []int{3, 4, 5, 6, 7}
		result := Except(items1, items2)
		assert.Equal(t, []int{1, 2}, result, "Should return elements not in second slice")
	})

	t.Run("Except with no overlapping elements", func(t *testing.T) {
		items1 := []int{1, 2, 3}
		items2 := []int{4, 5, 6}
		result := Except(items1, items2)
		assert.Equal(t, []int{1, 2, 3}, result, "Should return all elements from first slice")
	})

	t.Run("Except with empty first slice", func(t *testing.T) {
		var items1 []int
		items2 := []int{1, 2, 3}
		result := Except(items1, items2)
		assert.Empty(t, result, "Should return empty slice")
	})
}

func TestUnion(t *testing.T) {
	t.Run("Union with overlapping elements", func(t *testing.T) {
		items1 := []int{1, 2, 3, 4}
		items2 := []int{3, 4, 5, 6}
		result := Union(items1, items2)
		assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, result, "Should return union of both slices")
	})

	t.Run("Union with no overlapping elements", func(t *testing.T) {
		items1 := []int{1, 2, 3}
		items2 := []int{4, 5, 6}
		result := Union(items1, items2)
		assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, result, "Should return all elements from both slices")
	})

	t.Run("Union with empty slice", func(t *testing.T) {
		items1 := []int{1, 2, 3}
		var items2 []int
		result := Union(items1, items2)
		assert.Equal(t, []int{1, 2, 3}, result, "Should return elements from non-empty slice")
	})
}

func TestTake(t *testing.T) {
	t.Run("Take less than slice length", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5}
		result := Take(items, 3)
		assert.Equal(t, []int{1, 2, 3}, result, "Should return first 3 elements")
	})

	t.Run("Take more than slice length", func(t *testing.T) {
		items := []int{1, 2, 3}
		result := Take(items, 5)
		assert.Equal(t, []int{1, 2, 3}, result, "Should return all elements")
	})

	t.Run("Take from empty slice", func(t *testing.T) {
		var items []int
		result := Take(items, 3)
		assert.Empty(t, result, "Should return empty slice")
	})
}

func TestSkip(t *testing.T) {
	t.Run("Skip less than slice length", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5}
		result := Skip(items, 2)
		assert.Equal(t, []int{3, 4, 5}, result, "Should skip first 2 elements")
	})

	t.Run("Skip more than slice length", func(t *testing.T) {
		items := []int{1, 2, 3}
		result := Skip(items, 5)
		assert.Empty(t, result, "Should return empty slice")
	})

	t.Run("Skip from empty slice", func(t *testing.T) {
		var items []int
		result := Skip(items, 3)
		assert.Empty(t, result, "Should return empty slice")
	})
}

func TestAll(t *testing.T) {
	t.Run("All elements satisfy condition", func(t *testing.T) {
		items := []int{2, 4, 6, 8}
		result := All(items, func(i int) bool { return i%2 == 0 })
		assert.True(t, result, "Should return true when all elements are even")
	})

	t.Run("Not all elements satisfy condition", func(t *testing.T) {
		items := []int{2, 4, 5, 8}
		result := All(items, func(i int) bool { return i%2 == 0 })
		assert.False(t, result, "Should return false when not all elements are even")
	})

	t.Run("All with empty slice", func(t *testing.T) {
		var items []int
		result := All(items, func(i int) bool { return i > 0 })
		assert.True(t, result, "Should return true for empty slice")
	})
}

func TestAny(t *testing.T) {
	t.Run("Some elements satisfy condition", func(t *testing.T) {
		items := []int{1, 3, 4, 7}
		result := Any(items, func(i int) bool { return i%2 == 0 })
		assert.True(t, result, "Should return true when at least one element is even")
	})

	t.Run("No elements satisfy condition", func(t *testing.T) {
		items := []int{1, 3, 5, 7}
		result := Any(items, func(i int) bool { return i%2 == 0 })
		assert.False(t, result, "Should return false when no elements are even")
	})

	t.Run("Any with empty slice", func(t *testing.T) {
		var items []int
		result := Any(items, func(i int) bool { return i > 0 })
		assert.False(t, result, "Should return false for empty slice")
	})
}

func TestAt(t *testing.T) {
	t.Run("Get element at index", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5}
		result := At(items, 2)
		assert.Equal(t, 3, result, "Should return element at index 2")
	})

	t.Run("Get element at index out of bounds", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5}
		result := At(items, 10)
		assert.Equal(t, 0, result, "Should return zero value for out of bounds index")
	})

	t.Run("Get element at negative index", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5}
		result := At(items, -1)
		assert.Equal(t, 0, result, "Should return zero value for negative index")
	})

	t.Run("Get element at index of empty slice", func(t *testing.T) {
		var items []int
		result := At(items, 0)
		assert.Equal(t, 0, result, "Should return zero value for empty slice")
	})
}
