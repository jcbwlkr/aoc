package main

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		word     string
		expected bool
	}{
		{"hijklmmn", false}, // contains i and l
		{"abbceffg", false}, // does not have a straight
		{"abbcegjk", false}, // does not have two pair
		{"abbbeabc", false}, // two pairs overlap (bbb doesn't count)
		{"abcdffaa", true},
		{"ghjaabcc", true},
	}

	for i, test := range tests {
		actual := isValid(test.word)
		if actual != test.expected {
			t.Errorf("Test %d: isValid(%q): expected %v, actual %v", i, test.word, test.expected, actual)
		}
	}
}

func TestIncrement(t *testing.T) {
	tests := []struct {
		word     string
		expected string
	}{
		{"a", "b"},
		{"x", "y"},
		{"aa", "ab"},
		{"az", "ba"},
		{"abcdefg", "abcdefh"},
		{"azzzzzz", "baaaaaa"},
	}

	for i, test := range tests {
		actual := increment(test.word)
		if actual != test.expected {
			t.Errorf("Test %d: increment(%q): expected %v, actual %v", i, test.word, test.expected, actual)
		}
	}
}
