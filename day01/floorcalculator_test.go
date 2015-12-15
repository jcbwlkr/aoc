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
			t.Errorf("Test %d: expected %v, actual %v", i, test.expected, actual)
		}
	}
}
