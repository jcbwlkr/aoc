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
	nice := 0
	for scanner.Scan() {
		if isNice(scanner.Text()) {
			nice++
		}
	}

	fmt.Println(nice)
}
