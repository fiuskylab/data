package associative

type (
	// Array - Associative Array (aka Map) is a key-value
	// data structure.
	Array[K comparable, T any] struct {
		items []item[K, T]
	}

	item[K comparable, T any] struct {
		key   K
		value *T
	}
)

// NewArray returns a associative array instance.
func NewArray[K comparable, T any]() *Array[K, T] {
	return &Array[K, T]{
		items: make([]item[K, T], 10),
	}
}

// Put adds an item to current Associative Array.
func (a *Array[K, T]) Put(key K, value T) {
	a.items = append(a.items, item[K, T]{
		key:   key,
		value: &value,
	})
}

// Get retrieves an item associated with given Key.
func (a *Array[K, T]) Get(key K) *T {
	for _, i := range a.items {
		if i.key == key {
			return i.value
		}
	}

	return nil
}

// Delete removes an item from Associative Array, returns:
// 	- true: if item was found
// 	- false: if item wasn't found
func (a *Array[K, T]) Delete(key K) bool {
	for p, i := range a.items {
		if i.key == key {
			a.items = append(a.items[0:p], a.items[p-1:]...)
			return true
		}
	}

	return false
}
