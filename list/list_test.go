package list

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewList(t *testing.T) {
	require := require.New(t)

	expected := List[int]{}
	actual := New[int]()

	require.Equal(expected, *actual)
}
