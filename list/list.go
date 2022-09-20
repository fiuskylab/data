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

func (l *List[T]) Last() *T {
	current := l
	for {
		if current.Next == nil {
			return current.Value
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

// ListAt - returns the List instance at position `p`.
func (l *List[T]) ListAt(p int) *List[T] {
	current := *l
	if l.Next == nil && l.Value == nil {
		return nil
	}
	for i := 0; i <= p; i++ {
		if i == p {
			return current.Next
		}
		if current.Next == nil {
			break
		}
		current = *current.Next
	}
	return nil
}

// Pop - removes a value from the Linked List.
func (l *List[T]) Pop(p int) *T {
	current := *l
	currentPtr := l
	if l.Next == nil && l.Value == nil {
		return nil
	}
	for i := 0; i <= p; i++ {
		if i == p {
			val := *current.Value
			if currentPtr.Next != nil {
				currentPtr.Value = currentPtr.Next.Value
				currentPtr.Next = currentPtr.Next.Next
			} else {
				currentPtr.Value = nil
			}
			return &val
		}
		if current.Next == nil {
			break
		}
		current = *current.Next
		currentPtr = currentPtr.Next
	}
	return nil
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
func (l *List[T]) ValueAt(p int) *T {
	current := *l
	if l.Value == nil {
		return nil
	}
	for i := 0; i <= p; i++ {
		if i == p {
			return current.Value
		}
		if current.Next == nil {
			break
		}
		current = *current.Next
	}
	return nil
}
