package dequeue

import "github.com/fiuskylab/data/list"

type (
	// Dequeue - A collection of entities that are maintained in a sequence,
	// in this case a FIFO (first in, first out).
	Dequeue[T any] struct {
		*list.List[T]
	}
)

// New returns a new instance of Dequeue.
func New[T any]() *Dequeue[T] {
	return &Dequeue[T]{
		list.New[T](),
	}
}

// PushBack - Add a item to the end of the queue.
func (q *Dequeue[T]) PushBack(v T) {
	q.Append(v)
}

// PushFront - Add a item to the start of the queue.
func (q *Dequeue[T]) PushFront(v T) {
	q.Prepend(v)
}

// PopBack - Retrieves and delete Queue's last item.
func (q *Dequeue[T]) PopBack() *T {
	return q.Pop(q.Len() - 1)
}

// PopFront - Retrieves and delete Queue's first item.
func (q *Dequeue[T]) PopFront() *T {
	return q.Pop(0)
}

// Front - Retrieves Queue's first item.
func (q *Dequeue[T]) Front() *T {
	return q.Value
}

// Back - Retrieves Queue's last item.
func (q *Dequeue[T]) Back() *T {
	return q.Last()
}
