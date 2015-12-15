package main

import "testing"

func TestIsNiceOld(t *testing.T) {
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
		actual := isNiceOld(test.word)
		if actual != test.expected {
			t.Errorf("Test %d: isNiceOld(%q) expected %v, actual %v", i, test.word, test.expected, actual)
		}
	}
}

func TestIsNiceNew(t *testing.T) {
	tests := []struct {
		word     string
		expected bool
	}{
		{"ababeejkj", true},         // repeated ab and interrupted pair jkj
		{"qjhvhtzxzqqjkmpb", true},  // repeated qj and interrupted pair zxz
		{"xxyxx", true},             // repeated xx and interrupted pair xyx
		{"aaa", false},              // repeated aa overlaps so doesn't count
		{"uurcxstgmygtbstg", false}, // repeated tg but no interrupted pair
		{"eeodomkazucvgmuy", false}, // interrupted pair odo but no repeated pair
		{"fart", false},             // no repeat or interrupted pairs
	}

	for i, test := range tests {
		actual := isNiceNew(test.word)
		if actual != test.expected {
			t.Errorf("Test %d: isNiceNew(%q) expected %v, actual %v", i, test.word, test.expected, actual)
		}
	}
}
