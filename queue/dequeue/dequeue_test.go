package dequeue

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	require := require.New(t)

	actual := New[int]()

	require.NotNil(actual)
	require.NotNil(actual.List)
	require.Nil(actual.Value)
	require.Nil(actual.Next)
}

func TestPushBack(t *testing.T) {
	d := New[int]()

	t.Run("1 Item", func(t *testing.T) {
		require := require.New(t)

		expected := 10

		d.PushBack(expected)

		require.Equal(expected, *d.Value)
	})

	t.Run("2 Items", func(t *testing.T) {
		require := require.New(t)

		expected := 20

		d.PushBack(expected)

		require.Equal(10, *d.Value)
		require.Equal(expected, *d.Next.Value)
	})
}

func TestPushFront(t *testing.T) {
	d := New[int]()

	t.Run("1 Item", func(t *testing.T) {
		require := require.New(t)

		expected := 10

		d.PushFront(expected)

		require.Equal(expected, *d.Value)
	})

	t.Run("2 Items", func(t *testing.T) {
		require := require.New(t)

		expected := 20

		d.PushFront(expected)

		require.Equal(10, *d.Next.Value)
		require.Equal(expected, *d.Value)
	})
}

func TestPopBack(t *testing.T) {
	d := New[int]()

	t.Run("Empty", func(t *testing.T) {
		require := require.New(t)

		actual := d.PopBack()

		require.Nil(actual)
	})

	t.Run("With 1 Item", func(t *testing.T) {
		require := require.New(t)

		expected := 10
		d.PushBack(expected)

		actual := d.PopFront()

		require.NotNil(actual)
		require.Equal(expected, *actual)
	})

	t.Run("With 2 Items", func(t *testing.T) {
		require := require.New(t)

		expected := 10
		d.PushBack(expected)
		d.PushBack(20)

		actual := d.PopFront()

		require.NotNil(actual)
		require.Equal(expected, *actual)
	})

}

func TestPopFront(t *testing.T) {
	d := New[int]()

	t.Run("Empty", func(t *testing.T) {
		require := require.New(t)

		actual := d.PopFront()

		require.Nil(actual)
	})

	t.Run("With 1 Item", func(t *testing.T) {
		require := require.New(t)

		expected := 10
		d.PushFront(expected)

		actual := d.PopFront()

		require.NotNil(actual)
		require.Equal(expected, *actual)
	})

	t.Run("With 2 Items", func(t *testing.T) {
		require := require.New(t)

		expected := 10
		d.PushFront(20)
		d.PushFront(expected)

		actual := d.PopFront()

		require.NotNil(actual)
		require.Equal(expected, *actual)
	})
}

func TestFront(t *testing.T) {
	d := New[int]()

	t.Run("Empty Dequeue", func(t *testing.T) {
		require := require.New(t)

		actual := d.Front()

		require.Nil(actual)
	})

	t.Run("With 1 Item", func(t *testing.T) {
		require := require.New(t)

		expected := 10
		d.PushBack(expected)
		actual := d.Front()

		require.Equal(expected, *actual)

		actual = d.Front()
		require.Equal(expected, *actual)
	})
}

func TestBack(t *testing.T) {
	d := New[int]()

	t.Run("Empty Dequeue", func(t *testing.T) {
		require := require.New(t)

		actual := d.Back()

		require.Nil(actual)
	})

	t.Run("With 1 Item", func(t *testing.T) {
		require := require.New(t)

		expected := 10
		d.PushBack(expected)
		actual := d.Back()

		require.Equal(expected, *actual)

		actual = d.Back()
		require.Equal(expected, *actual)
	})

	t.Run("With 2 Items", func(t *testing.T) {
		require := require.New(t)

		expected := 20
		d.PushBack(expected)
		actual := d.Back()

		require.Equal(expected, *actual)

		actual = d.Back()
		require.Equal(expected, *actual)
	})
}
