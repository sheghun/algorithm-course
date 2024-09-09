package main

import "log"

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) walk(curr *TreeNode, path *[]int, order string) {
	// pre
	if curr == nil {
		return
	}

	switch order {
	case "pre":
		*path = append(*path, curr.Value)
		t.walk(curr.Left, path, "pre")
		t.walk(curr.Right, path, "pre")
		return

	case "in":
		t.walk(curr.Left, path, "in")
		*path = append(*path, curr.Value)
		t.walk(curr.Right, path, "in")
		return

	case "post":
		t.walk(curr.Left, path, "post")
		t.walk(curr.Right, path, "post")
		*path = append(*path, curr.Value)
		return

	default:
		log.Fatal("order not specified")
	}
}

func PreOrderSearch(head *TreeNode) []int {
	path := []int{}

	if head == nil {
		return path
	}

	head.walk(head, &path, "pre")
	return path
}

func InOrderSearch(head *TreeNode) []int {
	path := []int{}

	if head == nil {
		return path
	}

	head.walk(head, &path, "in")
	return path
}

func PostOrderSearch(head *TreeNode) []int {
	path := []int{}

	if head == nil {
		return path
	}

	head.walk(head, &path, "post")
	return path
}

func main() {
}
