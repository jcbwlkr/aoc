package main

import "bufio"
import "io"

// A House is a home
type House struct {
	X, Y int // The coordinates of the house in our grid
}

func houseVisits(r io.Reader) map[House]int {
	return work(r, 1)
}

func roboVisits(r io.Reader) map[House]int {
	return work(r, 2)
}

func work(r io.Reader, numWorkers int) map[House]int {
	var (
		rdr     = bufio.NewReader(r)
		workers = make([]Worker, numWorkers)
		m       = map[House]int{{X: 0, Y: 0}: len(workers)} // Starting location always gets visits
		i       = 0
	)

	for {
		c, _, err := rdr.ReadRune()
		if err != nil {
			break
		}
		if i >= numWorkers {
			i = 0
		}
		m[workers[i].Move(c)]++
		i++
	}
	return m
}

// A Worker delivers presents
type Worker struct {
	X, Y int // Current position
}

// Move accepts one rune of the list "^><v" and returns the Worker's new
// location.
func (w *Worker) Move(dir rune) House {
	switch dir {
	case '>':
		w.X++
	case '<':
		w.X--
	case '^':
		w.Y++
	case 'v':
		w.Y--
	}

	return House{X: w.X, Y: w.Y}
}
