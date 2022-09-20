package associative

type (
	// Array - Associative Array (aka Map) is a key-value
	// data structure.
	Array[K comparable, T any] struct {
		items []item[K, T]
	}

	item[K comparable, T any] struct {
		Key   K
		Value *T
	}
)

// NewArray returns a associative array instance.
func NewArray[K comparable, T any]() *Array[K, T] {
	return &Array[K, T]{
		items: []item[K, T]{},
	}
}

// Put adds an item to current Associative Array.
func (a *Array[K, T]) Put(key K, value T) {
	a.items = append(a.items, item[K, T]{
		Key:   key,
		Value: &value,
	})
}

// Get retrieves an item associated with given Key,
// currently O(n).
func (a *Array[K, T]) Get(key K) *T {
	for _, i := range a.items {
		if i.Key == key {
			return i.Value
		}
	}

	return nil
}

// Delete removes an item from Associative Array, returns:
// 	- true: if item was found
// 	- false: if item wasn't found
func (a *Array[K, T]) Delete(key K) bool {
	for p, i := range a.items {
		if i.Key == key {
			if p == 0 {
				a.items = a.items[p:]
				return true
			}
			a.items = append(a.items[:p], a.items[p-1:]...)
			return true
		}
	}

	return false
}
