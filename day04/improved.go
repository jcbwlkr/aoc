package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"math"
	"strconv"
)

func mineSerial(input string, prefix string) int {
	var (
		// l is the min length of the md5 byte slice that is needed to generate a
		// hex string at least as long as the desired prefix
		l = int(math.Ceil(float64(len(prefix)) / float64(2)))

		// in is the input cast as a []byte. Do this before the loop so we don't
		// have to keep doing it.
		in = []byte(input)

		p = []byte(prefix)

		ore = make([]byte, hex.EncodedLen(l))
	)

	for i := 1; i < math.MaxInt32; i++ {
		// Combine the bytes of in plus the bytes of the string version of i then
		// get the md5 sum as another []byte
		sum := md5.Sum(append(
			in,
			[]byte(strconv.Itoa(i))...,
		))

		// Get the hex encoding of just the first part of the md5 sum slice
		hex.Encode(ore, sum[:l])

		if bytes.Equal(p, ore[:len(p)]) {
			return i
		}
	}

	return -1
}
