package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	expected := []int{3, 4, 7, 9, 42, 69, 420}
	input := []int{9, 3, 7, 4, 69, 420, 42}
	quickSort(input)
	assert.Equal(t, input, expected, "The array's should be equal")
}
