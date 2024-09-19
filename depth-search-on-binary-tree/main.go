package main

type BinaryNode[T comparable] struct {
	Left  *BinaryNode[T]
	Right *BinaryNode[T]
	Val   T
}

func search(curr *BinaryNode[int], needle int) bool {
	if curr == nil {
		return false
	}

	if curr.Val == needle {
		return true
	}

	if curr.Val < needle {
		return search(curr.Right, needle)
	}

	return search(curr.Left, needle)
}

func DFSearch(head *BinaryNode[int], needle int) bool {
	return search(head, needle)
}

func main() {
}
