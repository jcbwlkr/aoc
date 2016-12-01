package main

import (
	"testing"
)

func TestNewBox(t *testing.T) {
	expected := Box{
		Length: 1,
		Width:  2,
		Height: 3,
	}

	actual := NewBox(1, 2, 3)

	if actual != expected {
		t.Errorf("NewBox(1, 2, 3): expected %v, actual %v", expected, actual)
	}
}

func TestNewBoxFromDimensions(t *testing.T) {
	tests := []struct {
		input string
		box   Box
		err   bool
	}{
		{"1x2x3", NewBox(1, 2, 3), false},
		{"1x22x3", NewBox(1, 22, 3), false},
		{"1x2x33", NewBox(1, 2, 33), false},
		{"11x22x33", NewBox(11, 22, 33), false},
		{"", Box{}, true},
		{"banana", Box{}, true},
		{"2", Box{}, true},
		{"2x4", Box{}, true},
		{"2x4x5x6", Box{}, true},
		{"AxBxC", Box{}, true},
		{"Ax2x3", Box{}, true},
		{"1xBx3", Box{}, true},
		{"1x2xC", Box{}, true},
	}

	for i, test := range tests {
		box, err := NewBoxFromDimensions(test.input)
		switch {
		case err != nil && !test.err:
			t.Errorf("Test %d, NewBoxFromDimensions(%v): should not have errored, err was %v", i, test.input, err)
		case err == nil && test.err:
			t.Errorf("Test %d, NewBoxFromDimensions(%v): should have errored, err was nil", i, test.input)
		case box != test.box:
			t.Errorf("Test %d, NewBoxFromDimensions(%v): expected %v, actual %v", i, test.input, test.box, box)
		}
	}
}

func TestBoxRequiredPaper(t *testing.T) {
	tests := []struct {
		box      Box
		expected int
	}{
		{NewBox(2, 3, 4), 58},
		{NewBox(3, 2, 4), 58},
		{NewBox(3, 4, 2), 58},
		{NewBox(1, 1, 10), 43},
		{NewBox(1, 1, 1), 7},
	}

	for i, test := range tests {
		actual := test.box.RequiredPaper()
		if actual != test.expected {
			t.Errorf("Test %d, %#v.RequiredPaper(): expected %v, actual %v", i, test.box, test.expected, actual)
		}
	}
}

func TestBoxRequiredRibbon(t *testing.T) {
	tests := []struct {
		box      Box
		expected int
	}{
		{NewBox(2, 3, 4), 34},
		{NewBox(1, 1, 10), 14},
	}

	for i, test := range tests {
		actual := test.box.RequiredRibbon()
		if actual != test.expected {
			t.Errorf("Test %d, %#v.RequiredRibbon(): expected %v, actual %v", i, test.box, test.expected, actual)
		}
	}
}

func TestSmallest(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 2, 3}, 1},
		{[]int{2, 1, 3}, 1},
		{[]int{2, 3, 1}, 1},
		{[]int{}, 0},
	}

	for i, test := range tests {
		actual := smallest(test.nums...)
		if actual != test.expected {
			t.Errorf("Test %d, smallest(%v...): expected %v, actual %v", i, test.nums, test.expected, actual)
		}
	}
}
