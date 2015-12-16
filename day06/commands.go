package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

// Command is an action to take on a whole range of lights addressed within
// Start and End points.
type Command struct {
	Action Action
	Start  Point
	End    Point
}

// Range gives all of the points in the range of a Command's Start and End.
func (c Command) Range() []Point {
	var ps []Point
	for x := c.Start.X; x <= c.End.X; x++ {
		for y := c.Start.Y; y <= c.End.Y; y++ {
			ps = append(ps, Point{X: x, Y: y})
		}
	}

	return ps
}

// ParseCommands parses the list of commands from a input file.
func ParseCommands(filename string) []Command {
	re := regexp.MustCompile("(turn on|turn off|toggle) ([0-9]*),([0-9]*) through ([0-9]*),([0-9]*)")

	input, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	var commands []Command
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		cmd := scanner.Text()
		m := re.FindStringSubmatch(cmd)
		commands = append(commands, Command{
			Action: Action(m[1]),
			Start:  Point{X: mustAtoi(m[2]), Y: mustAtoi(m[3])},
			End:    Point{X: mustAtoi(m[4]), Y: mustAtoi(m[5])},
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	return commands
}

func mustAtoi(a string) (i int) {
	var err error
	i, err = strconv.Atoi(a)
	if err != nil {
		log.Fatalln(err)
	}

	return i
}
