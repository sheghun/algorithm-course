package main

import (
	"sort"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{2, 3, 4, 5, 10, 9, 22, 11, 23, 43}
	arr = BubbleSort(arr)
	if !sort.IntsAreSorted(arr) {
		t.Fatalf("Bubble sort failed")
	}
}
