package main

import (
	"fmt"
	"math"
)

const (
	up = iota
	left
	down
	right
)

func main() {
	//const input = 368078 // Jacob's
	const input = 325489 // Ed's

	var x, y, max int
	pointing := right

	// cells[x][y] = 42
	//cells := make(map[int]map[int]int)

	var n int
	for n = 1; n < input; n++ {
		switch pointing {
		case up:
			y++
			if y == max {
				pointing = left
			}
		case left:
			x--
			if x == -max {
				pointing = down
			}
		case down:
			y--
			if y == -max {
				pointing = right
			}
		case right:
			x++

			// Only moving right is allowed to expand the board
			if x > max {
				max++
				pointing = up
			}
		}
	}

	fmt.Println("Distance from (0,0) is", math.Abs(float64(x))+math.Abs(float64(y)))
}

func dir(p int) string {
	switch p {
	case up:
		return "up"
	case left:
		return "left"
	case down:
		return "down"
	case right:
		return "right"
	}
	return "????"
}
