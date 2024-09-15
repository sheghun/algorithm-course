package main

import (
	"testing"
)

func TestCompareBinaryTree_EqualTrees(t *testing.T) {
	tree1 := &BinaryNode[int]{
		Value: 1,
		Left:  &BinaryNode[int]{Value: 2},
		Right: &BinaryNode[int]{Value: 3},
	}
	tree2 := &BinaryNode[int]{
		Value: 1,
		Left:  &BinaryNode[int]{Value: 2},
		Right: &BinaryNode[int]{Value: 3},
	}

	if !CompareBinaryTree(tree1, tree2) {
		t.Errorf("Expected trees to be equal, but they are not")
	}
}

func TestCompareBinaryTree_UnequalTrees(t *testing.T) {
	tree1 := &BinaryNode[int]{
		Value: 1,
		Left:  &BinaryNode[int]{Value: 2},
		Right: &BinaryNode[int]{Value: 3},
	}
	tree2 := &BinaryNode[int]{
		Value: 1,
		Left:  &BinaryNode[int]{Value: 4},
		Right: &BinaryNode[int]{Value: 3},
	}

	if CompareBinaryTree(tree1, tree2) {
		t.Errorf("Expected trees to be unequal, but they are equal")
	}
}

func TestCompareBinaryTree_OneTreeIsNil(t *testing.T) {
	tree1 := &BinaryNode[int]{
		Value: 1,
		Left:  &BinaryNode[int]{Value: 2},
		Right: &BinaryNode[int]{Value: 3},
	}
	tree2 := (*BinaryNode[int])(nil)

	if CompareBinaryTree(tree1, tree2) {
		t.Errorf("Expected trees to be unequal when one tree is nil, but they are equal")
	}
}

func TestCompareBinaryTree_BothTreesNil(t *testing.T) {
	var tree1 *BinaryNode[string]
	var tree2 *BinaryNode[string]

	if !CompareBinaryTree(tree1, tree2) {
		t.Errorf("Expected trees to be equal when both trees are nil, but they are not")
	}
}

func TestCompareBinaryTree_DifferentStructures(t *testing.T) {
	tree1 := &BinaryNode[int]{Value: 1, Left: &BinaryNode[int]{Value: 2}}
	tree2 := &BinaryNode[int]{Value: 1, Right: &BinaryNode[int]{Value: 2}}

	if CompareBinaryTree(tree1, tree2) {
		t.Errorf("Expected trees to be unequal due to different structures, but they are equal")
	}
}
