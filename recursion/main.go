package main

type Point struct {
	x int
	y int
}

var dir = [][]int{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

func walk(maze [][]rune, wall rune, curr Point, end Point, seen [][]bool, path *[]Point) bool {
	if maze[curr.y][curr.x] == wall {
		return false
	}

	if curr.y == end.y && curr.x == end.x {
		*path = append(*path, curr)
		return true
	}

	if seen[curr.y][curr.x] {
		return false
	}

	*path = append(*path, curr)
	i := len(*path)
	seen[curr.y][curr.x] = true

	for _, point := range dir {
		x, y := point[0], point[1]
		if walk(maze, wall, Point{x: curr.x + x, y: curr.y + y}, end, seen, path) {
			return true
		}
	}
	*path = (*path)[:i-1]
	return false
}

func mazeSolver(maze []string, wall rune, start Point, end Point) []Point {
	// convert to rune
	mazeRune := make([][]rune, len(maze))
	path := []Point{}
	for i, str := range maze {
		mazeRune[i] = []rune(str)
	}

	seen := [][]bool{}

	for i := 0; i < len(maze); i++ {
		seen = append(seen, make([]bool, len(maze[0])))
	}

	if walk(mazeRune, wall, start, end, seen, &path) {
		return path
	}
	return path
}

func main() {
}
