package main

import (
	"errors"
	"fmt"
	"log"
	"sync"
)

type RingBugger struct {
	Size     int
	Byte     []byte
	Length   int
	writePos int
	readPos  int
	*sync.Mutex
	wait *sync.Cond
}

func NewRingBugger(size int) *RingBugger {
	b := make([]byte, size)

	return &RingBugger{
		Size:     size,
		Byte:     b,
		Length:   len(b),
		writePos: 0,
		readPos:  0,
	}
}

func (r *RingBugger) Write(b []byte) (int, error) {
	byteLength := len(b)
	endIndex := byteLength

	if byteLength > r.Length {
		return 0, errors.New("bytes too large")
	}

	// check if current byte input is larger than the remain empty slots
	if byteLength > r.Length-r.writePos {
		// rotate to the end starting index
		endIndex = (r.writePos + byteLength) - r.Length
		copy(r.Byte[r.writePos:], b[:r.Length-r.writePos])
		copy(r.Byte[:endIndex], b[r.Length-r.writePos:])
		r.writePos = endIndex
		return r.writePos, nil
	}

	// check if curren byte input is lower than the reamining empty slots
	copy(r.Byte[r.writePos:], b)
	r.writePos += endIndex

	return r.writePos, nil
}

func (r *RingBugger) Read(b []byte) (int, error) {
	byteLength := len(b)
	endIndex := byteLength
	if byteLength == 0 {
		return 0, nil
	}

	if r.writePos > r.readPos {
		endIndex = r.writePos - r.readPos
		fmt.Println(endIndex)
		if endIndex > byteLength {
			endIndex = byteLength
		}
		copy(b, r.Byte[r.readPos:endIndex])
		r.readPos = endIndex
		return endIndex, nil
	}

	if byteLength > r.Length-r.readPos {
		endIndex = (r.readPos + byteLength) - r.Length
		if endIndex > r.writePos {
			endIndex = r.writePos
		}
		copy(b, r.Byte[r.readPos:])
		copy(b, r.Byte[:endIndex])
		totalRead := r.readPos + 1
		r.readPos = endIndex
		return totalRead, nil
	}

	copy(b, r.Byte[r.readPos:])
	r.readPos = endIndex
	return endIndex, nil
}

func main() {
	// fmt.Println("here")
	// users := []int{1, 2, 3, 4, 5}
	// usersAddon := []int{6, 7, 8, 9, 10}
	// copy(users[1:2], usersAddon[:5])
	//  copy(users[2:3], usersAddon[1:5])
	// fmt.Println(users)

	ringBuff := NewRingBugger(19)
	_, err := ringBuff.Write([]byte("Segun is a good boy"))
	if err != nil {
		log.Fatal(err)
	}

	readBuffer := make([]byte, 2)
	totalRead, err := ringBuff.Read(readBuffer)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("READ BUFFER: %s\n", readBuffer)
	fmt.Printf("TOAL READ: %d\n", totalRead)
	fmt.Printf("READ POSITION: %d\n", ringBuff.readPos)
	fmt.Printf("WRITE POSITION: %d\n", ringBuff.writePos)

	_, err = ringBuff.Write([]byte("Where is"))
	if err != nil {
		log.Fatal(err)
	}

	totalRead, err = ringBuff.Read(readBuffer)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("READ BUFFER: %s\n", readBuffer)
	fmt.Printf("TOAL READ: %d\n", totalRead)
	fmt.Printf("READ POSITION: %d\n", ringBuff.readPos)
	fmt.Printf("WRITE POSITION: %d\n", ringBuff.writePos)
}
