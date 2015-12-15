package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	scanner := bufio.NewScanner(file)
	var niceOld, niceNew int
	for scanner.Scan() {
		if isNiceOld(scanner.Text()) {
			niceOld++
		}
		if isNiceNew(scanner.Text()) {
			niceNew++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	fmt.Println(niceOld, niceNew)
}
