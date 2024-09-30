package main

type node struct {
	To     int
	Weight int
}

/*
//      (1) --- (4) ---- (5)
//    /  |       |       /|
// (0)   | ------|------- |
//    \  |/      |        |
//      (2) --- (3) ---- (6)
*/
var graph1 = [][]node{
	{
		{To: 1, Weight: 3},
		{To: 2, Weight: 1},
	},
	{
		{To: 0, Weight: 3},
		{To: 2, Weight: 4},
		{To: 4, Weight: 1},
	},
	{
		{To: 1, Weight: 4},
		{To: 3, Weight: 7},
		{To: 0, Weight: 1},
	},
	{
		{To: 2, Weight: 7},
		{To: 4, Weight: 5},
		{To: 6, Weight: 1},
	},
	{
		{To: 1, Weight: 1},
		{To: 3, Weight: 5},
		{To: 5, Weight: 2},
	},
	{
		{To: 6, Weight: 1},
		{To: 4, Weight: 2},
		{To: 2, Weight: 18},
	},
	{
		{To: 3, Weight: 1},
		{To: 5, Weight: 1},
	},
}

//     >(1)<--->(4) ---->(5)
//    /          |       /|
// (0)     ------|------- |
//    \   v      v        v
//     >(2) --> (3) <----(6)

var graph2 = [][]node{
	{
		{To: 1, Weight: 3},
		{To: 2, Weight: 1},
	},
	{
		{To: 4, Weight: 1},
	},
	{
		{To: 3, Weight: 7},
	},
	{},
	{
		{To: 1, Weight: 1},
		{To: 3, Weight: 5},
		{To: 5, Weight: 2},
	},
	{
		{To: 2, Weight: 18},
		{To: 6, Weight: 1},
	},
	{
		{To: 3, Weight: 1},
	},
}

func DFSGraphList(graph [][]node, source, needle int) []int {
	seen := make([]bool, len(graph))
	path := []int{}
	Walk(graph, source, needle, &path, seen)
	return path
}

func Walk(graph [][]node, curr, needle int, path *[]int, seen []bool) bool {
	if seen[curr] {
		return false
	}

	seen[curr] = true

	*path = append(*path, curr)
	if curr == needle {
		return true
	}

	for _, edge := range graph[curr] {
		if Walk(graph, edge.To, needle, path, seen) {
			return true
		}
	}
	*path = (*path)[:len(*path)-1]

	return false
}

func main() {
}
