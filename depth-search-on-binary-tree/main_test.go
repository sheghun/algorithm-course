package main

import (
	"testing"
)

func TestDFSearch(t *testing.T) {
	// Test case 1: Needle is in the tree
	root1 := &BinaryNode[int]{Val: 10}
	root1.Left = &BinaryNode[int]{Val: 5}
	root1.Right = &BinaryNode[int]{Val: 15}
	root1.Left.Left = &BinaryNode[int]{Val: 2}
	root1.Left.Right = &BinaryNode[int]{Val: 7}
	root1.Right.Left = &BinaryNode[int]{Val: 12}
	root1.Right.Right = &BinaryNode[int]{Val: 20}

	if !DFSearch(root1, 7) {
		t.Errorf("Expected to find 7 in the tree")
	}

	// Test case 2: Needle is not in the tree
	if DFSearch(root1, 8) {
		t.Errorf("Did not expect to find 8 in the tree")
	}

	// Test case 3: Needle is the root
	if !DFSearch(root1, 10) {
		t.Errorf("Expected to find 10 in the tree")
	}

	// Test case 4: Needle is in the left subtree
	if !DFSearch(root1, 2) {
		t.Errorf("Expected to find 2 in the tree")
	}

	// Test case 5: Needle is in the right subtree
	if !DFSearch(root1, 20) {
		t.Errorf("Expected to find 20 in the tree")
	}

	// Test case 6: Empty tree
	var root2 *BinaryNode[int]
	if DFSearch(root2, 1) {
		t.Errorf("Did not expect to find 1 in an empty tree")
	}
}
