package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestNewRoutes(t *testing.T) {
	rdr := strings.NewReader(`London to Dublin = 464
  London to Belfast = 518
  Dublin to Belfast = 141`)

	r, err := NewRoutes(rdr)
	if err != nil {
		t.Fatalf("NewRoutes() error should be nil, was %v", err)
	}

	expected := Routes{
		"London": map[string]int{
			"Dublin":  464,
			"Belfast": 518,
		},
		"Dublin": map[string]int{
			"London":  464,
			"Belfast": 141,
		},
		"Belfast": map[string]int{
			"London": 518,
			"Dublin": 141,
		},
	}

	if !reflect.DeepEqual(r, expected) {
		t.Errorf("NewRoutes() expected %v actual %v", expected, r)
	}
}

func TestShortest(t *testing.T) {
	rdr := strings.NewReader(`London to Dublin = 464
  London to Belfast = 518
  Dublin to Belfast = 141`)

	r, err := NewRoutes(rdr)
	if err != nil {
		t.Fatalf("NewRoutes() error should be nil, was %v", err)
	}

	if expected, actual := 605, r.ShortestCircuit(); actual != expected {
		t.Errorf("ShortestCircuit(): expected %v, actual %v", expected, actual)
	}
}

func TestAllCombinations(t *testing.T) {
	input := []string{"a", "b", "c"}
	expected := [][]string{
		{"c", "b", "a"},
		{"b", "c", "a"},
		{"c", "a", "b"},
		{"a", "c", "b"},
		{"b", "a", "c"},
		{"a", "b", "c"},
	}

	actual := allCombinations(input)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("allCombinations(%v):\nexpected: %v\nactual:   %v", input, expected, actual)
	}
}
