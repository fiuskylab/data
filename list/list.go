package list

type (
	// List aka Linked List
	List[T any] struct {
		value *T
		next  *List[T]
	}
)

// New returns a new List with given type T.
func New[T any]() *List[T] {
	return &List[T]{}
}

// Append will Append a value
func (l *List[T]) Append(v T) {
	if l.value == nil {
		l.value = &v
		return
	}

	l.next = &List[T]{
		value: &v,
	}
}
