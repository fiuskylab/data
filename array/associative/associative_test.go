package associative

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAssociativeArray(t *testing.T) {
	require := require.New(t)

	m := NewArray[int, string]()

	m.Put(1, "um")
	m.Put(2, "dois")
	m.Put(3, "tres")

	require.False(m.Delete(120))
	require.True(m.Delete(3))

	require.Equal("dois", *m.Get(2))
	require.Equal("um", *m.Get(1))
}
