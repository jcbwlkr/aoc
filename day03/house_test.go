package main

import (
	"strings"
	"testing"
)

func TestHouseVisits(t *testing.T) {
	tests := []struct {
		input     string
		numHouses int
	}{
		{"", 1},
		{">", 2},
		{"><", 2},
		{"^>v<", 4},
		{"^v^v^v^v^v", 2},
	}

	for i, test := range tests {
		houses := houseVisits(strings.NewReader(test.input))
		actual := len(houses)
		if actual != test.numHouses {
			t.Errorf("Test %d: len(housesFromInput) expected %d, actual %d", i, test.numHouses, actual)
		}
	}
}

func TestRoboVisits(t *testing.T) {
	tests := []struct {
		input     string
		numHouses int
	}{
		{"", 1},
		{">", 2},
		{"><", 3},
		{"^>v<", 3},
		{"^v^v^v^v^v", 11},
	}

	for i, test := range tests {
		houses := roboVisits(strings.NewReader(test.input))
		actual := len(houses)
		if actual != test.numHouses {
			t.Errorf("Test %d: len(roboVisits) expected %d, actual %d", i, test.numHouses, actual)
		}
	}
}
