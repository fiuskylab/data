package queue

import "github.com/fiuskylab/data/list"

type (
	// Queue - A collection of entities that are maintained in a sequence,
	// in this case a FIFO (first in, first out).
	Queue[T any] struct {
		l *list.List[T]
	}
)

// New returns a new instance of Queue.
func New[T any]() *Queue[T] {
	return &Queue[T]{
		l: list.New[T](),
	}
}

// Push - Add a item to the end of the queue.
func (q *Queue[T]) Push(v T) {
	q.l.Append(v)
}

// Get - Retrieves the Queue's first item, but
// doesn't delete it, if you call `Get()` again,
// it will return the same value as the first call.
func (q *Queue[T]) Get() *T {
	return q.l.Value
}

// Pop - Retrieves Queue's first item and
// removes it.
func (q *Queue[T]) Pop() *T {
	return q.l.Pop(0)
}
