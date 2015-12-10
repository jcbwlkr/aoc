package main

import (
	"bufio"
	"io"
)

// BoxScanner is an iterator providing instances of Box from an io.Reader split
// by newlines
type BoxScanner struct {
	s        *bufio.Scanner
	parseErr error
	next     Box
}

// NewBoxScanner creates a BoxScanner
func NewBoxScanner(r io.Reader) *BoxScanner {
	return &BoxScanner{
		s: bufio.NewScanner(r),
	}
}

// Scan reads the next Box from the reader into the BoxScanner. It returns true
// if a Box is read successfully. It returns false when an error is encountered
// or it reaches the end of its iput.
func (bs *BoxScanner) Scan() bool {
	if !bs.s.Scan() {
		bs.next = Box{}
		return false
	}
	box, err := NewBoxFromDimensions(bs.s.Text())
	if err != nil {
		bs.next = Box{}
		bs.parseErr = err
		return false
	}

	bs.next = box
	return true
}

// Box returns the most recent Box generated by a call to Scan
func (bs *BoxScanner) Box() Box {
	return bs.next
}

// Err returns the first error encountered by Scan
func (bs *BoxScanner) Err() error {
	if bs.parseErr != nil {
		return bs.parseErr
	}

	return bs.s.Err()
}
