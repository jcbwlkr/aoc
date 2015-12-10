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
	paper, ribbon := 0, 0
	for scanner.Scan() {
		paper += scanner.Box().RequiredPaper()
		ribbon += scanner.Box().RequiredRibbon()
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln("Error scanning input", err)
	}

	fmt.Printf("Part 1 required paper %d sqft\n", paper)
	fmt.Printf("Part 2 required ribbon %d ft\n", ribbon)
}
