package main

import "testing"

var tests = []struct {
	prefix   string
	expected int
}{
	{"00000", 117946},
	{"000000", 3938038},
}

func TestOriginal(t *testing.T) {
	for _, test := range tests {
		actual := mineOriginal("ckczppom", test.prefix)
		if actual != test.expected {
			t.Errorf("mineOriginal prefix %s = %d, expected %d", test.prefix, actual, test.expected)
		}
	}
}

func TestSerial(t *testing.T) {
	for _, test := range tests {
		actual := mineSerial("ckczppom", test.prefix)
		if actual != test.expected {
			t.Errorf("mineSerial prefix %s = %d, expected %d", test.prefix, actual, test.expected)
		}
	}
}

func TestConcurrent(t *testing.T) {
	for _, test := range tests {
		actual := mineConcurrent("ckczppom", test.prefix)
		if actual != test.expected {
			t.Errorf("mineConcurrent prefix %s = %d, expected %d", test.prefix, actual, test.expected)
		}
	}
}

var result int

func BenchmarkOriginal5(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = mineOriginal("ckczppom", "00000")
	}
	result = r
}

func BenchmarkOriginal6(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = mineOriginal("ckczppom", "000000")
	}
	result = r
}

func BenchmarkSerial5(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = mineSerial("ckczppom", "00000")
	}
	result = r
}

func BenchmarkSerial6(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = mineSerial("ckczppom", "000000")
	}
	result = r
}

func BenchmarkConcurrent5(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = mineConcurrent("ckczppom", "00000")
	}
	result = r
}

func BenchmarkConcurrent6(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = mineConcurrent("ckczppom", "000000")
	}
	result = r
}
