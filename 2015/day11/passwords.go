package main

import (
	"regexp"
)

func isValid(pw string) bool {
	var (
		forbidden = regexp.MustCompile("[iol]")
		pairs     = map[string]bool{}
		straight  bool
	)

	if forbidden.MatchString(pw) {
		return false
	}

	runes := []rune(pw)
	for i := 1; i < len(runes); i++ {
		r := runes[i]
		if r == runes[i-1] {
			pairs[string([]rune{r, runes[i-1]})] = true
		}

		if i > 1 && runes[i-2] == r-2 && runes[i-1] == r-1 {
			straight = true
		}
	}

	return straight && len(pairs) > 1
}

func increment(pw string) string {
	word := []rune(pw)
	for i := len(word) - 1; i >= 0; i-- {
		r := word[i] + 1

		if r > 'z' {
			r = 'a'
		}

		word[i] = r

		if r > 'a' {
			break
		}
	}

	return string(word)
}
