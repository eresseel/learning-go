// main_test.go
package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSlices(t *testing.T) {
	s := Slices()

	require.NotNil(t, s, "The slice should not be nil")
	require.Equal(t, 0, len(s), "Expected empty slice")
	require.Equal(t, 0, cap(s), "Expected empty slice with capacity 0")
	require.Equal(t, []int{}, s)
}

func TestSlices2(t *testing.T) {
	s := Slices2()

	require.NotNil(t, s)
	require.Equal(t, 4, len(s), "Expected slice of 4 elements")
	require.GreaterOrEqual(t, cap(s), 4, "Capacity should be at least 4")
	require.Equal(t, []string{"Go", "Slices", "Are", "Powerful"}, s)
}
