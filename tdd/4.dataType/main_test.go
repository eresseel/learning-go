package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDetectType(t *testing.T) {
	require.Equal(t, "int", DetectType(5))
	require.Equal(t, "string", DetectType("hello"))
	require.Equal(t, "float64", DetectType(3.14))
	require.Equal(t, "float32", DetectType(float32(3.14))) // <-- itt a float32 teszt
	require.Equal(t, "bool", DetectType(true))
	require.Equal(t, "slice", DetectType([]int{1, 2, 3}))
	require.Equal(t, "map", DetectType(map[string]int{"a": 1}))
}
