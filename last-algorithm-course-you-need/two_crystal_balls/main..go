package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(twoCrystalBallsProblem([]bool{false, false, false, true, true, true, true, true}))
}

func twoCrystalBallsProblem(breaks []bool) int {
	jmpAmount := int(math.Floor(math.Sqrt(float64(len(breaks)))))

	i := jmpAmount
	for ; i < len(breaks); i += jmpAmount {
		if breaks[i] {
			break
			fmt.Print("here")
		}
	}

	i -= jmpAmount
	for j := 0; j <= jmpAmount && i < len(breaks); i++ {
		j++
		if breaks[i] {
			return i
		}
	}

	return -1

}
