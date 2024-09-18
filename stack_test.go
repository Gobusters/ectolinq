package ectolinq

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewStack(t *testing.T) {
	t.Run("Create new stack", func(t *testing.T) {
		s := NewStack[int]()
		assert.NotNil(t, s)
		assert.Equal(t, 0, s.Count())
	})
}

func TestFromSlice(t *testing.T) {
	t.Run("Create stack from slice", func(t *testing.T) {
		slice := []int{1, 2, 3}
		s := FromSlice(slice)
		assert.NotNil(t, s)
		assert.Equal(t, 3, s.Count())
		assert.Equal(t, slice, s.ToArray())
	})
}

func TestStackPush(t *testing.T) {
	t.Run("Push items to stack", func(t *testing.T) {
		s := NewStack[int]()
		s.Push(1)
		s.Push(2)
		assert.Equal(t, 2, s.Count())
	})
}

func TestStackPop(t *testing.T) {
	t.Run("Pop from non-empty stack", func(t *testing.T) {
		s := FromSlice([]int{1, 2, 3})
		item, err := s.Pop()
		require.NoError(t, err)
		assert.Equal(t, 3, item)
		assert.Equal(t, 2, s.Count())
	})

	t.Run("Pop from empty stack", func(t *testing.T) {
		s := NewStack[int]()
		_, err := s.Pop()
		assert.Error(t, err)
	})
}

func TestStackPeek(t *testing.T) {
	t.Run("Peek from non-empty stack", func(t *testing.T) {
		s := FromSlice([]int{1, 2, 3})
		item, err := s.Peek()
		require.NoError(t, err)
		assert.Equal(t, 3, item)
		assert.Equal(t, 3, s.Count()) // Ensure count hasn't changed
	})

	t.Run("Peek from empty stack", func(t *testing.T) {
		s := NewStack[int]()
		_, err := s.Peek()
		assert.Error(t, err)
	})
}

func TestStackCount(t *testing.T) {
	t.Run("Count items in stack", func(t *testing.T) {
		s := FromSlice([]int{1, 2, 3})
		assert.Equal(t, 3, s.Count())

		s.Push(4)
		assert.Equal(t, 4, s.Count())

		_, _ = s.Pop()
		assert.Equal(t, 3, s.Count())
	})
}

func TestStackClear(t *testing.T) {
	t.Run("Clear stack", func(t *testing.T) {
		s := FromSlice([]int{1, 2, 3})
		s.Clear()
		assert.Equal(t, 0, s.Count())
	})
}

func TestStackContains(t *testing.T) {
	t.Run("Check if stack contains item", func(t *testing.T) {
		s := FromSlice([]int{1, 2, 3})
		assert.True(t, s.Contains(2))
		assert.False(t, s.Contains(4))
	})
}

func TestStackContainsWhere(t *testing.T) {
	t.Run("Check if stack contains item satisfying condition", func(t *testing.T) {
		s := FromSlice([]int{1, 2, 3})
		assert.True(t, s.ContainsWhere(func(i int) bool { return i > 2 }))
		assert.False(t, s.ContainsWhere(func(i int) bool { return i > 3 }))
	})
}

func TestStackToArray(t *testing.T) {
	t.Run("Convert stack to array", func(t *testing.T) {
		original := []int{1, 2, 3}
		s := FromSlice(original)
		array := s.ToArray()
		assert.Equal(t, original, array)
	})
}

func TestStackToList(t *testing.T) {
	t.Run("Convert stack to list", func(t *testing.T) {
		original := []int{1, 2, 3}
		s := FromSlice(original)
		list := s.ToList()
		assert.Equal(t, ToList(original), list)
	})
}
