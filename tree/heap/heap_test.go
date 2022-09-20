package heap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	require := require.New(t)

	h := New[int]()

	require.NotNil(h)
	require.Len(h.nodes, 0)
}
