package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(binarySearch([]int{-3, -2, 1, 2, 3, 4, 5, 6, 7, 8, 9}, -2))
}

func binarySearch(arr []int, val int) int {
	lo := 0
	h := len(arr)
	m := 0

	for lo < h {
		tmp := float64(lo + (h-lo)/2)
		m = int(math.Floor(tmp))
		if arr[m] == val {
			return m
		}
		if arr[m] > val {
			h = m
		} else {
			lo = m + 1
		}
	}
	return -1
}
