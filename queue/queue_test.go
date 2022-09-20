package queue

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	require := require.New(t)

	actual := New[int]()

	require.NotNil(actual)
	require.NotNil(actual.l)
	require.Nil(actual.l.Value)
	require.Nil(actual.l.Next)
}

func TestPush(t *testing.T) {
	q := New[int]()

	t.Run("Check First Item", func(t *testing.T) {
		require := require.New(t)

		actual := q.l.Value

		require.Nil(actual)
	})

	t.Run("Push", func(t *testing.T) {
		require := require.New(t)

		expected := 10
		q.Push(expected)

		actual := *q.l.Value

		require.Equal(expected, actual)
	})
}

func TestGet(t *testing.T) {
	q := New[int]()

	t.Run("Get Empty Queue", func(t *testing.T) {
		require := require.New(t)

		actual := q.Get()

		require.Nil(actual)
	})

	t.Run("Get w/ 1 value", func(t *testing.T) {
		require := require.New(t)

		expectedValue := 10
		q.Push(expectedValue)

		actualValue := q.Get()

		require.NotNil(actualValue)
		require.Equal(expectedValue, *actualValue)
	})
}

func TestPop(t *testing.T) {
	q := New[int]()

	t.Run("Pop Empty Queue", func(t *testing.T) {
		require := require.New(t)

		expectedValue := 0
		actualValue, actualErr := q.Pop()

		require.Equal(expectedValue, actualValue)
		require.NotNil(actualErr)
	})

	t.Run("Pop w/ 1 value", func(t *testing.T) {
		require := require.New(t)

		expectedValue := 10
		q.Push(expectedValue)

		actualValue, actualErr := q.Pop()

		require.Equal(expectedValue, actualValue)
		require.Nil(actualErr)
	})
}
