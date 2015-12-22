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

func parse(data interface{}, redOK bool) float64 {
	switch t := data.(type) {
	case float64:
		return t
	case string:
		return 0
	case []interface{}:
		var sum float64
		for _, d := range t {
			sum += parse(d, redOK)
		}
		return sum
	case map[string]interface{}:
		var sum float64
		for _, d := range t {
			if d == "red" && !redOK {
				return 0
			}
			sum += parse(d, redOK)
		}
		return sum
	}
	return 0
}
