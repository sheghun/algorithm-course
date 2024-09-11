package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

// Helper function to capture the output of BreadthFirstSearch
func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	f()
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func TestBreadthFirstSearch(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected string
	}{
		{
			name:     "Empty tree",
			root:     nil,
			expected: "",
		},
		{
			name:     "Single node",
			root:     &TreeNode{Value: 1},
			expected: "1\n",
		},
		{
			name: "Complete binary tree",
			root: &TreeNode{
				Value: 1,
				Left: &TreeNode{
					Value: 2,
					Left:  &TreeNode{Value: 4},
					Right: &TreeNode{Value: 5},
				},
				Right: &TreeNode{
					Value: 3,
					Left:  &TreeNode{Value: 6},
					Right: &TreeNode{Value: 7},
				},
			},
			expected: "1\n2\n3\n4\n5\n6\n7\n",
		},
		{
			name: "Left-skewed tree",
			root: &TreeNode{
				Value: 1,
				Left: &TreeNode{
					Value: 2,
					Left: &TreeNode{
						Value: 3,
						Left: &TreeNode{
							Value: 4,
						},
					},
				},
			},
			expected: "1\n2\n3\n4\n",
		},
		{
			name: "Right-skewed tree",
			root: &TreeNode{
				Value: 1,
				Right: &TreeNode{
					Value: 2,
					Right: &TreeNode{
						Value: 3,
						Right: &TreeNode{
							Value: 4,
						},
					},
				},
			},
			expected: "1\n2\n3\n4\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := captureOutput(func() {
				BreadthFirstSearch(tt.root)
			})
			if output != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, output)
			}
		})
	}
}
