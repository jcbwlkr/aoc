package main

import (
	"strings"
	"testing"
)

func TestFloorCalculator(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"((()))", 0},
		{"()()()", 0},
		{"(", 1},
		{"((((((((", 8},
		{")", -1},
		{"))))))))", -8},
		{"(()", 1},
	}

	for i, test := range tests {
		reader := strings.NewReader(test.input)
		actual := floorCalculator(reader)
		if actual != test.expected {
			t.Errorf("Test %d: floorCalculator expected %v, actual %v", i, test.expected, actual)
		}
	}
}

func TestBasementFinder(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{")", 1},
		{"()())", 5},
		{"(()))", 5},
		{"(((((((()))))))))", 17},
		{"(((()", -1},
	}

	for i, test := range tests {
		reader := strings.NewReader(test.input)
		actual := basementFinder(reader)
		if actual != test.expected {
			t.Errorf("Test %d: basementFinder expected %v, actual %v", i, test.expected, actual)
		}
	}
}
