package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSumWithForLoop(t *testing.T) {
	require.Equal(t, 15, SumWithForLoop(5)) // 1+2+3+4+5 = 15
	require.Equal(t, 0, SumWithForLoop(0))
	require.Equal(t, 1, SumWithForLoop(1))
}

func TestSumWithWhileStyle(t *testing.T) {
	require.Equal(t, 10, SumWithWhileStyle(4)) // 1+2+3+4 = 10
	require.Equal(t, 0, SumWithWhileStyle(0))
	require.Equal(t, 21, SumWithWhileStyle(6)) // 1+2+3+4+5+6 = 21
}

func TestSumSliceValues(t *testing.T) {
	require.Equal(t, 6, SumSliceValues([]int{1, 2, 3}))
	require.Equal(t, 0, SumSliceValues([]int{}))
	require.Equal(t, -1, SumSliceValues([]int{-1, 0, 0}))
}
