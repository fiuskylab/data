package list

import (
	"fmt"
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

func TestLastInstance(t *testing.T) {
	list := New[int]()

	t.Run("Empty List", func(t *testing.T) {
		require := require.New(t)

		actual := list.lastInstance()

		require.Nil(actual.Next)
		require.Nil(actual.Value)
	})

	t.Run("List w/ 1 item", func(t *testing.T) {
		require := require.New(t)

		expected := List[int]{
			Value: helper.ToPtr(1),
		}

		list.Append(1)

		actual := list.lastInstance()

		require.Nil(actual.Next)
		require.Equal(*expected.Value, *actual.Value)
	})

	for i := 2; i <= 10; i++ {
		list.Append(i)
	}

	t.Run("List w/ 10 item", func(t *testing.T) {
		require := require.New(t)

		expected := List[int]{
			Value: helper.ToPtr(10),
		}

		actual := list.lastInstance()

		require.Nil(actual.Next)
		require.Equal(*expected.Value, *actual.Value)
	})
}

func TestValueAt(t *testing.T) {
	list := New[int]()
	t.Run("Empty List", func(t *testing.T) {
		require := require.New(t)

		expectedVal := 0
		expectedErr := fmt.Errorf(errEmptyList)
		actualVal, actualErr := list.ValueAt(100)

		require.Equal(expectedVal, actualVal)
		require.Equal(expectedErr, actualErr)
	})

	t.Run("Shorter Than", func(t *testing.T) {
		require := require.New(t)

		pos := 100

		list.Append(1)

		expectedVal := 0
		expectedErr := fmt.Errorf(errShorterThan, pos)
		actualVal, actualErr := list.ValueAt(pos)

		require.Equal(expectedVal, actualVal)
		require.Equal(expectedErr, actualErr)
	})

	t.Run("Found Value", func(t *testing.T) {
		require := require.New(t)

		pos := 1

		list.Append(2)

		expectedVal := 2
		actualVal, actualErr := list.ValueAt(pos)

		require.Equal(expectedVal, actualVal)
		require.Nil(actualErr)
	})
}
