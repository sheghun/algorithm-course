package main

import "math"

type node struct {
	To     int
	Weight int
}

type WeightedAdjacencyList [][]node

func DjikstraList(source, sink int, graph WeightedAdjacencyList) {
	seen := make([]bool, len(graph))
	prev := make([]int, len(graph))
	for i := range prev {
		prev[i] = -1
	}
	dists := make([]int, len(graph))
	for i := range dists {
		dists[i] = math.MaxInt64
	}
	dists[source] = 0

	for hasUnvisited(seen, dists) {
		curr := getLowestUnvisited(seen, dists)
		seen[curr] = true
		nodes := graph[curr]
		for _, node := range nodes {
			if seen[node.To] {
				continue
			}

			dist := dists[curr] + node.Weight
			if dist < dists[node.To] {
				dists[node.To] = dist
				prev[node.To] = curr
			}
		}
	}
}

func hasUnvisited(seen []bool, dists []int) bool {
	for i := range seen {
		if !seen[i] && dists[i] < math.MaxInt64 {
			return true
		}
	}

	return false
}

func getLowestUnvisited(seen []bool, dists []int) int {
	idx := -1
	lowestDistance := math.MaxInt64

	for i := range seen {
		if seen[i] {
			continue
		}
		if lowestDistance > dists[i] {
			idx = i
			lowestDistance = dists[i]
		}
	}
	return idx
}

func main() {
}
