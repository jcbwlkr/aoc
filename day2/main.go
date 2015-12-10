package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := NewBoxScanner(file)
	var total int
	for scanner.Scan() {
		total += scanner.Box().RequiredPaper()
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln("Error scanning input", err)
	}

	fmt.Printf("Part 1 required paper %d sqft\n", total)
}
