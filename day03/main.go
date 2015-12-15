package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	houses, err := houseVisits(file)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(len(houses))

	file, err = os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	houses, err = roboVisits(file)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(len(houses))
}
