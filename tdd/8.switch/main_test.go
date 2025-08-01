package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCategorizeDay(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{1, "Odd weekday"},
		{3, "Odd weekday"},
		{5, "Odd weekday"},
		{2, "Even weekday"},
		{4, "Even weekday"},
		{6, "Weekend"},
		{7, "Weekend"},
		{0, "Invalid day of day number"},
		{8, "Invalid day of day number"},
		{-1, "Invalid day of day number"},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, CategorizeDay(tt.input))
	}
}

func TestGetDayName(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{1, "Monday"},
		{2, "Tuesday"},
		{3, "Wednesday"},
		{4, "Thursday"},
		{5, "Friday"},
		{6, "Saturday"},
		{7, "Sunday"},
		{0, "Invalid day"},
		{8, "Invalid day"},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, GetDayName(tt.input))
	}
}
