package main

import (
	"strings"
	"testing"
)

var (
	alice = Guest{
		Name: "Alice",
		Opinions: map[string]int{
			"Bob":   54,
			"Carol": 79,
			"David": -2,
		},
	}
	bob = Guest{
		Name: "Bob",
		Opinions: map[string]int{
			"Alice": 83,
			"Carol": -7,
			"David": -63,
		},
	}
	carol = Guest{
		Name: "Carol",
		Opinions: map[string]int{
			"Alice": -62,
			"Bob":   60,
			"David": 55,
		},
	}
	david = Guest{
		Name: "David",
		Opinions: map[string]int{
			"Alice": 46,
			"Bob":   -7,
			"Carol": 41,
		},
	}
)

func TestNewTable(t *testing.T) {
	const input = `Alice would gain 54 happiness units by sitting next to Bob.
Alice would lose 79 happiness units by sitting next to Carol.
Alice would lose 2 happiness units by sitting next to David.
Bob would gain 83 happiness units by sitting next to Alice.
Bob would lose 7 happiness units by sitting next to Carol.
Bob would lose 63 happiness units by sitting next to David.
Carol would lose 62 happiness units by sitting next to Alice.
Carol would gain 60 happiness units by sitting next to Bob.
Carol would gain 55 happiness units by sitting next to David.
David would gain 46 happiness units by sitting next to Alice.
David would lose 7 happiness units by sitting next to Bob.
David would gain 41 happiness units by sitting next to Carol.
`
	r := strings.NewReader(input)

	tbl := NewTable(r)
	expected := Table{alice, bob, carol, david}

	if !equal(tbl, expected) {
		t.Errorf("NewTable() = %v, expected %v", tbl, expected)
	}
}

func TestTableLeftAndRight(t *testing.T) {
	table := Table{carol, david, alice}

	tests := []struct {
		guest Guest
		left  Guest
		right Guest
	}{
		{carol, alice, david},
		{david, carol, alice},
		{alice, david, carol},
	}

	for i, test := range tests {
		left := table.Left(test.guest)
		if left.Name != test.left.Name {
			t.Errorf("%d: table.Left(%v) = %v, expected %v", i, test.guest, left, test.left)
		}
		right := table.Right(test.guest)
		if right.Name != test.right.Name {
			t.Errorf("%d: table.Right(%v) = %v, expected %v", i, test.guest, right, test.right)
		}
	}
}

func TestTableTotalHappiness(t *testing.T) {
	table := Table{carol, david, alice, bob}
	expected := 330
	actual := table.TotalHappiness()
	if actual != expected {
		t.Errorf("table.TotalHappiness() = %d, expected %d", actual, expected)
	}
}

func TestTableOptimize(t *testing.T) {
	table := Table{alice, bob, david, carol}
	expected := Table{carol, david, alice, bob}
	table.Optimize()
	if !equal(table, expected) {
		t.Errorf("table.Optimize()\nactual   %v\nexpected %v", table, expected)
	}
}

func equal(tblA, tblB Table) bool {
	if len(tblA) != len(tblB) {
		return false
	}

	for i := 0; i < len(tblA); i++ {
		// We need to find guest a[i] in the b table then ensure their left
		// and right guests are the same.
		guestA := tblA[i]

		var guestB Guest
		for j := range tblB {
			if tblB[j].Eq(guestA) {
				guestB = tblB[j]
				break
			}
		}
		if !guestA.Eq(guestB) {
			return false // Did not find guestA in table b
		}

		// If my left is your left and my right is your right we're good
		if tblA.Left(guestA).Eq(tblB.Left(guestB)) && tblA.Right(guestA).Eq(tblB.Right(guestB)) {
			continue
		}
		// If the tables are reversed then my left is your right and my right is your left
		if tblA.Left(guestA).Eq(tblB.Right(guestB)) && tblA.Right(guestA).Eq(tblB.Left(guestB)) {
			continue
		}

		return false
	}

	return true
}
