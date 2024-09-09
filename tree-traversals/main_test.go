package main

import (
	"reflect"
	"testing"
)

func TestTreeTraversals(t *testing.T) {
	// Construct the binary tree:
	//        1
	//       / \
	//      2   3
	//     / \
	//    4   5
	root := &TreeNode{
		1,
		&TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}},
		&TreeNode{3, nil, nil},
	}

	tests := []struct {
		name      string
		traversal func(*TreeNode) []int
		want      []int
	}{
		{"Preorder", PreOrderSearch, []int{1, 2, 4, 5, 3}},
		{"Inorder", InOrderSearch, []int{4, 2, 5, 1, 3}},
		{"Postorder", PostOrderSearch, []int{4, 5, 2, 3, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.traversal(root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%s() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
