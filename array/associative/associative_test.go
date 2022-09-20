package associative

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	require := require.New(t)

	m := NewArray[int, string]()

	require.NotNil(m)
	require.NotNil(m.items)
}

func TestPut(t *testing.T) {
	require := require.New(t)
	m := NewArray[int, string]()

	expectedValue := "um"
	m.Put(1, expectedValue)

	require.NotNil(m.items[0].Value)
	require.Equal(*m.items[0].Value, expectedValue)
}

func TestGet(t *testing.T) {
	m := NewArray[int, string]()

	t.Run("Empty Array", func(t *testing.T) {
		require := require.New(t)

		require.Nil(m.Get(1))
	})

	t.Run("With item", func(t *testing.T) {
		m.Put(1, "um")
		t.Run("Wrong key", func(t *testing.T) {
			require := require.New(t)

			require.Nil(m.Get(2))
		})

		t.Run("Found", func(t *testing.T) {
			require := require.New(t)

			require.NotNil(m.Get(1))
		})
	})
}

func TestDelete(t *testing.T) {
	m := NewArray[int, string]()

	m.Put(1, "um")

	t.Run("Wrong Key", func(t *testing.T) {
		require := require.New(t)

		require.False(m.Delete(20))
	})

	t.Run("Found", func(t *testing.T) {
		require := require.New(t)

		require.True(m.Delete(1))
	})
}
