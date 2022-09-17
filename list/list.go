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
