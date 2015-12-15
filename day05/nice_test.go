package main

import "testing"

func TestIsNice(t *testing.T) {
	tests := []struct {
		word     string
		expected bool
	}{
		{"aaa", true},
		{"ugknbfddgicrmopn", true},
		{"jchzalrnumimnmhp", false}, // No double letter
		{"dvszwmarrgswjxmb", false}, // Only has one vowel
		{"dvszwmarreswjxmb", false}, // Only has two vowels
		{"haegwjzuvuyypabu", false}, // Contains forbidden ab
		{"haegwjzuvuyypcdu", false}, // Contains forbidden cd
		{"haegwjzuvuyyppqu", false}, // Contains forbidden pq
		{"haegwjzuvuyypxyu", false}, // Contains forbidden xy
	}

	for i, test := range tests {
		actual := isNice(test.word)
		if actual != test.expected {
			t.Errorf("Test %d: isNice(%q) expected %v, actual %v", i, test.word, test.expected, actual)
		}
	}
}
