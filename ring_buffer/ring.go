// package main
//
// import (
// 	"errors"
// 	"sync"
// )
//
// type RingBugger struct {
// 	Size     int
// 	Byte     []byte
// 	Length   int
// 	writePos int
// 	readPos  int
// 	*sync.Mutex
// 	wait *sync.Cond
// }
//
// func NewRingBugger(size int) *RingBugger {
// 	b := make([]byte, size)
//
// 	return &RingBugger{
// 		Size:     size,
// 		Byte:     b,
// 		Length:   len(b),
// 		writePos: 1,
// 		readPos:  0,
// 	}
// }
//
// func (r *RingBugger) Write(b []byte) (int, error) {
// 	byteLength := len(b)
// 	endIndex := byteLength
//
// 	if byteLength > r.Length {
// 		return 0, errors.New("bytes too large")
// 	}
//
// 	// check if current byte input is larger than the remain empty slots
// 	if byteLength > r.Length-r.writePos {
// 		// rotate to the end starting index
// 		endIndex = (r.writePos + byteLength) - r.Length
// 		copy(r.Byte[r.writePos:], b[:r.Length-r.writePos])
// 		copy(r.Byte[:endIndex], b[r.Length-r.writePos:])
// 		r.writePos = endIndex
// 		return r.writePos, nil
// 	}
//
// 	// check if curren byte input is lower than the reamining empty slots
// 	copy(r.Byte[r.writePos:], b)
// 	r.writePos = endIndex
//
// 	return r.writePos, nil
// }
//
// func (r *RingBugger) Read(b []byte) (int, error) {
// 	byteLength := len(b)
// 	endIndex := byteLength
// 	if byteLength == 0 {
// 		return 0, nil
// 	}
//
// 	if r.writePos > r.readPos {
// 		endIndex = r.writePos - r.readPos
// 		if endIndex > byteLength {
// 			endIndex = byteLength
// 		}
// 		copy(b, r.Byte[r.readPos:endIndex])
// 		r.readPos = endIndex - 1
// 		return endIndex, nil
// 	}
//
// 	if byteLength > r.Length-r.readPos {
// 		endIndex = (r.readPos + byteLength) - r.Length
// 		if endIndex > r.writePos {
// 			endIndex = r.writePos
// 		}
// 		copy(b, r.Byte[r.readPos:])
// 		copy(b, r.Byte[:endIndex])
// 		r.readPos = endIndex
// 		return byteLength, nil
// 	}
//
// 	copy(b, r.Byte[r.readPos:])
// 	r.readPos = endIndex
// 	return endIndex, nil
// }
