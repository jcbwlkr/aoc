package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestNewCircuit(t *testing.T) {
	r := strings.NewReader(`NOT b -> c
  c -> d
  d LSHIFT 2 -> a
  e AND f -> b
  12345 -> e
  54321 -> f`)

	c, err := NewCircuit(r)

	if err != nil {
		t.Fatalf("NewCircuit() err be nil, was %v", err)
	}

	expected := map[string]string{
		"a": "d LSHIFT 2",
		"b": "e AND f",
		"c": "NOT b",
		"d": "c",
		"e": "12345",
		"f": "54321",
	}

	if !reflect.DeepEqual(c.wires, expected) {
		t.Errorf("NewCircuit().wires: expected %v, actual %v", expected, c.wires)
	}
}

func TestSimpleCircuit(t *testing.T) {
	c := &Circuit{
		wires: map[string]string{
			"x": "123",
			"y": "456",
			"d": "x AND y",
			"e": "x OR y",
			"f": "x LSHIFT 2",
			"g": "y RSHIFT 2",
			"h": "NOT x",
			"i": "NOT y",
			"j": "x",
			"k": "i",
			"l": "1 AND x",
		},
		signals: map[string]uint16{},
	}

	tests := []struct {
		wire  string
		value uint16
	}{
		{"x", 123},
		{"y", 456},
		{"d", 72},
		{"e", 507},
		{"f", 492},
		{"g", 114},
		{"h", 65412},
		{"i", 65079},
		{"j", 123},
		{"k", 65079},
		{"l", 1},
	}

	for i, test := range tests {
		actual := c.Read(test.wire)
		if actual != test.value {
			t.Errorf("Test %d: c.Read(%q), expected %v, actual %v", i, test.wire, test.value, actual)
		}
	}
}
