package main

import (
	"bufio"
	"errors"
	"io"
	"regexp"
	"strconv"
	"strings"
)

// Possible errors
var (
	ErrBadInput = errors.New("input is malformed")
)

// Circuit holds a series of connected wires with signals.
type Circuit struct {
	wires   map[string]string // nodes with instructions
	signals map[string]uint16 // resolved values
}

// NewCircuit reads a newline separated input stream to create a series of
// wires.
func NewCircuit(r io.Reader) (*Circuit, error) {
	c := &Circuit{
		wires:   make(map[string]string),
		signals: make(map[string]uint16),
	}

	s := bufio.NewScanner(r)
	for s.Scan() {
		ln := s.Text()
		parts := strings.Split(ln, "->")
		if len(parts) != 2 {
			return nil, ErrBadInput
		}
		c.wires[strings.TrimSpace(parts[1])] = strings.TrimSpace(parts[0])
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return c, nil
}

// RegExps used by Read
var (
	reFollow = regexp.MustCompile(`^[a-z]+$`)
	reAnd    = regexp.MustCompile(`^([a-z]+) AND ([a-z]+)$`)
	reAndInt = regexp.MustCompile(`^([0-9]+) AND ([a-z]+)$`)
	reOr     = regexp.MustCompile(`^([a-z]+) OR ([a-z]+)$`)
	reNot    = regexp.MustCompile(`^NOT ([a-z]+)$`)
	reLshift = regexp.MustCompile(`^([a-z]+) LSHIFT ([0-9]+)$`)
	reRshift = regexp.MustCompile(`^([a-z]+) RSHIFT ([0-9]+)$`)
)

func (c *Circuit) Read(wire string) (signal uint16) {
	// Before we're done record the resolution in our signal map
	defer func() {
		c.signals[wire] = signal
	}()

	// If we've already resolved this wire we can just return its signal.
	if s, ok := c.signals[wire]; ok {
		return s
	}

	val := c.wires[wire]
	// Base case. A wire with just a numeric value can be read directly.
	if n, err := strconv.Atoi(val); err == nil {
		return uint16(n)
	}

	switch {
	case reFollow.MatchString(val):
		// Wire just points to another wire. Read that one.
		return c.Read(val)
	case reAnd.MatchString(val):
		// Wire is the AND of two other wires
		pts := reAnd.FindStringSubmatch(val)
		return c.Read(pts[1]) & c.Read(pts[2])
	case reAndInt.MatchString(val):
		// Wire is the AND of some int and another wire
		pts := reAndInt.FindStringSubmatch(val)
		x, _ := strconv.Atoi(pts[1])
		return uint16(x) & c.Read(pts[2])
	case reOr.MatchString(val):
		// Wire is the OR of two other wires
		pts := reOr.FindStringSubmatch(val)
		return c.Read(pts[1]) | c.Read(pts[2])
	case reNot.MatchString(val):
		// Wire is bitwise complement of another wire
		pts := reNot.FindStringSubmatch(val)
		return ^c.Read(pts[1])
	case reLshift.MatchString(val):
		// Wire is another wire left shifted by some value
		pts := reLshift.FindStringSubmatch(val)
		shift, _ := strconv.Atoi(pts[2])
		return c.Read(pts[1]) << uint(shift)
	case reRshift.MatchString(val):
		// Wire is another wire right shifted by some value
		pts := reRshift.FindStringSubmatch(val)
		shift, _ := strconv.Atoi(pts[2])
		return c.Read(pts[1]) >> uint(shift)
	}

	return 0
}
