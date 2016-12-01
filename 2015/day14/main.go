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

	racers, err := ParseAll(file)
	if err != nil {
		log.Fatal(err)
	}

	winner1, distance := OldRace(racers, 2503)
	fmt.Println("Part 1:", winner1.Name, "wins with a total distance of", distance)

	winner2, points := NewRace(racers, 2503)
	fmt.Println("Part 2:", winner2.Name, "wins with a total of", points, "points")
}
