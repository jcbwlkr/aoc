package main

import (
	"bufio"
	"io"
	"regexp"
	"strconv"
	"strings"
)

// Routes knows the distances between cities.
type Routes map[string]map[string]int

// NewRoutes reads a newline separated list of routes.
func NewRoutes(rdr io.Reader) (Routes, error) {
	re := regexp.MustCompile(`^([A-Za-z ]+) to ([A-Za-z ]+) = ([0-9]+)$`)
	r := make(Routes)

	initMap := func(city string) {
		if _, ok := r[city]; !ok {
			r[city] = make(map[string]int)
		}
	}

	s := bufio.NewScanner(rdr)
	for s.Scan() {
		parts := re.FindStringSubmatch(strings.TrimSpace(s.Text()))
		initMap(parts[1])
		initMap(parts[2])
		distance, _ := strconv.Atoi(parts[3])

		r[parts[1]][parts[2]] = distance
		r[parts[2]][parts[1]] = distance
	}

	if err := s.Err(); err != nil {
		return Routes{}, err
	}

	return r, nil
}

// ShortestCircuit returns the distance of the shortest circuit that visits
// each city exactly once.
func (r Routes) ShortestCircuit() int {
	d, _ := r.distances()
	return d
}

// LongestCircuit returns the distance of the longest circuit that visits
// each city exactly once.
func (r Routes) LongestCircuit() int {
	_, d := r.distances()
	return d
}

func (r Routes) distances() (shortest, longest int) {
	var allCities []string
	for city := range r {
		allCities = append(allCities, city)
	}

	for i, path := range allCombinations(allCities) {
		var d int
		lastCity := path[0]
		for i := 1; i < len(path); i++ {
			d += r[lastCity][path[i]]
			lastCity = path[i]
		}

		if i == 0 {
			shortest, longest = d, d
		}
		if d < shortest {
			shortest = d
		}
		if d > longest {
			longest = d
		}
	}

	return
}

func allCombinations(in []string) [][]string {
	// Recursive base case
	if len(in) == 1 {
		return [][]string{[]string{in[0]}}
	}

	var combos [][]string

	for i, a := range in {
		var others []string
		for j, b := range in {
			if i != j {
				others = append(others, b)
			}
		}

		for _, combo := range allCombinations(others) {
			combos = append(combos, append(combo, a))
		}
	}

	// TODO eliminate duplicate combos that are the same forward and back

	return combos
}
