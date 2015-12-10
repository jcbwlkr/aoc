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
