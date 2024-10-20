package main

type Node[V any] struct {
	next *Node[V]
	prev *Node[V]
}

type LRU[K, V any] struct {
	head   *Node[V]
	tail   *Node[V]
	length int
	lookup map[any]V
}

func NewLRU(length) {
  return &
}

func main() {
}
