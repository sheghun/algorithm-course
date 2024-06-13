package main

import (
	"errors"
	"fmt"
)

type CircularBuffer[T any] struct {
	data []T
	size int
	w    int
	r    int
}

func NewCircularBuffer[T any](size int) *CircularBuffer[T] {
	d := make([]T, size)

	return &CircularBuffer[T]{
		data: d,
		size: size,
		w:    0,
		r:    0,
	}
}

func (c *CircularBuffer[T]) Write(data []T) (int, error) {
	s := len(data)

	if s > c.size {
		return 0, errors.New("data size too large")
	}

	e := c.w + s - c.size // no end of slice
	if e <= 0 {
		copy(c.data[c.w:], data)
		c.w += s
		return c.w, nil
	} // remainder of the slice
	copy(c.data[c.w:], data[:c.size-c.w])
	copy(c.data[:e], data[c.size-c.w:])
	c.w = e
	return c.w, nil
}

func (c *CircularBuffer[T]) Read(data []T) error {
	s := len(data)

	if c.r < c.w {
		if c.r+s > c.w {
			copy(data, c.data[c.r:c.w])
			c.r = c.w
			return nil
		}
		copy(data, c.data[c.r:c.r+s])
		c.r += s
		return nil
	}
	e := c.r + s - c.size
	if e <= 0 {
		copy(data, c.data[c.r:])
		c.r += s
		return nil
	}
	copy(data, c.data[c.r:])
	if e > c.w {
		e = c.w
	}
	copy(data[c.size-c.r:], c.data[:e])
	c.r = e
	return nil
}

func main() {
	c := NewCircularBuffer[int](10)
	c.Write([]int{9, 10, 2, 3, 2, 4, 5, 2, 1, 3})
	fmt.Println("WRITE INDEX: ", c.w)
	fmt.Println("READ INDEX: ", c.r)
	read := make([]int, 2)
	c.Read(read)
	fmt.Println(read)
	fmt.Println("READ INDEX: ", c.r)
	fmt.Println("----------------------------------------------------------")

	c.Write([]int{11, 12})
	fmt.Println("WRITE INDEX: ", c.w)
	read2 := make([]int, 10)
	c.Read(read2)
	fmt.Println(read2)
	fmt.Println("READ INDEX: ", c.r)
	fmt.Println("CIRCULAR DATA: ", c.data)
}
