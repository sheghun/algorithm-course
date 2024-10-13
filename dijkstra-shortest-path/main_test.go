package main

import (
	"testing"
)

func TestDijkstraList(t *testing.T) {
	graph := WeightedAdjacencyList{
		{{To: 1, Weight: 4}, {To: 2, Weight: 1}}, // Node 0
		{{To: 3, Weight: 1}},                     // Node 1
		{{To: 1, Weight: 2}, {To: 3, Weight: 5}}, // Node 2
		{},                                       // Node 3
	}

	source := 0
	sink := 3
	expectedDistance := 4 // Expected shortest path distance from node 0 to node 3

	// Run Dijkstra's algorithm
	dists := DjikstraList(source, sink, graph)

	// Check the distance to the sink node
	if dists[sink] != expectedDistance {
		t.Errorf(
			"Expected distance to node %d is %d, but got %d",
			sink,
			expectedDistance,
			dists[sink],
		)
	}
}
