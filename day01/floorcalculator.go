package main

import "io"

func floorCalculator(r io.Reader) int {
	var (
		scratch = make([]byte, 32)
		floor   int
		read    int
	)
	for {
		if read, _ = r.Read(scratch); read == 0 {
			break
		}
		for i := 0; i < read; i++ {
			switch scratch[i] {
			case '(':
				floor++
			case ')':
				floor--
			}
		}
	}

	return floor
}
