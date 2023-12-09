package main

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type Queue[T any] struct {
	head   *Node[T]
	tail   *Node[T]
	length int
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

func (q *Queue[T]) enqueue(item T) {
	q.length++
	node := &Node[T]{Value: item, Next: nil}
	if q.tail == nil {
		q.head = node
		q.tail = q.head
	}
	q.tail.Next = node
	q.tail = node
}

func (q *Queue[T]) deque() *T {
	if q.head == nil {
		return nil
	}
	q.length--
	head := q.head
	q.head = q.head.Next
	return &head.Value
}

func main() {

}
