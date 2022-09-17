package list

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

// Prepend will add a Value to end of the List.
func (l *List[T]) Prepend(v T) {
	currentList := *l
	newL := New[T]()

	newL.Value = &v
	newL.Next = &currentList

	*l = *newL
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
