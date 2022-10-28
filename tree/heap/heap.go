package heap

type (
	// Heap - A Tree data structure
	Heap[T comparable] struct {
		nodes []T
	}
)

// New returns an empty Heap
func New[T comparable]() *Heap[T] {
	return &Heap[T]{
		nodes: []T{},
	}
}
