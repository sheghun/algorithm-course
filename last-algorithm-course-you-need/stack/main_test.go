package main

import "testing"

func TestStack(t *testing.T) {

	stack := NewStack[int]()
	t.Run("test new stack", func(t *testing.T) {
		if stack.Len() != 0 || stack.Head != nil {
			t.Errorf("expected 0 and nil got %d and %v", stack.Len(), stack.Head)
		}
	})

	t.Run("test push", func(t *testing.T) {
		stack.Push(1)
		if stack.Len() != 1 || stack.Head == nil {
			t.Errorf("pushing failed ")
		}
	})

	t.Run("test pop", func(t *testing.T) {
		val := stack.Pop()
		if *val != 1 && stack.Len() != 0 {
			t.Errorf("error occured calling pop")
		}

		val = stack.Pop()
		if val != nil && stack.Len() != 0 {
			t.Errorf("calling pop the second time failed")
		}
	})
}
