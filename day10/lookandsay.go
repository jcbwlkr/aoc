package main

import (
	"bytes"
	"strconv"
)

func lookAndSay(starting string, iterations int) string {
	ret := describe(starting)
	for i := 1; i < iterations; i++ {
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
