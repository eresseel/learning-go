package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValuePI(t *testing.T) {
	got := ValuePI()
	require.Equal(t, float32(3.14), got)
}
