package main

import (
	"regexp"
	"strings"
)

func isNiceOld(word string) bool {
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

func isNiceNew(word string) bool {
	var (
		runes     = []rune(word)
		hasRepeat bool
		hasPair   bool
	)

	for i := 0; i < len(runes)-2; i++ {
		if runes[i] == runes[i+2] {
			hasPair = true
			break
		}
	}

	for i := 0; i < len(runes)-1; i++ {
		pair := string(runes[i]) + string(runes[i+1])
		rest := string(runes[i+2:])
		if strings.Count(rest, pair) > 0 {
			hasRepeat = true
			break
		}
	}

	return hasPair && hasRepeat
}
