package main

import (
	"bufio"
	"io"
)

func floorCalculator(r io.Reader) int {
	var floor int
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanBytes)

	for s.Scan() {
		switch s.Text() {
		case "(":
			floor++
		case ")":
			floor--
		}
	}

	return floor
}

func basementFinder(r io.Reader) int {
	var (
		floor int
		index int
	)
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanBytes)

	for s.Scan() {
		index++
		switch s.Text() {
		case "(":
			floor++
		case ")":
			floor--
		}
		if floor == -1 {
			return index
		}
	}

	return -1
}
