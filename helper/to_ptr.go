package helper

// ToPtr receives a value T and returns it's address.
func ToPtr[T any](v T) *T {
	return &v
}
