package main

import "math"

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

type Stack[T comparable] struct {
	length int
	Head   *Node[T]
}

func NewStack[T comparable]() *Stack[T] {
	return &Stack[T]{
		length: 0,
		Head:   nil,
	}
}

func (s *Stack[T]) Len() int {
	return s.length
}

func (s *Stack[T]) Push(item T) {
	s.length++
	node := &Node[T]{Value: item}
	if s.Head == nil {
		s.Head = node
		return
	}

	node.Next = s.Head
	s.Head = node
}

func (s *Stack[T]) Pop() *T {
	if s.Head == nil {
		return nil
	}

	if s.length == 0 {
		return nil
	}

	head := *s.Head
	s.Head = s.Head.Next

	s.length = int(math.Max(0, float64(s.length)-1))

	return &head.Value
}

func (s *Stack[T]) Peek(item T) T {
	return s.Head.Value
}

func main() {
}
