package main

import "strconv"

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

func describe(num string) string {
	var (
		lastRune rune
		count    int
		ret      string
	)
	for i, r := range []rune(num) {
		if i != 0 && r != lastRune {
			ret += strconv.Itoa(count) + string(lastRune)
			count = 1
			lastRune = r
			continue
		}

		count++
		lastRune = r
	}

	ret += strconv.Itoa(count) + string(lastRune)

	return ret
}
