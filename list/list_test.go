package list

import (
	"testing"

	"github.com/fiuskylab/data/helper"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	require := require.New(t)

	expected := List[int]{}
	actual := New[int]()

	require.Equal(expected, *actual)
}

func TestAppend(t *testing.T) {
	list := New[int]()

	t.Run("Empty List", func(t *testing.T) {
		require := require.New(t)
		expected := 1
		list.Append(expected)

		require.Equal(expected, *list.Value)
	})

	t.Run("Non Empty List", func(t *testing.T) {
		require := require.New(t)
		expected := 2
		list.Append(expected)

		require.Equal(expected, *list.Next.Value)
	})
}

func TestEmpty(t *testing.T) {
	list := New[int]()

	t.Run("Empty List", func(t *testing.T) {
		require := require.New(t)

		require.True(list.Empty())
	})

	t.Run("List Not Empty", func(t *testing.T) {
		require := require.New(t)

		list.Append(1)

		require.False(list.Empty())
	})
}

func TestHead(t *testing.T) {
	list := New[int]()

	t.Run("Empty List", func(t *testing.T) {
		require := require.New(t)

		require.Nil(list.Head())
	})

	t.Run("List Not Empty", func(t *testing.T) {
		require := require.New(t)

		expected := 1
		list.Append(expected)

		actual := list.Head()

		require.Equal(expected, *actual)
	})
}

func TestListAt(t *testing.T) {
	list := New[int]()
	t.Run("Empty List", func(t *testing.T) {
		require := require.New(t)

		actualVal := list.ListAt(100)

		require.Nil(actualVal)
	})

	t.Run("Shorter Than", func(t *testing.T) {
		require := require.New(t)

		pos := 100

		list.Append(1)

		actualVal := list.ListAt(pos)

		require.Nil(actualVal)
	})

	t.Run("Found Value", func(t *testing.T) {
		require := require.New(t)

		pos := 0

		list.Append(2)

		actualVal := list.ListAt(pos)

		require.NotNil(actualVal)
	})

	t.Run("Found Nil Value", func(t *testing.T) {
		require := require.New(t)

		pos := 1

		actualVal := list.ListAt(pos)

		require.Nil(actualVal)
	})
}

func TestPop(t *testing.T) {
	list := New[int]()
	t.Run("Empty List", func(t *testing.T) {
		require := require.New(t)

		actualVal := list.Pop(100)

		require.Nil(actualVal)
	})

	t.Run("Shorter Than", func(t *testing.T) {
		require := require.New(t)

		pos := 100

		list.Append(1)

		actual := list.Pop(pos)

		require.Nil(actual)
	})

	t.Run("Found Value", func(t *testing.T) {
		require := require.New(t)

		pos := 0

		list.Append(2)

		actualVal := list.Pop(pos)

		require.NotNil(actualVal)
	})

	t.Run("Found Nil Value", func(t *testing.T) {
		require := require.New(t)
		list.Append(3)

		pos := 1
		expectedVal := 3

		actualVal := list.Pop(pos)

		require.Equal(expectedVal, *actualVal)
	})
}

func TestPrepend(t *testing.T) {
	list := New[int]()

	t.Run("Empty List", func(t *testing.T) {
		require := require.New(t)

		expected := 1

		list.Prepend(expected)

		require.Equal(expected, *list.Value)
	})

	t.Run("List w/ 1 Item", func(t *testing.T) {
		require := require.New(t)

		expected := 2

		list.Prepend(expected)

		require.Equal(expected, *list.Value)
		require.Equal(1, *list.Next.Value)
	})
}

func TestLast(t *testing.T) {
	list := New[int]()

	t.Run("Empty List", func(t *testing.T) {
		require := require.New(t)

		actual := list.Last()

		require.Nil(actual)
	})

	t.Run("List w/ 1 item", func(t *testing.T) {
		require := require.New(t)

		expected := List[int]{
			Value: helper.ToPtr(1),
		}

		list.Append(1)

		actual := list.Last()

		require.Equal(*expected.Value, *actual)
	})

	for i := 2; i <= 10; i++ {
		list.Append(i)
	}

	t.Run("List w/ 10 item", func(t *testing.T) {
		require := require.New(t)

		expected := List[int]{
			Value: helper.ToPtr(10),
		}

		actual := list.Last()

		require.NotNil(actual)
		require.Equal(*expected.Value, *actual)
	})
}

func TestTail(t *testing.T) {
	list := New[int]()
	t.Run("Empty List", func(t *testing.T) {
		require := require.New(t)

		actual := list.Tail()

		require.Nil(actual)
	})

	t.Run("List w/ 1 item", func(t *testing.T) {
		require := require.New(t)

		list.Append(1)

		actual := list.Tail()

		require.Nil(actual)
	})

	t.Run("Non Nil Tail", func(t *testing.T) {
		require := require.New(t)

		list.Append(2)

		actual := list.Tail()

		require.NotNil(actual)
	})
}

func TestValueAt(t *testing.T) {
	list := New[int]()
	t.Run("Empty List", func(t *testing.T) {
		require := require.New(t)

		actual := list.ValueAt(100)

		require.Nil(actual)
	})

	t.Run("Shorter Than", func(t *testing.T) {
		require := require.New(t)

		pos := 100

		list.Append(1)

		actual := list.ValueAt(pos)

		require.Nil(actual)
	})

	t.Run("Found Value", func(t *testing.T) {
		require := require.New(t)

		pos := 1

		list.Append(2)

		expected := 2
		actual := list.ValueAt(pos)

		require.Equal(expected, *actual)
	})
}
