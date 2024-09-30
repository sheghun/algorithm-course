package main

var graph = [][]int{
	{0, 3, 1, 0, 0, 0, 0},
	{0, 0, 0, 0, 1, 0, 0},
	{0, 0, 7, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0},
	{0, 1, 0, 5, 0, 2, 0},
	{0, 0, 18, 0, 0, 0, 1},
	{0, 0, 0, 1, 0, 0, 1},
}

func bfs(graph [][]int, source int, needle int) []int {
	// seen
	seen := make([]bool, len(graph))
	prev := make([]int, len(graph))
	out := []int{}

	for i := range prev {
		prev[i] = -1
	}

	seen[source] = true
	q := []int{source}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		for i := 0; i < len(graph[curr]); i++ {
			if graph[curr][i] == 0 || seen[i] {
				continue
			}

			if curr == needle {
				break
			}

			prev[i] = curr
			seen[i] = true
			q = append(q, i)
		}
	}

	if prev[needle] == -1 {
		return []int{}
	}

	for i := needle; prev[i] != -1; i = prev[i] {
		out = append(out, i)
	}
	out = append(out, source)

	// Reverse the out slice
	for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
		out[i], out[j] = out[j], out[i]
	}

	return out
}

func main() {
	bfs(graph, 0, 6)
}
