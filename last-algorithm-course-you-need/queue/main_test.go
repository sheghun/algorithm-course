package main

import "testing"

func TestQueue(t *testing.T) {
	queue := NewQueue[int]()
	queue.enqueue(1)

	if queue.peek() == nil || *queue.peek() != 1 {
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
	if queue.length != 3 {
		t.Fatalf("expected queue.value.lengh to be %d, go %d", 3, queue.length)
	}

	if *queue.peek() != 1 {
		t.Fatalf("expected queue.value.head to be %d, go %d", 1, queue.head.Value)
	}

	queue.deque()
	if *queue.peek() != 2 {
		t.Fatalf("expected queue.value.head to be %d, go %d", 1, queue.head.Value)
	}

	queue.deque()
	if queue.head.Value != 3 {
		t.Fatalf("expected queue.value.head to be %d, go %d", 1, queue.head.Value)
	}
}
