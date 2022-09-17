package list

type (
	// List aka Linked List
	List[T any] struct {
		item T
		next *List[T]
	}
)

// New returns a new List with given type T.
func New[T any]() *List[T] {
	return &List[T]{}
}
