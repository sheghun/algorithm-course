package main

import (
	"testing"
)

func TestLRU(t *testing.T) {
	tests := []struct {
		name       string
		operations []func(lru *LRU[int]) (int, error)
		expected   []int
	}{
		{
			name: "Basic operations",
			operations: []func(lru *LRU[int]) (int, error){
				func(lru *LRU[int]) (int, error) {
					lru.Update("a", 1)
					return 0, nil
				},
				func(lru *LRU[int]) (int, error) {
					return lru.Get("a")
				},
				func(lru *LRU[int]) (int, error) {
					lru.Update("b", 2)
					return 0, nil
				},
				func(lru *LRU[int]) (int, error) {
					return lru.Get("b")
				},
			},
			expected: []int{0, 1, 0, 2},
		},
		{
			name: "Eviction test",
			operations: []func(lru *LRU[int]) (int, error){
				func(lru *LRU[int]) (int, error) {
					lru.Update("a", 1)
					return 0, nil
				},
				func(lru *LRU[int]) (int, error) {
					lru.Update("b", 2)
					return 0, nil
				},
				func(lru *LRU[int]) (int, error) {
					lru.Update("c", 3)
					return 0, nil
				},
				func(lru *LRU[int]) (int, error) {
					_, err := lru.Get("a")
					if err != nil {
						return -1, err
					}
					return 0, nil
				},
				func(lru *LRU[int]) (int, error) {
					return lru.Get("b")
				},
				func(lru *LRU[int]) (int, error) {
					return lru.Get("c")
				},
			},
			expected: []int{0, 0, 0, -1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lru := NewLRU[string, int](2)
			for i, op := range tt.operations {
				result, err := op(lru)
				if err != nil && result != -1 {
					t.Errorf("unexpected error: %v", err)
				}
				if result != tt.expected[i] {
					t.Errorf("expected %d, got %d", tt.expected[i], result)
				}
			}
		})
	}
}
