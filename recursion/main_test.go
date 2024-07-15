package main

import (
	"reflect"
	"testing"
)

func TestMaze(t *testing.T) {
	maze := []string{
		"xxxxxxxxxx x",
		"x        x x",
		"x        x x",
		"x xxxxxxxx x",
		"x          x",
		"x xxxxxxxxxx",
	}

	expectedMazeResult := []Point{
		{x: 10, y: 0},
		{x: 10, y: 1},
		{x: 10, y: 2},
		{x: 10, y: 3},
		{x: 10, y: 4},
		{x: 9, y: 4},
		{x: 8, y: 4},
		{x: 7, y: 4},
		{x: 6, y: 4},
		{x: 5, y: 4},
		{x: 4, y: 4},
		{x: 3, y: 4},
		{x: 2, y: 4},
		{x: 1, y: 4},
		{x: 1, y: 5},
	}

	actualMazeResult := mazeSolver(maze, 'x', Point{x: 10, y: 0}, Point{x: 1, y: 5})

	if !reflect.DeepEqual(actualMazeResult, expectedMazeResult) {
		t.Errorf("expected: %#v, got: %#v", expectedMazeResult, actualMazeResult)
	}
}
