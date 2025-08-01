package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetCarInfo(t *testing.T) {
	car := GetCarInfo()
	require.Equal(t, "Ford", car["brand"])
	require.Equal(t, "Mustang", car["model"])
	require.Equal(t, "1964", car["year"])
}

func TestGetCityCodes(t *testing.T) {
	cities := GetCityCodes()
	require.Equal(t, 1, cities["Oslo"])
	require.Equal(t, 2, cities["Bergen"])
	require.Equal(t, 3, cities["Trondheim"])
	require.Equal(t, 4, cities["Stavanger"])
}
