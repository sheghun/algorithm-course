package main

import (
	"fmt"
)

func qs(arr []int, lo, hi int) {
	if lo >= hi {
		return
	}
	pivotIdx := partition(arr, lo, hi)

	qs(arr, lo, pivotIdx-1)
	qs(arr, pivotIdx+1, hi)
}

func partition(arr []int, lo, hi int) int {
	pivot := arr[hi]
	idx := lo - 1
	for i := lo; i < hi; i++ {
		if arr[i] <= pivot {
			idx++
			arr[idx], arr[i] = arr[i], arr[idx]
		}
	}
	idx++
	arr[idx], arr[hi] = arr[hi], arr[idx]
	return idx
}

func quickSort(arr []int) {
	qs(arr, 0, len(arr)-1)
}

func main() {
	arr := []int{9, 3, 7, 4, 69, 420, 42}
	quickSort(arr)
	fmt.Println(arr)
}
