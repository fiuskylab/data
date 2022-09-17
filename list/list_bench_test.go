package list

import "testing"

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New[int]()
	}
}

func BenchmarkListAppend(b *testing.B) {
	list := New[int]()
	for i := 0; i < b.N; i++ {
		list.Append(i)
	}
}

func BenchmarkSliceAppend(b *testing.B) {
	l := []int{}
	for i := 0; i < b.N; i++ {
		l = append(l, i)
	}
}

func BenchmarkListPrepend(b *testing.B) {
	list := New[int]()
	for i := 0; i < b.N; i++ {
		list.Prepend(i)
	}
}

func BenchmarkSlicePrepend(b *testing.B) {
	l := []int{}
	for i := 0; i < b.N; i++ {
		l = append([]int{i}, l...)
	}
}
