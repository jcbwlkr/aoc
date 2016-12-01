package main

import (
	"crypto/md5"
	"fmt"
	"math"
	"strings"
)

func mineOriginal(input string, prefix string) int {
	for i := 1; i < math.MaxInt32; i++ {
		shaft := input + fmt.Sprintf("%d", i)
		ore := fmt.Sprintf("%x", md5.Sum([]byte(shaft)))

		if strings.HasPrefix(ore, prefix) {
			return i
		}
	}

	return -1
}
