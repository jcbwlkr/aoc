package main

import (
	"regexp"
)

func isNice(word string) bool {
	var (
		forbidden = regexp.MustCompile(`(ab|cd|pq|xy)`)
		vowels    int
		last      rune
		hasDouble bool
	)

	if len(forbidden.FindAllIndex([]byte(word), -1)) > 0 {
		return false
	}

	for _, r := range word {
		switch r {
		case 'a', 'e', 'i', 'o', 'u':
			vowels++
		}
		if r == last {
			hasDouble = true
		}
		last = r
		if hasDouble && vowels > 2 {
			return true
		}
	}

	return false
}
