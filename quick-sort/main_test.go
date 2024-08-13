package main

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	expected := []int{3, 4, 7, 9, 42, 69, 420}
	input := []int{9, 3, 7, 4, 69, 420, 42}
	quickSort(input)
	if !reflect.DeepEqual(input, expected) {
		t.Errorf("expected: %v\n got: %v", expected, input)
	}
}
