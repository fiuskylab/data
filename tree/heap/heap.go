package heap

type (
	// Heap - A Tree data structure
	Heap[T comparable] struct {
		nodes []node[T]
	}

	// node - basic unit of data structure.
	node[T comparable] struct {
		Value *T
		Left  *node[T]
		Right *node[T]
	}
)

// New returns an empty Heap
func New[T comparable]() *Heap[T] {
	return &Heap[T]{
		nodes: []node[T]{},
	}
}
