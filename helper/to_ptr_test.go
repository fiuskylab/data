package helper

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToPtr(t *testing.T) {
	require := require.New(t)

	got := ToPtr(int(1))

	typeOf := reflect.TypeOf(got)

	expected := reflect.Pointer
	actual := typeOf.Kind()

	require.Equal(expected, actual)
}
