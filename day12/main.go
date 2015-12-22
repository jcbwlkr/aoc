package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.json")
	if err != nil {
		log.Fatal(err)
	}

	var data interface{}
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		log.Fatal(err)
	}

	fmt.Println("AoC Day 12")
	fmt.Print("Part 1: ")
	fmt.Println(parse(data, true))
	fmt.Print("Part 2: ")
	fmt.Println(parse(data, false))
}

func parse(data interface{}, redOK bool) (sum float64) {
	switch t := data.(type) {
	case float64:
		sum = t
	case []interface{}:
		for _, d := range t {
			sum += parse(d, redOK)
		}
	case map[string]interface{}:
		for _, d := range t {
			if d == "red" && !redOK {
				sum = 0
				break
			}
			sum += parse(d, redOK)
		}
	}
	return sum
}
