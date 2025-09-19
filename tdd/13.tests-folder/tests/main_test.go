package tests

import (
	"testing"

	"example.com/hellotdd"
	"github.com/stretchr/testify/require"
)

func TestHello(t *testing.T) {
	got := hellotdd.Hello("TDD")
	want := "Hello, TDD!"
	require.Equal(t, want, got)
}
