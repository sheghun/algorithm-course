package main

import "fmt"

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Value int
}

func BreadthFirstSearch(head *TreeNode) {
	queue := []*TreeNode{}
	if head == nil {
		return
	}
	queue = append(queue, head)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		// Process the current node (e.g., print its value)
		fmt.Println(node.Value)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
}

func main() {
}
