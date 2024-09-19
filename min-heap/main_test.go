package main

import (
	"testing"
)

func TestMinHeap_Insert(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{3, 1, 6, 5, 2, 4}, expected: []int{1, 2, 4, 5, 3, 6}},
		{input: []int{10, 20, 15, 30, 40}, expected: []int{10, 20, 15, 30, 40}},
		{input: []int{5, 3, 8, 1, 2}, expected: []int{1, 2, 8, 5, 3}},
	}

	for _, test := range tests {
		heap := &MinHeap{}
		for _, val := range test.input {
			heap.Insert(val)
		}

		for i, val := range test.expected {
			if heap.data[i] != val {
				t.Errorf("expected %d at index %d, got %d", val, i, heap.data[i])
			}
		}
	}
}

func TestMinHeap_Delete(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{3, 1, 6, 5, 2, 4}, expected: []int{1, 2, 3, 4, 5, 6}},
		{input: []int{10, 20, 15, 30, 40}, expected: []int{10, 15, 20, 30, 40}},
		{input: []int{5, 3, 8, 1, 2}, expected: []int{1, 2, 3, 5, 8}},
	}

	for _, test := range tests {
		heap := &MinHeap{}
		for _, val := range test.input {
			heap.Insert(val)
		}

		for _, expectedVal := range test.expected {
			if val := heap.Delete(); val != expectedVal {
				t.Errorf("expected %d, got %d", expectedVal, val)
			}
		}
	}
}

func TestMinHeap_EmptyDelete(t *testing.T) {
	heap := &MinHeap{}
	if val := heap.Delete(); val != -1 {
		t.Errorf("expected -1, got %d", val)
	}
}
