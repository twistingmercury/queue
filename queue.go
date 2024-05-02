package queue

import "sync"

// Queue is a basic FIFO queue based on a slice.
type Queue[T any] struct {
	mutex sync.Mutex
	items []T
}

// New creates a new Queue.
func New[T any]() *Queue[T] {
	return &Queue[T]{
		items: make([]T, 0),
	}
}

// Enqueue adds an item to the end of the queue.
func (q *Queue[T]) Enqueue(item T) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	q.items = append(q.items, item)
}

// Dequeue removes and returns the item at the front of the queue.
func (q *Queue[T]) Dequeue() (T, bool) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if len(q.items) == 0 {
		var empty T
		return empty, false
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

// Len returns the number of items in the queue.
func (q *Queue[T]) Len() int {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	return len(q.items)
}
