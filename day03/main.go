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

	houses := houseVisits(file)
	fmt.Println(len(houses))

	file, err = os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	houses = roboVisits(file)
	fmt.Println(len(houses))
}
