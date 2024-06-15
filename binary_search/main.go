package main

import (
	"fmt"
)

func main() {
	fmt.Println(binarySearch([]int{-3, -2, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 12))
}

func binarySearch(arr []int, needle int) int {
	lo := 0
	hi := len(arr)

	for lo < hi {
		m := int(float64(lo + (hi-lo)/2))
		val := arr[m]

		if val == needle {
			return m
		}

		if val > needle {
			hi = m
		} else {
			lo = m + 1
		}
	}
	return -1
}
