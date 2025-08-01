// main_test.go
package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTestCount(t *testing.T) {
	got := TestCount(1)
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	require.Equal(t, expected, got)
}
