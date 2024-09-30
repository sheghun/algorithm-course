package main

import (
	"reflect"
	"testing"
)

func TestDFSGraphList(t *testing.T) {
	got := DFSGraphList(graph2, 0, 6)
	expected := []int{0, 1, 4, 5, 6}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected: %v got: %v", expected, got)
	}
}
