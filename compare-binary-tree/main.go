package main

type BinaryNode[T comparable] struct {
	Left  *BinaryNode[T]
	Right *BinaryNode[T]
	Value T
}

func CompareBinaryTree[T comparable](a *BinaryNode[T], b *BinaryNode[T]) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if a.Value != b.Value {
		return false
	}

	return CompareBinaryTree(a.Left, b.Left) && CompareBinaryTree(a.Right, b.Right)
}

func main() {
}
