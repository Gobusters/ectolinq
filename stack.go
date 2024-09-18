package ectolinq

import "fmt"

type Stack[T any] struct {
	items List[T]
}

// NewStack creates a new stack
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		items: NewList[T](),
	}
}

// Rename ToStack to FromSlice for consistency
func FromSlice[T any](items []T) *Stack[T] {
	return &Stack[T]{
		items: items,
	}
}

// Push pushes an item onto the stack
// item: The item to push
func (s *Stack[T]) Push(item T) {
	s.items = s.items.Push(item)
}

// Pop removes and returns the item on the top of the stack
func (s *Stack[T]) Pop() (T, error) {
	if s.Count() == 0 {
		var zero T
		return zero, fmt.Errorf("cannot pop from an empty stack")
	}
	item, items := s.items.Pop()
	s.items = items
	return item, nil
}

// Peek returns an element that is at the top of the stack without removing it
func (s *Stack[T]) Peek() (T, error) {
	if s.Count() == 0 {
		var zero T
		return zero, fmt.Errorf("cannot peek an empty stack")
	}
	return s.items.Last(), nil
}

// Count returns the number of items in the stack
func (s *Stack[T]) Count() int {
	return s.items.Length()
}

// Clear removes all items from the stack
func (s *Stack[T]) Clear() {
	s.items = NewList[T]()
}

// Contains returns if the stack contains the given item
// item: The item to check for
func (s *Stack[T]) Contains(item T) bool {
	return s.items.Contains(item)
}

// ContainsWhere returns if the stack contains an item that satisfies the given predicate
// fn: The predicate to check for
func (s *Stack[T]) ContainsWhere(fn func(T) bool) bool {
	return s.items.Any(fn)
}

// ToArray returns the items in the stack as an array
func (s *Stack[T]) ToArray() []T {
	return s.items
}

// ToList returns the items in the stack as a list
func (s *Stack[T]) ToList() List[T] {
	return s.items
}
