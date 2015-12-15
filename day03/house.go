package main

import "io"
import "io/ioutil"

// A House is a home
type House struct {
	X, Y int // The coordinates of the house in our grid
}

func houseVisits(r io.Reader) (map[House]int, error) {
	var x, y int
	m := map[House]int{
		{X: 0, Y: 0}: 1, // Starting location always gets a house
	}

	cmds, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	for _, c := range cmds {
		switch c {
		case '>':
			x++
		case '<':
			x--
		case '^':
			y++
		case 'v':
			y--
		}
		m[House{X: x, Y: y}]++
	}

	return m, nil
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

func roboVisits(r io.Reader) (map[House]int, error) {
	var santa, robot Worker
	var i int
	m := map[House]int{
		{X: 0, Y: 0}: 2, // Starting location always gets two visits
	}

	cmds, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	for _, c := range cmds {
		if i%2 == 0 {
			m[santa.Move(rune(c))]++
		} else {
			m[robot.Move(rune(c))]++
		}
		i++
	}

	return m, nil
}
