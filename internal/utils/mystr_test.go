package mystr

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToLowerKebab(t *testing.T) {
	require.Equal(t, "hello", ToLowerKebab("hello"))
	require.Equal(t, "hello", ToLowerKebab("Hello"))
	require.Equal(t, "hello-world", ToLowerKebab("hello world"))
	require.Equal(t, "hello-world", ToLowerKebab("Hello World"))
	require.Equal(t, "hello-world-this-is-a-pretty-long-string", ToLowerKebab("Hello World this is a pretty long string"))
}