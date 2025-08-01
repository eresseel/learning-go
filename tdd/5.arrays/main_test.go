package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetArrays(t *testing.T) {
	expected1 := [3]int{1, 2, 3}
	expected2 := [5]int{4, 5, 6, 7, 8}

	got1, got2 := GetArrays()

	require.Equal(t, expected1, got1)
	require.Equal(t, expected2, got2)
}
