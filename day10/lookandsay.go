package main

import (
	"bytes"
	"strconv"
)

// Printlner can print a ln
type Printlner interface {
	Println(args ...interface{})
}

func lookAndSay(starting string, iterations int, logger Printlner) string {
	logger.Println("starting iteration 1")
	ret := describe(starting)
	for i := 1; i < iterations; i++ {
		logger.Println("starting iteration", i+1)
		ret = describe(ret)
	}
	return ret
}

func check(_ interface{}, err error) {
	if err != nil {
		panic(err)
	}
}

func describe(num string) string {
	var (
		lastRune rune
		count    int
		ret      bytes.Buffer
	)
	for i, r := range []rune(num) {
		if i != 0 && r != lastRune {
			check(ret.WriteString(strconv.Itoa(count)))
			check(ret.WriteRune(lastRune))
			count = 1
			lastRune = r
			continue
		}

		count++
		lastRune = r
	}

	check(ret.WriteString(strconv.Itoa(count)))
	check(ret.WriteRune(lastRune))

	return ret.String()
}
