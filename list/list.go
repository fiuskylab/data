package list

import (
	"fmt"
)

type (
	// List aka Linked List.
	List[T any] struct {
		Value *T
		Next  *List[T]
	}
)

// New returns a new List with given type T.
func New[T any]() *List[T] {
	return &List[T]{}
}

// Append will Append a Value.
func (l *List[T]) Append(v T) {
	if l.Value == nil {
		l.Value = &v
		return
	}

	l.Next = &List[T]{
		Value: &v,
	}
}

// Empty returns whether the List is empty or not.
func (l *List[T]) Empty() bool {
	return (l.Value == nil) && (l.Next == nil)
}

// Head returns List's head.
func (l *List[T]) Head() *T {
	return l.Value
}

func (l *List[T]) lastInstance() *List[T] {
	current := l
	for {
		if current.Next == nil {
			return current
		}
		current = current.Next
	}
}

// Len return the amount of item in the List.
func (l *List[T]) Len() int {
	i := 0
	current := l
	for {
		if current.Next == nil {
			return i
		}
		i++
		current = current.Next
	}
}

// Prepend will add a Value to end of the List.
func (l *List[T]) Prepend(v T) {
	currentList := *l
	newL := New[T]()

	newL.Value = &v
	newL.Next = &currentList

	*l = *newL
}

// Tail - returns the List without the first item.
func (l *List[T]) Tail() *List[T] {
	return l.Next
}

// ValueAt - Returns a Value at position p.
func (l *List[T]) ValueAt(p int) (T, error) {
	current := *l
	if l.Value == nil {
		return *new(T), fmt.Errorf(errEmptyList)
	}
	for i := 0; i <= p; i++ {
		if i == p {
			return *current.Value, nil
		}
		if current.Next == nil {
			break
		}
		current = *current.Next
	}
	return *new(T), fmt.Errorf(errShorterThan, p)
}
