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
	fmt.Println("Santa first entered the basement on ", basementFinder(file))

	file, err = os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Santa ends on ", floorCalculator(file))
}
