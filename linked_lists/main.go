package main

import "fmt"

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
	Prev  *Node[T]
}

type List[T comparable] struct {
	length int
	head   *Node[T]
	tail   *Node[T]
}

type LinkedList[T any] interface {
	Len() int
	InsertAt(item T, index int)
	Remove(item T) *T
	RemoveAt(index int) *T
	Append(item T)
	Prepend(item T)
	Get(index int) *T
}

func NewLinkedList[T comparable]() LinkedList[T] {
	var list List[T]

	return &list
}

func (l *List[T]) AllWithIndex(yield func(int, T) bool) {
	for h, i := l.head, 0; h != nil && i < l.Len(); h, i = h.Next, i+1 {
		if !yield(i, h.Value) {
			return
		}
	}
}

func (l *List[T]) AllWithNode(yield func(int, *Node[T]) bool) {
	for h, i := l.head, 0; h != nil && i < l.Len(); h, i = h.Next, i+1 {
		if !yield(i, h) {
			return
		}
	}
}

// Length implements LinkedList.
func (l *List[T]) Len() int {
	return l.length
}

// append implements LinkedList.
func (l *List[T]) Append(item T) {
	l.length++
	n := &Node[T]{Value: item, Next: nil, Prev: nil}
	if l.tail == nil {
		l.head = n
		l.tail = n
		return
	}
	n.Prev = l.tail
	l.tail.Next = n
	l.tail = n
}

// get implements LinkedList.
func (l *List[T]) Get(index int) *T {
	for i, val := range l.AllWithIndex {
		if i == index {
			return &val
		}
	}
	return nil
}

// insertAt implements LinkedList.
func (l *List[T]) InsertAt(item T, index int) {
	n := &Node[T]{Value: item, Next: nil, Prev: nil}
	for i, val := range l.AllWithNode {
		if i == index {
			l.length++
			a := val.Prev
			a.Next = n
			n.Next = val
			n.Prev = a
			val.Prev = n
		}
	}
}

// preprend implements LinkedList.
func (l *List[T]) Prepend(item T) {
	n := &Node[T]{Value: item, Next: nil, Prev: nil}
	l.length++
	l.head.Prev = n
	n.Next = l.head
	l.head = n
}

// remove implements LinkedList.
func (l *List[T]) Remove(item T) *T {
	for _, val := range l.AllWithNode {
		fmt.Println(val.Value)
		if item == val.Value {
			a := val.Prev
			c := val.Next
			if c != nil {
				c.Prev = a
			}
			a.Next = c
			l.length--
			return &val.Value
		}
	}
	return nil
}

// removeAt implements LinkedList.
func (l *List[T]) RemoveAt(index int) *T {
	for i, val := range l.AllWithNode {
		if i == index {
			a := val.Prev
			c := val.Next
			c.Prev = a
			a.Next = c

			l.length--
			return &val.Value
		}
	}
	return nil
}

func main() {
}
