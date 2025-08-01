package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCheckNumber(t *testing.T) {
	t.Run("num < 10", func(t *testing.T) {
		got := CheckNumber(5)
		require.Equal(t, []string{"Num is less than 10"}, got)
	})

	t.Run("10 <= num <= 15", func(t *testing.T) {
		got := CheckNumber(12)
		require.Equal(t, []string{"Num is more than 10"}, got)
	})

	t.Run("num > 15", func(t *testing.T) {
		got := CheckNumber(20)
		require.Equal(t, []string{
			"Num is more than 10",
			"Num is also more than 15",
		}, got)
	})
}
