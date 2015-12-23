package main

import (
	"bufio"
	"io"
	"regexp"
	"strconv"
)

// Reindeer can fly.
type Reindeer struct {
	Name     string
	Speed    int
	Duration int
	Rest     int
}

var parser = regexp.MustCompile(`^(\w+) can fly (\d+) km/s for (\d+) seconds, but then must rest for (\d+) seconds.$`)

// ParseLine parses a single line into a Reindeer.
func ParseLine(s string) Reindeer {
	if !parser.MatchString(s) {
		return Reindeer{}
	}

	parts := parser.FindStringSubmatch(s)

	return Reindeer{
		Name:     parts[1],
		Speed:    mustAtoi(parts[2]),
		Duration: mustAtoi(parts[3]),
		Rest:     mustAtoi(parts[4]),
	}
}

// ParseAll reads lines from an input and parses them into a list of Reindeer.
func ParseAll(r io.Reader) ([]Reindeer, error) {
	var rnd []Reindeer

	s := bufio.NewScanner(r)
	for s.Scan() {
		rnd = append(rnd, ParseLine(s.Text()))
	}

	if err := s.Err(); err != nil {
		return []Reindeer{}, err
	}

	return rnd, nil
}

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

// OldRace starts a race among a group of Reindeer lasting up to duration. At
// the end the winner is announced along with the distance they traveled.
func OldRace(racers []Reindeer, duration int) (winner Reindeer, distance int) {
	pos := make(map[Reindeer]int)
	energy := make(map[Reindeer]int)
	// Prime the energy map with the number of seconds each Reindeer can sprint
	// to begin with. Every second the value will be deprecated. When it goes
	// negative enough to hit their Rest value they reset.
	for _, r := range racers {
		energy[r] = r.Duration
	}

	for i := 0; i < duration; i++ {
		for _, r := range racers {
			if energy[r] == -r.Rest {
				energy[r] = r.Duration
			}
			if energy[r] > 0 {
				pos[r] += r.Speed
			}
			energy[r]--
		}
	}

	for r, d := range pos {
		if d > distance {
			winner, distance = r, d
		}
	}

	return
}

// NewRace starts a race among a group of Reindeer lasting up to duration. At
// the end the winner is announced along with the points they accumulates.
func NewRace(racers []Reindeer, duration int) (winner Reindeer, points int) {
	pos := make(map[Reindeer]int)
	pts := make(map[Reindeer]int)
	energy := make(map[Reindeer]int)
	// Prime the energy map with the number of seconds each Reindeer can sprint
	// to begin with. Every second the value will be deprecated. When it goes
	// negative enough to hit their Rest value they reset.
	for _, r := range racers {
		energy[r] = r.Duration
	}

	for i := 0; i < duration; i++ {
		for _, r := range racers {
			if energy[r] == -r.Rest {
				energy[r] = r.Duration
			}
			if energy[r] > 0 {
				pos[r] += r.Speed
			}
			energy[r]--
		}

		// Figure out who's in the lead
		lead := 0
		for _, d := range pos {
			if d > lead {
				lead = d
			}
		}

		// Give a point to everyone in the lead
		for r, d := range pos {
			if d == lead {
				pts[r]++
			}
		}
	}

	for r, p := range pts {
		if p > points {
			winner, points = r, p
		}
	}

	return
}
