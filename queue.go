package ectolinq

type Queue[T any] struct {
	items List[T]
}

// NewQueue creates a new queue
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		items: NewList[T](),
	}
}

// ToQueue creates a new queue from an array
// items: The array to create the queue from
func ToQueue[T any](items []T) *Queue[T] {
	return &Queue[T]{
		items: items,
	}
}

// Enqueue adds an item to the end of the queue
// item: The item to add
func (q *Queue[T]) Enqueue(item T) {
	q.items.Push(item)
}

// Dequeue removes and returns the item at the beginning of the queue
func (q *Queue[T]) Dequeue() T {
	item, items := q.items.Pop()
	q.items = items
	return item
}

// Peek returns an element that is at the beginning of the queue without removing it
func (q *Queue[T]) Peek() T {
	return q.items.First()
}

// Count returns the number of items in the queue
func (q *Queue[T]) Count() int {
	return q.items.Length()
}

// Clear removes all items from the queue
func (q *Queue[T]) Clear() {
	q.items = NewList[T]()
}

// Contains returns if the queue contains the given item
// item: The item to check for
func (q *Queue[T]) Contains(item T) bool {
	return q.items.Contains(item)
}

// ContainsWhere returns if the queue contains an item that satisfies the given predicate
// fn: The predicate to check for
func (q *Queue[T]) ContainsWhere(fn func(T) bool) bool {
	return q.items.Any(fn)
}

// ToArray returns the items in the queue as an array
func (q *Queue[T]) ToArray() []T {
	return q.items
}

// ToList returns the items in the queue as a list
func (q *Queue[T]) ToList() List[T] {
	return q.items
}
