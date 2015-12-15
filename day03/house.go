package main

import "io"
import "io/ioutil"

// A House is a home
type House struct {
	X, Y int // The coordinates of the house in our grid
}

func houseVisits(r io.Reader) (map[House]int, error) {
	cmds, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return work(cmds, 1), nil
}

func roboVisits(r io.Reader) (map[House]int, error) {
	cmds, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return work(cmds, 2), nil
}

func work(cmds []byte, numWorkers int) map[House]int {
	workers := make([]Worker, numWorkers)
	m := map[House]int{
		{X: 0, Y: 0}: len(workers), // Starting location always gets visits
	}
	i := 0

	for _, c := range cmds {
		if i >= numWorkers {
			i = 0
		}
		m[workers[i].Move(rune(c))]++
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
