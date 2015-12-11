package main

import "strconv"

func lookAndSay(starting string, iterations int) string {
	ret := describe(starting)
	for i := 1; i < iterations; i++ {
		ret = describe(ret)
	}
	return ret
}

func describe(num string) string {
	var (
		runs     [][]rune
		run      []rune
		lastRune rune
		ret      string
	)
	for i, r := range []rune(num) {
		if i != 0 && r != lastRune {
			runs = append(runs, run)
			run = []rune{r}
			lastRune = r
			continue
		}

		lastRune = r
		run = append(run, r)
	}

	runs = append(runs, run)

	for _, run := range runs {
		ret += strconv.Itoa(len(run)) + string(run[0])
	}

	return ret
}
