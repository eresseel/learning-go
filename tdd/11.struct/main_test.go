package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewPerson(t *testing.T) {
	p := NewPerson("Hege", 45, "Teacher", 6000)

	require.Equal(t, "Hege", p.Name)
	require.Equal(t, 45, p.Age)
	require.Equal(t, "Teacher", p.Job)
	require.Equal(t, 6000, p.Salary)
}
