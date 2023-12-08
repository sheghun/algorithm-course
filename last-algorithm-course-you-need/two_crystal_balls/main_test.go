package main

import "testing"

func TestTwoCrystalBalls(t *testing.T) {
	floors := []bool{false, false, false, true, true, true, true}
	val := twoCrystalBallsProblem(floors)
	if val != 3 {
		t.Fatalf("expected %d got %d", 3, val)
	}
}
