package main

import (
	"errors"
	"fmt"
)

type Node[V any] struct {
	next  *Node[V]
	key   string
	value V
	prev  *Node[V]
}

type LRU[V any] struct {
	cap    int
	head   *Node[V]
	length int
	lookup map[string]*Node[V]
	tail   *Node[V]
}

func NewLRU[K, V any](cap int) *LRU[V] {
	return &LRU[V]{
		head:   nil,
		tail:   nil,
		length: 0,
		cap:    cap,
		lookup: make(map[string]*Node[V]),
	}
}

func (this *LRU[V]) Update(key string, value V) {
	node, ok := this.lookup[key]
	if ok {
		node.value = value
		this.detach(node)
		this.prepend(node)
		return
	}

	this.length++
	node = &Node[V]{value: value, key: key}
	this.lookup[key] = node
	this.prepend(node)
	this.trimCache()
}

func (this *LRU[V]) Get(key string) (V, error) {
	var v V
	node, ok := this.lookup[key]
	if !ok {
		return v, errors.New(fmt.Sprintf("Key %s does not exist", key))
	}

	return node.value, nil
}

func (this *LRU[V]) detach(node *Node[V]) {
	a, b := node.prev, node.next
	if a != nil {
		a.next = b
	}
	if b != nil {
		b.prev = a
	}

	if node == this.head {
		this.head = node.next
	}
	if node == this.tail {
		this.tail = node.prev
	}

	node.prev = nil
	node.next = nil
}

func (this *LRU[V]) prepend(node *Node[V]) {
	if this.head == nil {
		this.head = node
		this.tail = this.head
		return
	}

	if this.head == node {
		return
	}

	node.next, this.head.prev = this.head, node
	this.head = node
}

func (this *LRU[V]) trimCache() {
	if this.length <= this.cap {
		return
	}
	oldTail := this.tail
	this.detach(oldTail)
	delete(this.lookup, oldTail.key)
	this.length--
}

func main() {
}
