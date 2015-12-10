package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestBoxScanner(t *testing.T) {
	reader := strings.NewReader("1x2x3\n2x1x3\n2x3x1")
	expected := []Box{
		NewBox(1, 2, 3),
		NewBox(2, 1, 3),
		NewBox(2, 3, 1),
	}

	scanner := NewBoxScanner(reader)

	var actual []Box
	for scanner.Scan() {
		actual = append(actual, scanner.Box())
	}

	if err := scanner.Err(); err != nil {
		t.Fatalf("BoxScanner.Err() expected nil, actual %v", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("BoxScanner run expected %v, actual %v", expected, actual)
	}
}

func TestBoxScannerError(t *testing.T) {
	reader := strings.NewReader("1x2x3\n2x3x4\nBanana\n2x3x1")
	expectedIterCount := 2

	scanner := NewBoxScanner(reader)

	var iterCount int
	for scanner.Scan() {
		iterCount++
	}

	if err := scanner.Err(); err == nil {
		t.Error("BoxScanner.Err() expected err, got nil")
	}
	if iterCount != expectedIterCount {
		t.Errorf("BoxScanner.Scan() should have returned true %d times, actually %d", iterCount, expectedIterCount)
	}
}
