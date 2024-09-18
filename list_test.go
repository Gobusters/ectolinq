package ectolinq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewList(t *testing.T) {
	t.Run("Empty list", func(t *testing.T) {
		list := NewList[int]()
		assert.Equal(t, 0, list.Length(), "New list should be empty")
	})
}

func TestFrom(t *testing.T) {
	t.Run("From slice", func(t *testing.T) {
		items := []int{1, 2, 3}
		list := ToList(items)
		assert.Equal(t, 3, list.Length(), "List should have same length as input slice")
		assert.Equal(t, items, []int(list), "List should contain same elements as input slice")
	})
}

func TestListForEach(t *testing.T) {
	t.Run("Apply function to all elements", func(t *testing.T) {
		list := ToList([]int{1, 2, 3})
		sum := 0
		list.ForEach(func(i int) {
			sum += i
		})
		assert.Equal(t, 6, sum, "ForEach should apply function to all elements")
	})
}

func TestListFind(t *testing.T) {
	t.Run("Find existing element", func(t *testing.T) {
		list := ToList([]int{1, 2, 3, 4, 5})
		result := list.Find(func(i int) bool { return i > 3 })
		assert.Equal(t, 4, result, "Should find first element greater than 3")
	})
}

func TestListReverse(t *testing.T) {
	t.Run("Reverse list", func(t *testing.T) {
		list := ToList([]int{1, 2, 3, 4, 5})
		reversed := list.Reverse()
		assert.Equal(t, []int{5, 4, 3, 2, 1}, []int(reversed), "Should reverse the list")
	})
}

func TestListFindIndex(t *testing.T) {
	t.Run("Find index of existing element", func(t *testing.T) {
		list := ToList([]string{"apple", "banana", "cherry"})
		index := list.FindIndex("banana")
		assert.Equal(t, 1, index, "Should return correct index for existing element")
	})
}

func TestListFindLastIndex(t *testing.T) {
	t.Run("Find last index of existing element", func(t *testing.T) {
		list := ToList([]string{"apple", "banana", "cherry", "banana"})
		index := list.FindLastIndex("banana")
		assert.Equal(t, 3, index, "Should return last index of 'banana'")
	})
}

func TestListContains(t *testing.T) {
	t.Run("Contains existing element", func(t *testing.T) {
		list := ToList([]int{1, 2, 3, 4, 5})
		assert.True(t, list.Contains(3), "Should return true for existing element")
	})

	t.Run("Does not contain non-existing element", func(t *testing.T) {
		list := ToList([]int{1, 2, 3, 4, 5})
		assert.False(t, list.Contains(6), "Should return false for non-existing element")
	})
}

func TestListAll(t *testing.T) {
	t.Run("All elements satisfy condition", func(t *testing.T) {
		list := ToList([]int{2, 4, 6, 8})
		allEven := list.All(func(i int) bool { return i%2 == 0 })
		assert.True(t, allEven, "All elements should be even")
	})
}

func TestListAny(t *testing.T) {
	t.Run("Some elements satisfy condition", func(t *testing.T) {
		list := ToList([]int{1, 3, 4, 7})
		hasEven := list.Any(func(i int) bool { return i%2 == 0 })
		assert.True(t, hasEven, "Should have at least one even number")
	})
}

func TestListCount(t *testing.T) {
	t.Run("Count elements satisfying condition", func(t *testing.T) {
		list := ToList([]int{1, 2, 3, 4, 5, 6})
		evenCount := list.Count(func(i int) bool { return i%2 == 0 })
		assert.Equal(t, 3, evenCount, "Should count 3 even numbers")
	})
}

func TestListDistinct(t *testing.T) {
	t.Run("Remove duplicates", func(t *testing.T) {
		list := ToList([]int{1, 2, 2, 3, 3, 4, 5, 5})
		distinct := list.Distinct()
		assert.Equal(t, []int{1, 2, 3, 4, 5}, []int(distinct), "Should return unique elements")
	})
}

func TestListExcept(t *testing.T) {
	t.Run("Remove elements from second list", func(t *testing.T) {
		list := ToList([]int{1, 2, 3, 4, 5})
		except := list.Except([]int{3, 4, 5, 6, 7})
		assert.Equal(t, []int{1, 2}, []int(except), "Should return elements not in second list")
	})
}

func TestListIntersect(t *testing.T) {
	t.Run("Find common elements", func(t *testing.T) {
		list := ToList([]int{1, 2, 3, 4, 5})
		intersect := list.Intersect([]int{4, 5, 6, 7, 8})
		assert.Equal(t, []int{4, 5}, []int(intersect), "Should return common elements")
	})
}

func TestListUnion(t *testing.T) {
	t.Run("Combine unique elements", func(t *testing.T) {
		list := ToList([]int{1, 2, 3, 4})
		union := list.Union([]int{3, 4, 5, 6})
		assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, []int(union), "Should return union of both lists")
	})
}

func TestListConcat(t *testing.T) {
	t.Run("Concatenate two lists", func(t *testing.T) {
		list1 := ToList([]int{1, 2, 3})
		list2 := ToList([]int{4, 5, 6})
		concat := list1.Concat(list2)
		assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, []int(concat), "Should concatenate two lists")
	})
}

func TestListSequenceEqual(t *testing.T) {
	t.Run("Equal sequences", func(t *testing.T) {
		list := ToList([]int{1, 2, 3, 4, 5})
		assert.True(t, list.SequenceEqual([]int{1, 2, 3, 4, 5}), "Should be equal to identical slice")
	})

	t.Run("Unequal sequences", func(t *testing.T) {
		list := ToList([]int{1, 2, 3, 4, 5})
		assert.False(t, list.SequenceEqual([]int{1, 2, 3, 4}), "Should not be equal to different slice")
	})
}

func TestListReduce(t *testing.T) {
	t.Run("Reduce to sum", func(t *testing.T) {
		list := ToList([]int{1, 2, 3, 4, 5})
		sum := list.Reduce(func(a, b int) int { return a + b })
		assert.Equal(t, 15, sum, "Should reduce to sum of all elements")
	})
}

func TestListReduceWhere(t *testing.T) {
	t.Run("Reduce with initial value", func(t *testing.T) {
		list := ToList([]string{"a", "b", "c"})
		concat := list.ReduceWhere("", func(a, b string) string { return a + b })
		assert.Equal(t, "abc", concat, "Should reduce to concatenation of all elements")
	})
}

func TestListFilter(t *testing.T) {
	t.Run("Filter even numbers", func(t *testing.T) {
		list := ToList([]int{1, 2, 3, 4, 5, 6})
		evens := list.Filter(func(i int) bool { return i%2 == 0 })
		assert.Equal(t, []int{2, 4, 6}, []int(evens), "Should filter even numbers")
	})
}

func TestListRemoveAt(t *testing.T) {
	t.Run("Remove element at index", func(t *testing.T) {
		list := ToList([]int{1, 2, 3, 4, 5})
		removed := list.RemoveAt(2)
		assert.Equal(t, []int{1, 2, 4, 5}, []int(removed), "Should remove element at index 2")
	})
}

func TestListRemove(t *testing.T) {
	t.Run("Remove first occurrence", func(t *testing.T) {
		list := ToList([]int{1, 2, 3, 4, 5})
		removed := list.Remove(3)
		assert.Equal(t, []int{1, 2, 4, 5}, []int(removed), "Should remove first occurrence of 3")
	})
}

func TestListRemoveWhere(t *testing.T) {
	t.Run("Remove first element satisfying condition", func(t *testing.T) {
		list := ToList([]int{1, 2, 3, 4, 5})
		removed := list.RemoveWhere(func(i int) bool { return i > 3 })
		assert.Equal(t, []int{1, 2, 3, 5}, []int(removed), "Should remove first element greater than 3")
	})
}

func TestListRandomize(t *testing.T) {
	t.Run("Randomize list", func(t *testing.T) {
		list := ToList([]int{1, 2, 3, 4, 5})
		randomized := list.Randomize()
		assert.NotEqual(t, []int{1, 2, 3, 4, 5}, []int(randomized), "Should randomize the list")
		assert.Equal(t, 5, randomized.Length(), "Should maintain the same length")
	})
}

func TestListSlice(t *testing.T) {
	t.Run("Slice list", func(t *testing.T) {
		list := ToList([]int{1, 2, 3, 4, 5})
		sliced := list.Slice(1, 4)
		assert.Equal(t, []int{2, 3, 4}, []int(sliced), "Should return sliced portion of list")
	})
}

func TestListSortWhere(t *testing.T) {
	t.Run("Sort list", func(t *testing.T) {
		list := ToList([]int{3, 1, 4, 1, 5, 9, 2, 6})
		sorted := list.SortWhere(func(a, b int) bool { return a < b })
		assert.Equal(t, []int{1, 1, 2, 3, 4, 5, 6, 9}, []int(sorted), "Should sort the list")
	})
}

func TestListTake(t *testing.T) {
	t.Run("Take first n elements", func(t *testing.T) {
		list := ToList([]int{1, 2, 3, 4, 5})
		taken := list.Take(3)
		assert.Equal(t, []int{1, 2, 3}, []int(taken), "Should take first 3 elements")
	})
}

func TestListTakeLast(t *testing.T) {
	t.Run("Take last n elements", func(t *testing.T) {
		list := ToList([]int{1, 2, 3, 4, 5})
		taken := list.TakeLast(3)
		assert.Equal(t, []int{3, 4, 5}, []int(taken), "Should take last 3 elements")
	})
}

func TestListTakeWhile(t *testing.T) {
	t.Run("Take while condition is true", func(t *testing.T) {
		list := ToList([]int{1, 2, 3, 4, 5})
		taken := list.TakeWhile(func(i int) bool { return i < 4 })
		assert.Equal(t, []int{1, 2, 3}, []int(taken), "Should take elements while less than 4")
	})
}

func TestListSkip(t *testing.T) {
	t.Run("Skip first n elements", func(t *testing.T) {
		list := ToList([]int{1, 2, 3, 4, 5})
		skipped := list.Skip(2)
		assert.Equal(t, []int{3, 4, 5}, []int(skipped), "Should skip first 2 elements")
	})
}

func TestListSkipLast(t *testing.T) {
	t.Run("Skip last n elements", func(t *testing.T) {
		list := ToList([]int{1, 2, 3, 4, 5})
		skipped := list.SkipLast(2)
		assert.Equal(t, []int{1, 2, 3}, []int(skipped), "Should skip last 2 elements")
	})
}

func TestListSkipWhile(t *testing.T) {
	t.Run("Skip while condition is true", func(t *testing.T) {
		list := ToList([]int{1, 2, 3, 4, 5})
		skipped := list.SkipWhile(func(i int) bool { return i < 3 })
		assert.Equal(t, []int{3, 4, 5}, []int(skipped), "Should skip elements while less than 3")
	})
}

func TestListChunk(t *testing.T) {
	t.Run("Chunk list", func(t *testing.T) {
		list := ToList([]int{1, 2, 3, 4, 5, 6, 7})
		chunks := list.Chunk(3)
		assert.Equal(t, []List[int]{{1, 2, 3}, {4, 5, 6}, {7}}, chunks, "Should chunk the list into groups of 3")
	})
}

func TestListPush(t *testing.T) {
	t.Run("Push element to list", func(t *testing.T) {
		list := ToList([]int{1, 2, 3})
		pushed := list.Push(4)
		assert.Equal(t, []int{1, 2, 3, 4}, []int(pushed), "Should add element to the end")
	})
}

func TestListPop(t *testing.T) {
	t.Run("Pop element from list", func(t *testing.T) {
		list := ToList([]int{1, 2, 3})
		item, popped := list.Pop()
		assert.Equal(t, 3, item, "Should return last element")
		assert.Equal(t, []int{1, 2}, []int(popped), "Should remove last element")
	})
}

func TestListUnshift(t *testing.T) {
	t.Run("Unshift element to list", func(t *testing.T) {
		list := ToList([]int{1, 2, 3})
		unshifted := list.Unshift(0)
		assert.Equal(t, []int{0, 1, 2, 3}, []int(unshifted), "Should add element to the beginning")
	})
}

func TestListShift(t *testing.T) {
	t.Run("Shift element from list", func(t *testing.T) {
		list := ToList([]int{1, 2, 3})
		item, shifted := list.Shift()
		assert.Equal(t, 1, item, "Should return first element")
		assert.Equal(t, []int{2, 3}, []int(shifted), "Should remove first element")
	})
}

func TestListLast(t *testing.T) {
	t.Run("Get last element", func(t *testing.T) {
		list := ToList([]int{1, 2, 3})
		last := list.Last()
		assert.Equal(t, 3, last, "Should return last element")
	})
}

func TestListFirst(t *testing.T) {
	t.Run("Get first element", func(t *testing.T) {
		list := ToList([]int{1, 2, 3})
		first := list.First()
		assert.Equal(t, 1, first, "Should return first element")
	})
}
