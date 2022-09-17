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

		require.Equal(expected, *list.value)
	})

	t.Run("Non Empty List", func(t *testing.T) {
		require := require.New(t)
		expected := 2
		list.Append(expected)

		require.Equal(expected, *list.next.value)
	})
}

func TestLastInstance(t *testing.T) {
	list := New[int]()

	t.Run("Empty List", func(t *testing.T) {
		require := require.New(t)

		actual := list.lastInstance()

		require.Nil(actual.next)
		require.Nil(actual.value)
	})

	t.Run("List w/ 1 item", func(t *testing.T) {
		require := require.New(t)

		expected := List[int]{
			value: helper.ToPtr(1),
		}

		list.Append(1)

		actual := list.lastInstance()

		require.Nil(actual.next)
		require.Equal(*expected.value, *actual.value)
	})

	for i := 2; i <= 10; i++ {
		list.Append(i)
	}

	t.Run("List w/ 10 item", func(t *testing.T) {
		require := require.New(t)

		expected := List[int]{
			value: helper.ToPtr(10),
		}

		actual := list.lastInstance()

		require.Nil(actual.next)
		require.Equal(*expected.value, *actual.value)
	})
}
