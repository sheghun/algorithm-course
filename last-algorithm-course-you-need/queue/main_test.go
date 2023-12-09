package main

import "testing"

func TestQueue(t *testing.T) {
	queue := NewQueue[int]()
	queue.enqueue(1)

	if queue.head.Value != 1 {
		t.Fatalf("expected queue.value.head to be %d, go %d", 1, queue.head.Value)
	}

	queue.enqueue(2)
	if queue.tail.Value != 2 {
		t.Fatalf("expected queue.value.tail to be %d, go %d", 1, queue.head.Value)
	}

	queue.enqueue(3)
	if queue.tail.Value != 3 {
		t.Fatalf("expected queue.value.tail to be %d, go %d", 1, queue.head.Value)
	}

	if queue.head.Value != 1 {
		t.Fatalf("expected queue.value.head to be %d, go %d", 1, queue.head.Value)
	}

	queue.deque()
	if queue.head.Value != 2 {
		t.Fatalf("expected queue.value.head to be %d, go %d", 1, queue.head.Value)
	}

	queue.deque()
	if queue.head.Value != 3 {
		t.Fatalf("expected queue.value.head to be %d, go %d", 1, queue.head.Value)
	}

}
