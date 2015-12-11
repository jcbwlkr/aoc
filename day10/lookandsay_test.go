package main

import "testing"

func TestLookAndSay(t *testing.T) {
	tests := []struct {
		starting   string
		iterations int
		expected   string
	}{
		{"1", 1, "11"},
		{"1", 2, "21"},
		{"1", 3, "1211"},
		{"1211", 1, "111221"},
	}

	for i, test := range tests {
		actual := lookAndSay(test.starting, test.iterations)
		if actual != test.expected {
			t.Errorf("Test %d: lookAndSay(%v, %d): expected %v, actual %v", i, test.starting, test.iterations, test.expected, actual)
		}
	}
}

func TestDescribeNum(t *testing.T) {
	tests := []struct {
		num      string
		expected string
	}{
		{"1", "11"},
		{"11", "21"},
		{"21", "1211"},
		{"1211", "111221"},
	}

	for i, test := range tests {
		actual := describe(test.num)
		if actual != test.expected {
			t.Errorf("Test %d: describe(%v): expected %v, actual %v", i, test.num, test.expected, actual)
		}
	}
}

func BenchmarkLookAndSay(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lookAndSay("1113122113", 30)
	}
}
