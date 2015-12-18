package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var code, chars, bigcode int

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ln := scanner.Text()

		code += len(ln)
		s, err := strconv.Unquote(ln)
		if err != nil {
			log.Fatal(err)
		}
		chars += len(s)

		bc := strconv.Quote(ln)
		bigcode += len(bc)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d chars of code yielding %d chars of strings with a difference of %d\n", code, chars, code-chars)
	fmt.Printf("Part 2: %d chars of double encoded code yielding a difference of %d\n", bigcode, bigcode-code)
}
