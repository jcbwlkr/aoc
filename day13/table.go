package main

import (
	"bufio"
	"io"
	"regexp"
	"strconv"
)

// Guest is a person attending our party and their opinions of all other
// Guests (by name).
type Guest struct {
	Name     string
	Opinions map[string]int
}

// String implements stringer for easier debugging
func (g Guest) String() string {
	return g.Name
}

// Eq provides an equality comparison for two Guests
func (g Guest) Eq(other Guest) bool {
	return g.Name == other.Name
}

// Table is the arrangements of all Guests at our party. The table is
// circular so the guests at index 0 and len(table)-1 are adjacent.
type Table []Guest

// NewTable scans an io.Reader for the expected textual representation of
// a Table.
func NewTable(r io.Reader) Table {
	var (
		t  Table
		g  Guest
		re = regexp.MustCompile(`(\w+) would (gain|lose) (\d+) happiness units by sitting next to (\w+).`)
		s  = bufio.NewScanner(r)
	)

	for s.Scan() {
		m := re.FindStringSubmatch(s.Text())
		var (
			name  = m[1]
			dir   = m[2]
			pts   = mustAtoi(m[3])
			other = m[4]
		)
		if dir == "lose" {
			pts *= -1
		}

		if name != g.Name {
			if g.Name != "" {
				t = append(t, g)
			}
			g.Name = name
			g.Opinions = make(map[string]int)
		}
		g.Opinions[other] = pts
	}
	t = append(t, g)
	if err := s.Err(); err != nil {
		panic(err)
	}

	return t
}

// TotalHappiness returns the total change in happiness of the current
// seating order.
func (t Table) TotalHappiness() int {
	var h int
	l := len(t)
	for i := 0; i < l; i++ {
		g := t[i]
		h += g.Opinions[t.Left(g).Name]
		h += g.Opinions[t.Right(g).Name]
	}

	return h
}

// Left is the Guest sitting to the "left" of the specified guest.
func (t Table) Left(g Guest) Guest {
	var idx int
	for i := range t {
		if t[i].Eq(g) {
			idx = i - 1
			break
		}
	}

	if idx < 0 {
		idx = len(t) - 1
	}

	return t[idx]
}

// Right is the Guest sitting to the "right" of the specified guest.
func (t Table) Right(g Guest) Guest {
	var idx int
	for i := range t {
		if t[i].Eq(g) {
			idx = i + 1
			break
		}
	}

	if idx >= len(t) {
		idx = 0
	}

	return t[idx]
}

// Optimize will shuffle the order of Guests at the table to find the
// optimum value for TotalHappiness.
func (t *Table) Optimize() {
	var (
		bestScore int
		bestT     Table
	)

	for _, tbl := range allCombinations(*t) {
		if s := tbl.TotalHappiness(); s > bestScore {
			bestScore = s
			bestT = tbl
		}
	}

	*t = bestT
}

func allCombinations(t Table) []Table {
	if len(t) == 1 {
		return []Table{t}
	}

	var combos []Table

	for i, a := range t {
		var others Table
		for j, b := range t {
			if i != j {
				others = append(others, b)
			}
		}

		for _, combo := range allCombinations(others) {
			combos = append(combos, append(combo, a))
		}
	}

	return combos
}

func mustAtoi(a string) int {
	i, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	}

	return i
}
