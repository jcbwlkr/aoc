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

	r, err := NewRoutes(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Shortest route is: ")
	fmt.Println(r.ShortestCircuit())
	fmt.Print("Longest route is: ")
	fmt.Println(r.LongestCircuit())
}
