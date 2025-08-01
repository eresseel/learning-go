package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFormatValue(t *testing.T) {
	require.Equal(t, "Decimal: 42", FormatValue(42, "%d", "Decimal"))
	require.Equal(t, "Binary: 101010", FormatValue(42, "%b", "Binary"))
	require.Equal(t, "Hex: 2a", FormatValue(42, "%x", "Hex"))
	require.Equal(t, "Quoted: \"hello\"", FormatValue("hello", "%q", "Quoted"))
	require.Equal(t, "Type: string", FormatValue("hello", "%T", "Type"))
	require.Equal(t, "Float: 3.14", FormatValue(3.14159, "%.2f", "Float"))
}
