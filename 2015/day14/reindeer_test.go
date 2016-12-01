package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseLine(t *testing.T) {
	input := `Rudolph can fly 22 km/s for 8 seconds, but then must rest for 165 seconds.`
	expected := Reindeer{
		Name:     "Rudolph",
		Speed:    22,
		Duration: 8,
		Rest:     165,
	}

	if actual := ParseLine(input); actual != expected {
		t.Errorf("ParseLine(%q): expected %v, actual %v", input, expected, actual)
	}
}

func TestParseAll(t *testing.T) {
	input := strings.NewReader(`Rudolph can fly 22 km/s for 8 seconds, but then must rest for 165 seconds.
Cupid can fly 8 km/s for 17 seconds, but then must rest for 114 seconds.`)

	expected := []Reindeer{
		{Name: "Rudolph", Speed: 22, Duration: 8, Rest: 165},
		{Name: "Cupid", Speed: 8, Duration: 17, Rest: 114},
	}

	actual, err := ParseAll(input)
	if err != nil {
		t.Errorf("Error should be nil, was %v", err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("ParseAll(...): expected %v, actual %v", expected, actual)
	}
}

func TestOldRace(t *testing.T) {
	comet := Reindeer{Name: "Comet", Speed: 14, Duration: 10, Rest: 127}
	dancer := Reindeer{Name: "Dancer", Speed: 16, Duration: 11, Rest: 162}

	winner, distance := OldRace([]Reindeer{comet, dancer}, 1000)

	if winner != comet {
		t.Errorf("Winner should have been comet. Was %v", winner)
	}
	if distance != 1120 {
		t.Errorf("Distance should have been 1120. Was %v", distance)
	}
}

func TestNewRace(t *testing.T) {
	comet := Reindeer{Name: "Comet", Speed: 14, Duration: 10, Rest: 127}
	dancer := Reindeer{Name: "Dancer", Speed: 16, Duration: 11, Rest: 162}

	winner, points := NewRace([]Reindeer{comet, dancer}, 1000)

	if winner != dancer {
		t.Errorf("Winner should have been dancer. Was %v", winner)
	}
	if points != 689 {
		t.Errorf("Points should have been 689. Was %v", points)
	}
}
